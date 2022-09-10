package services

import (
	"net/url"
	"p-med/formvalidate"
	"p-med/models"
	"p-med/utils/page"

	"github.com/beego/beego/v2/client/orm"
)

// CorService struct
type CorService struct {
	BaseService
}

// GetPaginateData
func (us *CorService) GetPaginateData(listRows int, params url.Values) ([]*models.Cor, page.Pagination) {
	//Pesquisa, atribuição de campo de consulta
	us.SearchField = append(us.SearchField, new(models.Cor).SearchField()...)

	var cores []*models.Cor
	o := orm.NewOrm().QueryTable(new(models.Cor))
	_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&cores)
	if err != nil {
		return nil, us.Pagination
	}
	return cores, us.Pagination
}

// Create
func (*CorService) Create(form *formvalidate.CorForm) int {

	cor := models.Cor{
		Nome: form.Nome,
		Cor:  form.Cor[1:len(form.Cor)], // remove #

		Estado: 1,
	}

	//criptografia de senha
	id, err := orm.NewOrm().Insert(&cor)

	if err == nil {
		return int(id)
	}
	return 0
}

// GetCorById
func (*CorService) GetCorById(id int) *models.Cor {
	o := orm.NewOrm()
	cor := models.Cor{Id: id}
	err := o.Read(&cor)
	if err != nil {
		return nil
	}
	return &cor
}

// Update
func (*CorService) Update(form *formvalidate.CorForm) int {
	o := orm.NewOrm()
	cor := models.Cor{Id: form.Id}
	if o.Read(&cor) == nil {
		//
		cor.Nome = form.Nome
		cor.Cor = form.Cor[1:len(form.Cor)] // remove #
		cor.Estado = int8(form.Estado)

		num, err := o.Update(&cor)

		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Enable
func (*CorService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.Cor)).Filter("id__in", ids).Update(orm.Params{
		"estado": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable
func (*CorService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.Cor)).Filter("id__in", ids).Update(orm.Params{
		"estado": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del
func (*CorService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.Cor)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// GetExportData
func (us *CorService) GetExportData(params url.Values) []*models.Cor {
	//
	us.SearchField = append(us.SearchField, new(models.Cor).SearchField()...)
	var cor []*models.Cor
	o := orm.NewOrm().QueryTable(new(models.Cor))
	_, err := us.ScopeWhere(o, params).All(&cor)
	if err != nil {
		return nil
	}
	return cor
}

/* //Métodos públicos de cor
type corList struct {
	Id   int
	Nome string
	Cor  string
}

func corLists() (cl []corList) {
	corFilters := make([]interface{}, 0)
	corFilters = append(corFilters, "estado", 1)
	corResult, _ := GetPaginateData(1, 100, corFilters...)
	for _, cv := range corResult {
		corRow := corList{}
		corRow.Id = int(cv.Id)
		corRow.Nome = cv.Nome
		corRow.Cor = "#" + cv.Cor
		cl = append(cl, corRow)
	}
	return cl
} */
