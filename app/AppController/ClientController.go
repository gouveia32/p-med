package AppController

import (
	"pm/app/AppService"
	"pm/app/AppValidate"
	"pm/models"
)

type ClientController struct {
	AppBaseController
}

//Lista de clientes
func (c *ClientController) Index() {
	if c.IsHtml() {
		c.Display()
	} else {
		lists, err := AppService.GetClientLists(0, 15, "All")
		if err != nil {
			c.Abort500(err.Error(), "/")
		}
		c.Abort200(lists, "", "")
	}
}

func (c *ClientController) Save() {
	var (
		data models.Client
		maps models.Client
		err  error
	)

	if c.IsHtml() {
		id, _ := c.GetInt64("id", 0)
		c.Data["option"] = ""
		c.Data["isEdit"] = false

		if id != 0 {
			c.Data["isEdit"] = true
			//Nome de pesquisa
			maps.Id = id
			data, _ = AppService.GetClientByMap(&maps)
			c.Data["option"] = ""
		}
		//maps = data
		c.Data["data"] = data
		c.Display()
	} else {
		//Validação de campos de entrada
		params := models.Client{}
		if err := c.ParseForm(&params); err != nil {
			c.Abort500("Erro de parâmetro de entrada:"+err.Error(), "")
		}
		//Validação dos campos digitados
		ClientValidate := AppValidate.ClientValidate{}

		_ = c.ParseForm(&ClientValidate)
		err = ClientValidate.ValidClient()
		if err != nil {
			c.Abort500(err.Error(), "")
		}
		//Grava os dados
		err = AppService.SaveClient(&params)
		if err != nil {
			c.Abort500(err.Error(), "")
		}
		c.Abort200("", "Adicionado com sucesso", c.URLFor("ClientController.Lists"))
	}
}

//Exibição de lista
func (c *ClientController) Lists() {

	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := c.GetInt("limit")
	if err != nil {
		limit = 15
	}

	//paginator := pagination.SetPaginator(this, clientsPerPage, CountClients())

	if c.IsHtml() {
		c.Display("client/index")
	} else {
		lists, err := AppService.GetClientLists(page, limit, "All")
		if err != nil {
			c.Abort500(err.Error(), "/")
		}
		c.Data["json"] = lists
		//println("Estou aqui Lists: ", c)
		c.Abort200(lists, "", "")
	}
}

//Exibição de lista
func (c *ClientController) Search() {

	clientsPerPage := 15
	page, _ := c.GetInt("pag", 1)

	//paginator := pagination.SetPaginator(this, clientsPerPage, CountClients())

	if c.IsHtml() {
		c.Display("client/search")
	} else {
		lists, err := AppService.GetClientLists(page, clientsPerPage, "All")
		if err != nil {
			c.Abort500(err.Error(), "/")
		}
		c.Data["json"] = lists
		//println("Estou aqui Lists: ", c)
		c.Abort200(lists, "", "")
	}
}

//excluir
func (c *ClientController) Del() {
	//Excluir
	id, _ := c.GetInt64("id", 0)
	if id == 0 {
		c.Abort500("Dados não existem", "")
	}
	//Nome de pesquisa
	var maps models.Client
	maps.Id = id
	data, _ := AppService.GetClientByMap(&maps)
	if data.Id == 0 {
		c.Abort500("Os dados não existem, atualize e opere", "")
	}
	if err := AppService.DelClientByMap(&maps); err != nil {
		c.Abort500("Os dados não existem, atualize e opere:"+err.Error(), "")
	}
	c.Abort200("", "excluído com sucesso", "")
}
