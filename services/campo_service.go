package services

import (
	"net/url"
	"p-med/formvalidate"
	"p-med/models"
	"p-med/utils/page"
	//"encoding/json"
	"strings"
	//"fmt"

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
	campo := new(models.Campo) //[[descricao?nome:tipo=resultador]]
	campo.Original = cpo
	w1 := strings.Split(cpo, "?")
	if len(w1) > 0 {
		campo.Descricao = w1[0]
		w2 := strings.Split(w1[1], ":")
		if len(w2) > 0 {
			campo.Nome = w2[0]
			w3 := strings.Split(w2[1], "=")
			campo.Tipo = w3[0]
			if len(w3) > 1 {
				campo.Resposta = w3[1]
				rsps := strings.Split(w3[1], ",")
				var respostas []*models.Resposta
				if len(rsps) > 1 {
					for _, r := range rsps {
						//fmt.Println("r: ", r)
						rs := new(models.Resposta)

						rs.Resposta = r

						respostas = append(respostas, rs)
					}
					campo.RespostaStrut = respostas
				} else {
					campo.Resposta = rsps[0]
				}

				//err := json.Unmarshal([]byte(campo.Resposta), &campo.RespostaStrut)
				//fmt.Println(err)
			} else {
				campo.Resposta = ""
			}
		}
	}
	
	//fmt.Println("\ncampo.Descricao: ",campo.Descricao)
	//fmt.Println("campo.Nome: ",campo.Nome)
	//fmt.Println("campo.Tipo: ",campo.Tipo)
	//fmt.Println("campo.Resposta: ",campo.Resposta)

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
	//fmt.Println("err",err)
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
		Original: form.Original,
		ValorInicial: form.ValorInicial,
		Tipo: form.Tipo,

	}

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
		campo.Original = form.Original
		campo.Tipo = form.Tipo
		campo.ValorInicial = form.ValorInicial

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
