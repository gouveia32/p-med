package services

import (
	"fmt"
	"net/url"
	"os"
	"p-med/formvalidate"
	"p-med/models"
	"p-med/utils/page"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// ModeloService struct
type ModeloService struct {
	BaseService
}

// GetPaginateData
func (us *ModeloService) GetPaginateData(listRows int, params url.Values) ([]*models.Modelo, page.Pagination) {
	//Pesquisa, atribuição de campo de consulta
	us.SearchField = append(us.SearchField, new(models.Modelo).SearchField()...)

	var modelos []*models.Modelo
	o := orm.NewOrm().QueryTable(new(models.Modelo))
	_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&modelos)
	if err != nil {
		return nil, us.Pagination
	}
	return modelos, us.Pagination
}

// Create
func (*ModeloService) Create(form *formvalidate.ModeloForm) int {

	modelo := models.Modelo{
		Nome:       form.Nome,
		Detalhe:    form.Detalhe,
		CriadoEm:   int(time.Now().Unix()),
		AlteradoEm: int(time.Now().Unix()),

		Estado: 1,
	}

	//criptografia de senha
	id, err := orm.NewOrm().Insert(&modelo)

	if err == nil {
		return int(id)
	}
	return 0
}

// GetModeloById
func (*ModeloService) GetModeloById(id int64) *models.Modelo {
	o := orm.NewOrm()
	modelo := models.Modelo{Id: id}
	err := o.Read(&modelo)
	if err != nil {
		return nil
	}
	return &modelo
}

// Ajustes
func (*ModeloService) Ajustes(modelo *models.Modelo, campos []*models.Campo) string {

	// open output file
	fo, err := os.Create("views/import/consulta.tpl")
	if err != nil {
		panic(err)
	}


	//processar o arquivo
	strOriginal := "<section class='content'>" +
		"<form id='form' name='form' class='form-horizontal dataForm' action='/admin/import/consulta' method='post'	enctype='multipart/form-data'>" +
		"	<label class='label-micro'>Paciente: {{.paciente.Nome}}</label>" +
/* 		"	<button type='submit' class='btn flat btn-info dataFormSubmit toolTop'>" +
		"		Processar" +
		"	</button>" +
		"	<button type='reset' class='btn flat btn-default dataFormReset'>" +
		"		Restaurar" +
		"	</button>" + */
		"	<hr>" + modelo.Detalhe

	for i, v := range campos {
		switch v.Tipo {
		case "data":
			cmd := fmt.Sprintf("<input type='text' title='%s' name='%s' value='%s' ></input>",
				v.Descricao, v.Nome, v.ValorInicial)

			p1 := strings.Index(strOriginal, v.Original) - 2 // desconta [[
			strOriginal = strOriginal[:p1] +
				cmd +
				strOriginal[p1+len(v.Original)+4:]
		case "texto":
			cmd := fmt.Sprintf("<input type='text' title='%s' name='cpo%s' value='%s' ></input>",
				v.Descricao, v.Nome, v.Resposta)

			p1 := strings.Index(strOriginal, v.Original) - 2 // desconta [[
			strOriginal = strOriginal[:p1] +
				cmd +
				strOriginal[p1+len(v.Original)+4:]
		case "escolha":
			cmd := ""
			params := strings.Split(v.Resposta, ",")
			for _, p := range params {
				cmd += fmt.Sprintf("<input type='radio' title='%s' name='cpo%d' id='%s' onchange='atualiza(this.id,this,this.value)' > %s </input>",
					v.Descricao, i, v.Nome, p)
			}

			p1 := strings.Index(strOriginal, v.Original) - 2 // desconta [[
			strOriginal = strOriginal[:p1] +
				cmd +
				strOriginal[p1+len(v.Original)+4:]
		}
		//fmt.Println(strOriginal)

	}
	strOriginal += "	</form>" +
		"</section>" +
		"<script src='/static/layui/layui.js?s=36'></script>" +
		"<script language='javascript'>" +
		"var $,form;" +

		"layui.config({" +
		"base : 'js/'" +
		"}).use(['form','element','layer','jquery'],function(){" +
		"form = layui.form;" +
		"$ = layui.jquery;" +

		"});" +

		"function atualiza(key,valor) {" +
		"$('form[name=form]').find('input[name=cpo' + key + ']').val(valor);" +

		"console.log(key);" +
		"console.log(valor);" +

		"}" +
		"$('.toolTop').click(function () {" +
		"alert('Estou aqui...');" +
		"});" +
		"</script>"

	// write a chunk

	_, err = fo.WriteString(strOriginal)

	if err != nil {
		panic(err)
	}

	//fmt.Println(strOriginal)
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	return strOriginal
}

// Update
func (*ModeloService) Update(form *formvalidate.ModeloForm) int {
	o := orm.NewOrm()
	modelo := models.Modelo{Id: form.Id}
	if o.Read(&modelo) == nil {
		//
		modelo.Nome = form.Nome
		modelo.Detalhe = form.Detalhe

		modelo.Estado = int8(form.Estado)
		modelo.AlteradoEm = int(time.Now().Unix())

		num, err := o.Update(&modelo)

		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Enable
func (*ModeloService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.Modelo)).Filter("id__in", ids).Update(orm.Params{
		"estado": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable
func (*ModeloService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.Modelo)).Filter("id__in", ids).Update(orm.Params{
		"estado": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del
func (*ModeloService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.Modelo)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// GetExportData
func (us *ModeloService) GetExportData(params url.Values) []*models.Modelo {
	//
	us.SearchField = append(us.SearchField, new(models.Modelo).SearchField()...)
	var modelo []*models.Modelo
	o := orm.NewOrm().QueryTable(new(models.Modelo))
	_, err := us.ScopeWhere(o, params).All(&modelo)
	if err != nil {
		return nil
	}
	return modelo
}
