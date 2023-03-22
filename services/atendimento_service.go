package services

import (
	"net/url"
	"p-med/formvalidate"
	"p-med/models"
	"p-med/utils/page"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type PacienteAtend struct {
	Id       int
	Nome     string
	Email    string
	Telefone string
	Estado   int
	QtAtend  int
}

// AtendimentoService struct
type AtendimentoService struct {
	BaseService
}

// GetCount Pega quantidade de atendimentos
func (*AtendimentoService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.Atendimento)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

func (us *AtendimentoService) PacienteAtendGetList(page, pageSize int, where string) ([]*PacienteAtend, page.Pagination) {
	us.Pagination.CurrentPage = page
	us.Pagination.ListRows = pageSize

	offset := (page - 1) * pageSize
	list := make([]*PacienteAtend, 0)
	sql := "SELECT id, nome, email, telefone, (SELECT COUNT(*) FROM atendimento AS a WHERE a.paciente_id = p.id  AND estado=1) AS qt_atend, estado FROM paciente AS p " +
		where + " ORDER BY p.id DESC LIMIT ?,?;"

	orm.NewOrm().Raw(sql, offset, pageSize).QueryRows(&list)

	//total := len(list)
	us.Pagination.Total = len(list)
	return list, us.Pagination //precisa ajustes na paginação
}

func (us *AtendimentoService) AtendimentoGetList(page, pageSize int, where string) ([]*models.Atendimento, page.Pagination) {
	us.Pagination.CurrentPage = page
	us.Pagination.ListRows = pageSize

	offset := (page - 1) * pageSize
	list := make([]*models.Atendimento, 0)
	sql := "SELECT * FROM atendimento " +
		where + " ORDER BY id DESC LIMIT ?,?;"

	orm.NewOrm().Raw(sql, offset, pageSize).QueryRows(&list)

	//total := len(list)
	us.Pagination.Total = len(list)
	return list, us.Pagination //precisa ajustes na paginação
}

// GetPaginateData Obtenha atendimento por paginação
func (us *AtendimentoService) GetPaginateDataAtend(listRows int, params url.Values) ([]*models.Atendimento, page.Pagination) {
	//Pesquisa, atribuição de campo de consulta
	us.SearchField = append(us.SearchField, new(models.Atendimento).SearchField()...)

	var atendimentos []*models.Atendimento

	db := orm.NewOrm()
	o := db.QueryTable(new(models.Atendimento)).Distinct().RelatedSel("Paciente")
	_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&atendimentos)

	///sql := "SELECT id, nome, email, telefone, (SELECT COUNT(*) FROM atendimento AS a WHERE a.paciente_id = p.id  AND estado=1) AS qt_atend, estado FROM paciente AS p "
	///orm.NewOrm().Raw(sql, params).QueryRows(&atendimentos)
	///_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&atendimentos)
	if err != nil {
		return nil, us.Pagination
	}
	return atendimentos, us.Pagination
}

// GetPaginateData Obtenha atendimento por paginação
func (us *AtendimentoService) GetPaginateData(listRows int, params url.Values) ([]*models.Atendimento, page.Pagination) {
	//Pesquisa, atribuição de campo de consulta
	us.SearchField = append(us.SearchField, new(models.Atendimento).SearchField()...)

	var atendimentos []*models.Atendimento
	o := orm.NewOrm().QueryTable(new(models.Atendimento))
	///sql := "SELECT id, nome, email, telefone, (SELECT COUNT(*) FROM atendimento AS a WHERE a.paciente_id = p.id  AND estado=1) AS qt_atend, estado FROM paciente AS p "
	///orm.NewOrm().Raw(sql, params).QueryRows(&atendimentos)
	_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&atendimentos)
	if err != nil {
		return nil, us.Pagination
	}
	return atendimentos, us.Pagination
}

// Create
func (*AtendimentoService) Create(form *formvalidate.AtendimentoForm) int {

	atendimento := models.Atendimento{
		Nome:     form.Nome,
		Conteudo: form.Conteudo,

		CorId:      form.CorId,
		PacienteId: form.PacienteId,
		EtiquetaId: 1,
		CriadoEm:   int(time.Now().Unix()),
		AlteradoEm: int(time.Now().Unix()),
		Estado:     1,
	}

	//criptografia de senha
	id, err := orm.NewOrm().Insert(&atendimento)

	if err == nil {
		return int(id)
	}
	return 0
}

// GetAtendimentoById
func (*AtendimentoService) GetAtendimentoById(id int) *models.Atendimento {
	o := orm.NewOrm()
	atendimento := models.Atendimento{Id: id}
	err := o.Read(&atendimento)
	if err != nil {
		return nil
	}
	return &atendimento
}

// Update
func (*AtendimentoService) Update(form *formvalidate.AtendimentoForm) int {
	o := orm.NewOrm()
	atendimento := models.Atendimento{Id: form.AtendId}
	if o.Read(&atendimento) == nil {
		//
		atendimento.Nome = form.Nome
		atendimento.Conteudo = form.Conteudo
		atendimento.CorId = form.CorId
		atendimento.PacienteId = form.PacienteId
		atendimento.EtiquetaId = 1

		atendimento.AlteradoEm = int(time.Now().Unix())
		atendimento.Estado = 1

		num, err := o.Update(&atendimento)

		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Enable
func (*AtendimentoService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.Atendimento)).Filter("id__in", ids).Update(orm.Params{
		"estado": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable
func (*AtendimentoService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.Atendimento)).Filter("id__in", ids).Update(orm.Params{
		"estado": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del
func (*AtendimentoService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.Atendimento)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// GetExportData
func (us *AtendimentoService) GetExportData(params url.Values) []*models.Atendimento {
	//
	us.SearchField = append(us.SearchField, new(models.Atendimento).SearchField()...)
	var atendimento []*models.Atendimento
	o := orm.NewOrm().QueryTable(new(models.Atendimento))
	_, err := us.ScopeWhere(o, params).All(&atendimento)
	if err != nil {
		return nil
	}
	return atendimento
}
