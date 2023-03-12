package services

import (
	"net/url"
	"p-med/formvalidate"
	"p-med/models"
	"p-med/utils/page"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// NotaService struct
type NotaService struct {
	BaseService
}

// GetCount Pega quantidade de notas
func (*NotaService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.Nota)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

/* func (us *NotaService) NotaGetList(page, pageSize int, where string) ([]*PacienteAtend, page.Pagination) {
	us.Pagination.CurrentPage = page
	us.Pagination.ListRows = pageSize

	offset := (page - 1) * pageSize
	list := make([]*PacienteAtend, 0)
	sql := "SELECT id, nome, email, telefone, (SELECT COUNT(*) FROM nota AS a WHERE a.paciente_id = p.id  AND estado=1) AS qt_atend, estado FROM paciente AS p " +
		where + " ORDER BY p.id DESC LIMIT ?,?;"

	orm.NewOrm().Raw(sql, offset, pageSize).QueryRows(&list)

	//total := len(list)
	us.Pagination.Total = len(list)
	return list, us.Pagination //precisa ajustes na paginação
} */

func (us *NotaService) NotasGetList(page, pageSize int, where string) ([]*models.Nota, int64) {
	offset := (page - 1) * pageSize
	list := make([]*models.Nota, 0)
	sql := "SELECT * FROM nota " +
		where + " ORDER BY id LIMIT ?,?;"

	orm.NewOrm().Raw(sql, offset, pageSize).QueryRows(&list)

	total := len(list)
	return list, int64(total)
}

// GetPaginateData 
func (us *NotaService) GetPaginateData(listRows int, params url.Values) ([]*models.Nota, page.Pagination) {
	//Pesquisa, atribuição de campo de consulta
	us.SearchField = append(us.SearchField, new(models.Nota).SearchField()...)

	var notas []*models.Nota
	o := orm.NewOrm().QueryTable(new(models.Nota))
	_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&notas)
	if err != nil {
		return nil, us.Pagination
	}
	return notas, us.Pagination
}

// Create
func (*NotaService) Create(form *formvalidate.NotaForm) int {

	nota := models.Nota{
		Nome: form.Nome,

		Conteudo:      form.Conteudo,
		AtendimentoId: form.AtendId,
		CorId:         form.CorId,
		EtiquetaId:    form.EtiquetaId,
		CriadoEm:      int(time.Now().Unix()),
		AlteradoEm:    int(time.Now().Unix()),
		TipoNota:		1,
		Estado:        1,
	}

	//密码加密
	id, err := orm.NewOrm().Insert(&nota)

	if err == nil {
		return int(id)
	}
	return 0
}

// GetNotaById
func (*NotaService) GetNotaById(id int) *models.Nota {
	o := orm.NewOrm()
	nota := models.Nota{Id: id}

	err := o.Read(&nota)

	if err != nil {
		return nil
	}
	return &nota
}

// Update
func (*NotaService) Update(form *formvalidate.NotaForm) int {
	o := orm.NewOrm()
	nota := models.Nota{Id: form.NotaId}
	if o.Read(&nota) == nil {
		//
		nota.Nome = form.Nome
		nota.Conteudo = form.Conteudo
		nota.TipoNota = 1
		nota.CorId = form.CorId
		nota.EtiquetaId = form.EtiquetaId

		nota.CriadoEm = form.CriadoEm
		nota.AlteradoEm = int(time.Now().Unix())

		nota.Estado = 1

		num, err := o.Update(&nota)

		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Enable
func (*NotaService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.Nota)).Filter("id__in", ids).Update(orm.Params{
		"estado": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable
func (*NotaService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.Nota)).Filter("id__in", ids).Update(orm.Params{
		"estado": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del
func (*NotaService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.Nota)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// GetExportData
func (us *NotaService) GetExportData(params url.Values) []*models.Nota {
	//
	us.SearchField = append(us.SearchField, new(models.Nota).SearchField()...)
	var nota []*models.Nota
	o := orm.NewOrm().QueryTable(new(models.Nota))
	_, err := us.ScopeWhere(o, params).All(&nota)
	if err != nil {
		return nil
	}
	return nota
}
