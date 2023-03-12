package services

import (
	"net/url"
	"p-med/formvalidate"
	"p-med/models"
	"p-med/utils/page"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

// CampoService struct
type CampoService struct {
	BaseService
}

// GetCampoByNome
func (us *CampoService) GetCampoByNome(nome string) []*models.Campo {
	var campos []*models.Campo
	orm.NewOrm().QueryTable(new(models.Campo)).Filter("nome__in", nome).All(&campos)
	if len(campos) > 0 {
		return campos
	}
	return nil
}

// MontaCampo
func (us *CampoService) MontaCampo(cpo string) *models.Campo {
	campo := new(models.Campo)
	campo.Original = cpo
	w1 := strings.Split(cpo, "?")
	if len(w1) > 0 {
		campo.Descricao = w1[0]
		w2 := strings.Split(w1[1], ":")
		if len(w2) > 0 {
			campo.Nome = w2[0]
			w3 := strings.Split(w2[1], "=")
			if len(w3) > 1 {
				campo.Tipo = w3[0] //define multiplas respostas
				campo.Resposta = w3[1]
			} else {
				campo.Tipo = w3[0] //define multiplas respostas
				campo.Resposta = w3[0]
			}
		}
	}

	return campo
}

func (us *CampoService) CampoGetInConteudo(conteudo string) []*models.Campo {
	list := make([]*models.Campo, 0)

	words := strings.Split(conteudo, "]]")
	for _, word := range words {
		subwords := strings.Split(word, "[[")
		if len(subwords) > 1 {
			w := subwords[1]

			campo := us.MontaCampo(w)

			list = append(list, campo)

		}
	}

	return list
}

// GetPaginateData
func (us *CampoService) GetPaginateData(listRows int, params url.Values) ([]*models.Campo, page.Pagination) {
	//Pesquisa, atribuição de campo de consulta
	us.SearchField = append(us.SearchField, new(models.Campo).SearchField()...)

	var campos []*models.Campo
	o := orm.NewOrm().QueryTable(new(models.Campo))
	_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&campos)
	if err != nil {
		return nil, us.Pagination
	}
	return campos, us.Pagination
}

// Del
func (*CampoService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.Campo)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// Create
func (*CampoService) Create(form *formvalidate.CampoForm) int {

	campo := models.Campo{
		Nome:      form.Nome,
		Descricao: form.Descricao,
		Resposta:  form.Resposta,
	}

	//criptografia de senha
	id, err := orm.NewOrm().Insert(&campo)

	if err == nil {
		return int(id)
	}
	return 0
}

// GetCampoById
func (*CampoService) GetCampoById(id int64) *models.Campo {
	o := orm.NewOrm()
	campo := models.Campo{Id: id}
	err := o.Read(&campo)
	if err != nil {
		return nil
	}
	return &campo
}

// Update
func (*CampoService) Update(form *formvalidate.CampoForm) int {
	o := orm.NewOrm()
	campo := models.Campo{Id: form.Id}
	if o.Read(&campo) == nil {
		//
		campo.Nome = form.Nome
		campo.Descricao = form.Descricao
		campo.Resposta = form.Resposta

		num, err := o.Update(&campo)

		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// GetExportData
func (us *CampoService) GetExportData(params url.Values) []*models.Campo {
	//
	us.SearchField = append(us.SearchField, new(models.Campo).SearchField()...)
	var campo []*models.Campo
	o := orm.NewOrm().QueryTable(new(models.Campo))
	_, err := us.ScopeWhere(o, params).All(&campo)
	if err != nil {
		return nil
	}
	return campo
}
