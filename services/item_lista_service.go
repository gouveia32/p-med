package services

import (
	"net/url"
	"p-med/formvalidate"
	"p-med/models"
	"p-med/utils/page"
	"strconv"

	//"encoding/json"
	//"fmt"

	"github.com/beego/beego/v2/client/orm"
)

// ListaService struct
type ItemListaService struct {
	BaseService
}

// getDataBySettingGroupId Obtenha várias informações de configuração de acordo com o ID do grupo de configuração
func (*ItemListaService) getDataByItemListaId(listaId int) []*models.Lista {
	var listas []*models.Lista
	_, err := orm.NewOrm().QueryTable(new(models.Lista)).Filter("lista_id", listaId).All(&listas)
	if err != nil {
		return nil
	}
	return listas
}

// GetItemListaById
func (us *ItemListaService) GetItemListaById(id int64) *models.ItemLista {

	strId := strconv.FormatInt(int64(id), 10)
	item := make([]*models.ItemLista, 0)
	sql := "SELECT item_lista.id as id, lista.nome as nome, item_lista.descricao as descricao FROM item_lista JOIN lista ON  item_lista.lista_id = lista.id " +
		" WHERE lista.id=" + strId + " LIMIT 1;"

	orm.NewOrm().Raw(sql).QueryRows(&item)

	//fmt.Println("lista:", lista)

	return item[0]
}

// GetItemListaByNome
func (*ItemListaService) GetItemListaByNome(nome string) []*models.ItemLista {

	itens := make([]*models.ItemLista, 0)
	sql := "SELECT item_lista.id as id, lista.nome as nome, item_lista.descricao as descricao FROM item_lista JOIN lista ON  item_lista.lista_id = lista.id " +
		" WHERE lista.nome=" + nome

	orm.NewOrm().Raw(sql).QueryRows(&itens)

	//fmt.Println("itens:", itens)

	return itens
}

// GetPaginateData
func (us *ItemListaService) GetPaginateData(page, pageSize int, where string) ([]*models.ItemLista, page.Pagination) {
	us.Pagination.CurrentPage = page
	us.Pagination.ListRows = pageSize

	offset := (page - 1) * pageSize
	itens := make([]*models.ItemLista, 0)
	sql := "SELECT lista.id as id, lista.nome as nome, item_lista.descricao as descricao FROM item_lista JOIN lista ON  item_lista.lista_id = lista.id " +
		where + " ORDER BY lista.id DESC LIMIT ?,?;"

	orm.NewOrm().Raw(sql, offset, pageSize).QueryRows(&itens)

	//total := len(list)
	us.Pagination.Total = len(itens)
	return itens, us.Pagination //precisa ajustes na paginação
}

// Del
func (*ItemListaService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.Lista)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// Create
func (*ItemListaService) Create(form *formvalidate.ListaForm) int {

	lista := models.Lista{
		Nome: form.Nome,
		///Valor: form.Valor,
	}

	id, err := orm.NewOrm().Insert(&lista)

	if err == nil {
		return int(id)
	}
	return 0
}

// GetListaById
func (*ItemListaService) GetListaById(id int64) *models.Lista {
	o := orm.NewOrm()
	lista := models.Lista{Id: id}
	err := o.Read(&lista)
	if err != nil {
		return nil
	}
	return &lista
}

// Update
func (*ItemListaService) Update(form *formvalidate.ListaForm) int {
	o := orm.NewOrm()
	lista := models.Lista{Id: form.Id}

	if o.Read(&lista) == nil {
		//
		lista.Nome = form.Nome
		///lista.Valor = form.Valor

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
func (us *ItemListaService) GetExportData(params url.Values) []*models.Lista {
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
