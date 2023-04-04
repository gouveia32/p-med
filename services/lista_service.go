package services

import (
	"net/url"
	"p-med/formvalidate"
	"p-med/models"
	"p-med/utils/page"
	//"encoding/json"
	//"fmt"

	"github.com/beego/beego/v2/client/orm"
)

// ListaService struct
type ListaService struct {
	BaseService
}

// getDataByListaId 
func (*ListaService) getDataByListaId(listaId int) []*models.Lista {
	var listas []*models.Lista
	_, err := orm.NewOrm().QueryTable(new(models.Lista)).Filter("lista_id", listaId).All(&listas)
	if err != nil {
		return nil
	}
	return listas
}


// GetListaByNome
func (us *ListaService) GetListaByNome(nome string) []*models.Lista {
	var listas []*models.Lista
	orm.NewOrm().QueryTable(new(models.Lista)).Filter("nome__in", nome).All(&listas)
	if len(listas) > 0 {
		return listas
	}
	return nil
}


// GetPaginateData
func (us *ListaService) GetPaginateData(listRows int, params url.Values) ([]*models.Lista, page.Pagination) {
	//Pesquisa, atribuição de lista de consulta
	
	us.SearchField = append(us.SearchField, new(models.Lista).SearchField()...)

	var listas []*models.Lista
	o := orm.NewOrm().QueryTable(new(models.Lista))
	
	_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&listas)
	//fmt.Println("err",err)
	if err != nil {
		return nil, us.Pagination
	}
	
	return listas, us.Pagination
}

// Del
func (*ListaService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.Lista)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// Create
func (*ListaService) Create(form *formvalidate.ListaForm) int {

	lista := models.Lista{
		Nome:   form.Nome,
		Valor: 	form.Valor,
	}

	id, err := orm.NewOrm().Insert(&lista)

	if err == nil {
		return int(id)
	}
	return 0
}

// GetListaById
func (*ListaService) GetListaById(id int64) *models.Lista {
	o := orm.NewOrm()
	lista := models.Lista{Id: id}
	err := o.Read(&lista)
	if err != nil {
		return nil
	}
	return &lista
}

// Update
func (*ListaService) Update(form *formvalidate.ListaForm) int {
	o := orm.NewOrm()
	lista := models.Lista{Id: form.Id}
	
	if o.Read(&lista) == nil {
		//
		lista.Nome = form.Nome
		lista.Valor = form.Valor

		num, err := o.Update(&lista)

/* 		fmt.Println("err:",err)
		fmt.Println("valor:",lista.Valor) */
		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// GetExportData
func (us *ListaService) GetExportData(params url.Values) []*models.Lista {
	//
	us.SearchField = append(us.SearchField, new(models.Lista).SearchField()...)
	var lista []*models.Lista
	o := orm.NewOrm().QueryTable(new(models.Lista))
	_, err := us.ScopeWhere(o, params).All(&lista)
	if err != nil {
		return nil
	}
	return lista
}
