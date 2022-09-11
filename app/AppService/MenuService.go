/**
后台菜单的服务层
*/
package AppService

import (
	"pm/models"

	"github.com/astaxie/beego"
)

type Menu struct {
	models.Menu
	Url      string
	Children []Menu
}

var (
	ModelMenuLists models.Menu
	ModelMenuAdd   models.Menu
	ModelMenuOne   models.Menu
	ModelMenuDel   models.Menu
)

//A combinação torna-se um menu mais espesso
func GetMenuLists(module_id int8, data_type string) (lists []map[string]interface{}, err error) {
	ModelMenuLists.ModuleBelong = module_id
	var menus []models.Menu
	if data_type == "All" {
		ModelMenuLists.Status = 0 //Quando a estrutura consulta 0, a condição da consulta será inválida. Se você precisar consultar o 0 especial, deve ser a consulta da condição
		menus, err = ModelMenuLists.AllMenu()
	} else {
		//Encontre o top primeiro
		ModelMenuLists.Status = 1
		menus, err = ModelMenuLists.QueryMenuLists(0)
		if err != nil {
			return lists, err
		}
	}
	lists = getMenuData(menus, data_type)

	return lists, nil
}

//Menu Infinitus para obter um subconjunto
func getMenuData(menus []models.Menu, data_type string) []map[string]interface{} {
	var MenusData = make([]map[string]interface{}, len(menus))
	for key, item := range menus {
		row := make(map[string]interface{})
		row["mid"] = item.Id
		row["parent_id"] = item.ParentId
		row["title"] = item.Name
		row["icon"] = item.Icon
		row["href"] = beego.URLFor(item.Controller + "." + item.Action)
		row["type"] = item.Type
		row["status"] = item.Status
		row["list_order"] = item.ListOrder
		if data_type != "All" {
			children, _ := ModelMenuLists.QueryMenuLists(item.Id)
			row["childs"] = getMenuData(children, data_type)
		}
		MenusData[key] = row
	}
	return MenusData
}

//Adicionar e editar dados de operação
func SaveMenu(menu *models.Menu) (err error) {
	ModelMenuAdd = *menu
	if ModelMenuAdd.Id == 0 {
		err = ModelMenuAdd.AddMenuData()
	} else {
		err = ModelMenuAdd.EditMenuData()
	}
	return err
}

//Consultar dados com base nas condições da estrutura
func GetMenuByMap(menu *models.Menu) (data models.Menu, err error) {
	ModelMenuOne = *menu
	return ModelMenuOne, ModelMenuOne.OneMenu()
}

//Exclua os dados com base nas condições da estrutura
func DelMenuByMap(menu *models.Menu) (err error) {
	ModelMenuDel = *menu
	//Adicione uma transação para evitar erros na exclusão
	tx := models.Db.Begin()
	//Divida uma matriz vazia e armazene os dados encontrados em cada ciclo em um loop
	var ChildIds = []int{}
	AllChildIds := getMenuChildId(ChildIds, ModelMenuDel.Id)
	for _, id := range AllChildIds {
		ModelMenuDel.Id = id
		err = ModelMenuDel.DelMenu()
		if err != nil {
			tx.Rollback()
		}
	}
	tx.Commit()
	return nil
}

//Obtém recursivamente todo o próprio menu
func getMenuChildId(ChildIds []int, MenuId int) (NewChildIds []int) {
	NewChildIds = append(ChildIds, MenuId)
	var Child models.Menu
	ChildDatas, _ := Child.QueryMenuLists(MenuId)
	for _, v := range ChildDatas {
		//Depois de escrever seu próprio id, execute uma consulta filho
		NewChildIds = getMenuChildId(NewChildIds, v.Id)
	}
	return NewChildIds
}
