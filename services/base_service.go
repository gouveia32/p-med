package services

import (
	"net/url"
	"p-med/utils"
	"p-med/utils/page"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// BaseService struct
type BaseService struct {
	//campos pesquisáveis
	SearchField []string
	//Campos que podem ser usados ​​como condições
	WhereField []string
	//Campos que podem ser usados ​​como consultas de intervalo de tempo
	TimeField []string
	//O id de dados proibido de ser excluído pode ser declarado no modelo, mas não é necessário declará-lo aqui
	//NoDeletionId []int
	//paginação
	Pagination page.Pagination
}

// Paginate paginação
func (bs *BaseService) Paginate(seter orm.QuerySeter, listRows int, parameters url.Values) orm.QuerySeter {
	var pagination page.Pagination
	qs := pagination.Paginate(seter, listRows, parameters)
	bs.Pagination = pagination
	return qs
}

// ScopeWhere processamento de consultas
func (bs *BaseService) ScopeWhere(seter orm.QuerySeter, parameters url.Values) orm.QuerySeter {
	//palavra-chave como pesquisa
	keywords := parameters.Get("_keywords")
	cond := orm.NewCondition()
	if keywords != "" && len(bs.SearchField) > 0 {
		for _, v := range bs.SearchField {
			cond = cond.Or(v+"__icontains", keywords)
		}
	}

	//Consulta de condição de campo
	if len(bs.WhereField) > 0 && len(parameters) > 0 {
		for k, v := range parameters {
			if v[0] != "" && utils.InArrayForString(bs.WhereField, k) {
				cond = cond.And(k, v[0])
			}
		}
	}

	//consulta de intervalo de tempo
	if len(bs.TimeField) > 0 && len(parameters) > 0 {
		for key, value := range parameters {
			if value[0] != "" && utils.InArrayForString(bs.TimeField, key) {
				timeRange := strings.Split(value[0], " - ")
				startTimeStr := timeRange[0]
				endTimeStr := timeRange[1]

				loc, _ := time.LoadLocation("Local")
				startTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, loc)

				if err == nil {
					unixStartTime := startTime.Unix()
					if len(endTimeStr) == 10 {
						endTimeStr += "23:59:59"
					}

					endTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, loc)
					if err == nil {
						unixEndTime := endTime.Unix()
						cond = cond.And(key+"__gte", unixStartTime).And(key+"__lte", unixEndTime)
					}
				}
			}
		}
	}

	//Montar declarações condicionais na declaração principal
	seter = seter.SetCond(cond)

	//ordenar
	order := parameters.Get("_order")
	by := parameters.Get("_by")
	if order == "" {
		order = "id"
	}

	if by == "" {
		by = "-"
	} else {
		if by == "asc" {
			by = ""
		} else {
			by = "-"
		}
	}

	//ordenar
	seter = seter.OrderBy(by + order)

	return seter
}

// PaginateAndScopeWhere Mesclagem de paginação e consulta, usada principalmente para exibição e pesquisa da lista da página inicial
func (bs *BaseService) PaginateAndScopeWhere(seter orm.QuerySeter, listRows int, parameters url.Values) orm.QuerySeter {
	return bs.Paginate(bs.ScopeWhere(seter, parameters), listRows, parameters)
}
