package services

import (
	"fmt"
	"p-med/formvalidate"
	"p-med/models"
	"p-med/utils/page"
	"strconv"

	//"encoding/json"

	"github.com/beego/beego/v2/client/orm"
)

// ListaService struct
type ListaService struct {
	BaseService
}

// getDataByListaId
func (us *ListaService) getDataByListaId(listaId int) []*models.Lista {
	var listas []*models.Lista
	_, err := orm.NewOrm().QueryTable(new(models.Lista)).Filter("lista_id", listaId).All(&listas)
	if err != nil {
		return nil
	}
	return listas
}

// GetListaByNome
/* func (us *ListaService) GetListaByNome(nome string) []*models.Lista {
	var listas []*models.Lista
	orm.NewOrm().QueryTable(new(models.Lista)).Filter("nome__in", nome).All(&listas)
	if len(listas) > 0 {
		return listas
	}
	return nil
} */

// GetPaginateDataLista
func (us *ListaService) GetPaginateDataLista(page, pageSize int, where string) ([]*models.Lista, page.Pagination) {
	us.Pagination.CurrentPage = page
	us.Pagination.ListRows = pageSize

	offset := (page - 1) * pageSize
	itens := make([]*models.Lista, 0)
	sql := "SELECT id, nome FROM lista " +
		where + " ORDER BY id DESC LIMIT ?,?;"

	orm.NewOrm().Raw(sql, offset, pageSize).QueryRows(&itens)

	//total := len(list)
	us.Pagination.Total = len(itens)
	return itens, us.Pagination //precisa ajustes na paginação
}

// GetPaginateDataItemLista
func (us *ListaService) GetPaginateDataItemLista(page, pageSize int, where string) ([]*models.ItemLista, page.Pagination) {
	us.Pagination.CurrentPage = page
	us.Pagination.ListRows = pageSize

	offset := (page - 1) * pageSize
	itens := make([]*models.ItemLista, 0)
	sql := "SELECT item_lista.id as id, lista.nome as nome, item_lista.descricao as descricao FROM lista JOIN item_lista ON lista.id = item_lista.lista_id " +
		where + " ORDER BY id DESC LIMIT ?,?;"

	orm.NewOrm().Raw(sql, offset, pageSize).QueryRows(&itens)

	//total := len(list)
	us.Pagination.Total = len(itens)
	return itens, us.Pagination //precisa ajustes na paginação
}


// GetListaById
func (us *ListaService) GetListaById(id int64) *models.Lista {

	strId := strconv.FormatInt(int64(id), 10)
	lista := make([]*models.Lista, 0)
	sql := "SELECT id, nome, descricao FROM lista JOIN lista" +
		" WHERE id=" + strId + " LIMIT 1;"

	orm.NewOrm().Raw(sql).QueryRows(&lista)

	//fmt.Println("lista:", lista)

	return lista[0]
}

// GetListaByNome
func (*ListaService) GetListaByNome(nome string) []*models.Lista {

	lista := make([]*models.Lista, 0)
	sql := "SELECT lista.id, lista.nome, item_lista.descricao FROM item_lista JOIN lista ON  item_lista.lista_id = lista.id " +
		" WHERE lista.nome=" + nome

	orm.NewOrm().Raw(sql).QueryRows(&lista)

	//fmt.Println("lista:", lista)

	return lista
}

// Update
func (us *ListaService) Update(form *formvalidate.ListaForm) int {

	o := orm.NewOrm()
	lista := models.Lista{Id: form.Id}

	//lista1 := us.GetListaById(form.Id)

	fmt.Println("lista:", lista)

	if o.Read(&lista) == nil {
		//
		lista.Nome = form.Nome
		//lista.Valor = form.Valor

		num, err := o.Update(&lista)

		fmt.Println("err:", err)

		if err == nil {
			return int(num)
		}
		return 0
	}

	return 0
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
		Nome: form.Nome,
		///Valor: form.Valor,
	}

	//criptografia de senha
	id, err := orm.NewOrm().Insert(&lista)

	if err == nil {
		return int(id)
	}
	return 0
}

