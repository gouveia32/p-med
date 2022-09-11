package services

import (
	"fmt"
	"net/url"
	"p-med/formvalidate"
	"p-med/models"
	"p-med/utils/page"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// PacienteService struct
type PacienteService struct {
	BaseService
}

// GetCount Pega quantidade de pacientes
func (*PacienteService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.Paciente)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetPaginateDataObter paciente por paginação
func (us *PacienteService) GetPaginateData(listRows int, params url.Values) ([]*models.Paciente, page.Pagination) {
	//Pesquisa, atribuição de campo de consulta
	us.SearchField = append(us.SearchField, new(models.Paciente).SearchField()...)

	var pacientes []*models.Paciente
	o := orm.NewOrm().QueryTable(new(models.Paciente))
	_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&pacientes)
	if err != nil {
		return nil, us.Pagination
	}
	return pacientes, us.Pagination
}

// Create
func (*PacienteService) Create(form *formvalidate.PacienteForm) int {

	dn := form.Nascimento
	if dn == "" {
		dn = "001-01-01"
	}

	paciente := models.Paciente{
		Nome: form.Nome,

		Sexo:       form.Sexo,
		Logradoro:  form.Logradoro,
		Municipio:  form.Municipio,
		Numero:     form.Numero,
		Uf:         form.Uf,
		Cep:        form.Cep,
		Altura:     form.Altura,
		Peso:       form.Peso,
		Telefone:   form.Telefone,
		Email:      form.Email,
		CriadoEm:   form.CriadoEm,
		AlteradoEm: form.AlteradoEm,

		Nascimento: dn,
		Estado:     1,
	}

	//criptografia de senha
	id, err := orm.NewOrm().Insert(&paciente)

	if err == nil {
		return int(id)
	}
	return 0
}

// GetPacienteById
func (*PacienteService) GetPacienteById(id int64) *models.Paciente {
	o := orm.NewOrm()
	paciente := models.Paciente{Id: id}
	err := o.Read(&paciente)
	if err != nil {
		return nil
	}
	return &paciente
}

// Update
func (*PacienteService) Update(form *formvalidate.PacienteForm) int {
	o := orm.NewOrm()
	paciente := models.Paciente{Id: form.Id}
	if o.Read(&paciente) == nil {
		//
		paciente.Nome = form.Nome
		paciente.Sexo = form.Sexo
		paciente.Logradoro = form.Logradoro
		paciente.Municipio = form.Municipio
		paciente.Numero = form.Numero
		paciente.Uf = form.Uf
		paciente.Cep = form.Cep
		paciente.Altura = form.Altura
		paciente.Peso = form.Peso
		paciente.Telefone = form.Telefone
		paciente.Email = form.Email
		paciente.CriadoEm = form.CriadoEm
		paciente.AlteradoEm = int(time.Now().Unix())
		paciente.Estado = int8(form.Estado)

		if form.Nascimento != "" {
			paciente.Nascimento = form.Nascimento
		} else {
			paciente.Nascimento = "0001-01-01"
		}

		num, err := o.Update(&paciente)

		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Enable
func (*PacienteService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.Paciente)).Filter("id__in", ids).Update(orm.Params{
		"estado": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable
func (*PacienteService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.Paciente)).Filter("id__in", ids).Update(orm.Params{
		"estado": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del
func (*PacienteService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.Paciente)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// GetExportData
func (us *PacienteService) GetExportData(params url.Values) []*models.Paciente {
	//
	us.SearchField = append(us.SearchField, new(models.Paciente).SearchField()...)
	var paciente []*models.Paciente
	o := orm.NewOrm().QueryTable(new(models.Paciente))
	_, err := us.ScopeWhere(o, params).All(&paciente)
	if err != nil {
		return nil
	}
	return paciente
}

func (us *PacienteService) PacienteChangeNoSelecionado(id int64, no int64) (num int64, err error) {

	sql := fmt.Sprintf("UPDATE paciente set no_selecionado=? where id = %d;", id)
	res, err := orm.NewOrm().Raw(sql, no).Exec()
	num = 0
	if err == nil {
		num, _ = res.RowsAffected()
	}
	return num, err
}
