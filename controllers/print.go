/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 10:21:13
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:04:41
***********************************************/
package controllers

import (
	"p-med/models"
	"p-med/services"
)

type PrintController struct {
	baseController
}

func (self *PrintController) Atendimento() {
	paciente_id, _ := self.GetInt64("pacienteId", 0)
	paciente := new(models.Paciente)
	if (paciente_id) != 0 {
		var pacienteService services.PacienteService
		paciente = pacienteService.GetPacienteById(paciente_id)
		self.Data["paciente"] = paciente
	}

	aId, _ := self.GetInt("AtendId", 0)
	if (aId) != 0 {
		var atendimentoService services.AtendimentoService
		atendimento := atendimentoService.GetAtendimentoById(aId)
		self.Data["atendimento"] = atendimento
	}
	self.Data["pageTitle"] = "Impressão de Atendimento"
	//self.display()
	self.TplName = "print/atendimento.html"
}

func (self *PrintController) Nota() {
	paciente_id, _ := self.GetInt64("pacienteId", 0)
	if (paciente_id) != 0 {
		var pacienteService services.PacienteService
		paciente := pacienteService.GetPacienteById(paciente_id)
		self.Data["paciente"] = paciente
	}
	nId, _ := self.GetInt("NotaId", 0)
	if (nId) != 0 {
		var notaService services.NotaService
		nota := notaService.GetNotaById(nId)
		self.Data["nota"] = nota
	}
	self.Data["pageTitle"] = "Impressão de Nota"
	//self.display()
	self.TplName = "print/nota.html"
}
