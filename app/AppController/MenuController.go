package AppController

import (
	"database/sql"
	"strconv"

	"pm/app/AppService"
	"pm/app/AppValidate"
	"pm/models"
)

type MenuController struct {
	AppBaseController
}

//查询菜单列表
func (c *MenuController) Index() {
	if c.IsHtml() {
		c.Display()
	} else {
		lists, err := AppService.GetMenuLists(1, "")
		if err != nil {
			c.Abort500(err.Error(), "/")
		}
		c.Abort200(lists, "", "")
	}
}

//Adicionar e editar menu
func (c *MenuController) Save() {
	var (
		maps      models.Menu
		data      models.Menu
		maps2     models.Menu
		err       error
		null      sql.NullInt64
		parent_id int
	)
	if c.IsHtml() {
		id, _ := c.GetInt("id", 0)
		parent_id, _ := c.GetInt("parent_id", 0)
		c.Data["option"] = ""
		c.Data["isEdit"] = false
		if parent_id != 0 {
			//Nome de pesquisa
			maps.Id = parent_id
			data, _ := AppService.GetMenuByMap(&maps)
			c.Data["option"] = "<option value='" + strconv.FormatInt(int64(parent_id), 10) + "' selected>" + data.Name + "</option>"
		} else if id != 0 {
			c.Data["isEdit"] = true
			//Nome de pesquisa
			maps.Id = id
			//println("Estou aqui GetMenuByMap antes: ", data.Id)
			data, _ = AppService.GetMenuByMap(&maps)
			//println("Estou aqui GetMenuByMap depois: ", data.Id)
			c.Data["option"] = ""
			if data.NullParentId != null { //Encontre informações sobre os pais
				parent_id = int(data.NullParentId.Int64)
				//Nome de pesquisa
				maps2.Id = parent_id
				data, _ := AppService.GetMenuByMap(&maps2)
				c.Data["option"] = "<option value='" + strconv.FormatInt(int64(parent_id), 10) + "' selected>" + data.Name + "</option>"
			}
		}
		c.Data["data"] = data
		c.Display()
	} else {
		params := models.Menu{}
		parent_id, err = c.GetInt("parent_id", 0)
		if parent_id != 0 { //Atribua um valor a nullParentId como os dados finais adicionados
			params.NullParentId.Int64, params.NullParentId.Valid = int64(parent_id), true
		}
		if err := c.ParseForm(&params); err != nil {
			c.Abort500("Erro de parâmetro de entrada:"+err.Error(), "")
		}

		//Validação de campos de entrada
		MenuValidate := AppValidate.MenuValidate{}
		_ = c.ParseForm(&MenuValidate)
		err = MenuValidate.ValidMenu()
		if err != nil {
			c.Abort500(err.Error(), "")
		}
		//Adição de dados
		params.ParentId = int(params.NullParentId.Int64)
		err = AppService.SaveMenu(&params)
		if err != nil {
			c.Abort500(err.Error(), "")
		}
		c.Abort200("", "Adicionado com sucesso", c.URLFor("MenuController.Lists"))
	}
}

//Exibição de lista
func (c *MenuController) Lists() {
	if c.IsHtml() {
		//println("Estou aqui: ",c.IsHtml)
		c.Display("menu/index")
	} else {
		lists, err := AppService.GetMenuLists(1, "All")
		if err != nil {
			c.Abort500(err.Error(), "/")
		}
		c.Data["json"] = lists
		c.Abort200(lists, "", "")
	}
}

//excluir
func (c *MenuController) Del() {
	//Excluir pai excluir todos os filhos
	id, _ := c.GetInt("id", 0)
	if id == 0 {
		c.Abort500("Dados não existem", "")
	}
	//Nome de pesquisa
	var maps models.Menu
	maps.Id = id
	data, _ := AppService.GetMenuByMap(&maps)
	if data.Id == 0 {
		c.Abort500("Os dados não existem, atualize e opere", "")
	}
	if err := AppService.DelMenuByMap(&maps); err != nil {
		c.Abort500("Os dados não existem, atualize e opere:"+err.Error(), "")
	}
	c.Abort200("", "excluído com sucesso", "")
}
