/**
@author : CCB_f
@后台功能模板
*/
package models

import (
	"database/sql"
	"strings"
)

type Menu struct {
	Id           int    `form:"id"          alias:"ID" `
	ParentId     int    `form:"parent_id"  alias:"nível superior"  `
	Type         int    `form:"type"       alias:"Tipo"   `
	Status       int    `form:"status"     alias:"Estado"      `
	ListOrder    int    `form:"list_order" alias:"Ordenação"       valid:"Max(100);NonNegativeInteger"`
	Controller   string `form:"controller" alias:"Nome do controlador"    valid:"Required;MaxSize(30)"`
	Action       string `form:"action"     alias:"Nome da Operação"    valid:"Required;MaxSize(30)"`
	Param        string `form:"param"      alias:"Parâmetros extras"    valid:"MaxSize(30)"`
	Name         string `form:"name"       alias:"Nome"    valid:"Required;MaxSize(30)"`
	Icon         string `form:"icon"       alias:"Ícone do menu"    valid:"MaxSize(50)"`
	Remark       string `form:"remark"     alias:"Observação"        valid:"MaxSize(250)"`
	ModuleBelong int8   `form:"module_belong"   alias:"Módulo de leitura" `
	//Como atribuição, este campo não existe no banco de dados
	NullParentId sql.NullInt64 `gorm:"-" form:"parent_id"`
}

//根据struct条件查询用户信息  传入需要查询的字段
func (m *Menu) QueryMenuLists(parent_id int, field ...interface{}) (menu []Menu, err error) {
	//Db.LogMode(true)
	var fieldStr string
	if len(field) > 0 {
		for k, _ := range field {
			fieldStr = fieldStr + "," + field[k].(string)
		}
	} else {
		fieldStr = " * "
	}
	fieldStr = strings.Trim(fieldStr, ",")
	err = Db.Model(Menu{}).Where(&m).Where("parent_id = ?", parent_id).
		Order("list_order asc").
		Select(fieldStr).
		Find(&menu).Error
	return
}

//Adicionar
func (m *Menu) AddMenuData() (err error) {
	Db.LogMode(true)
	return Db.Model(Menu{}).Create(&m).Error
}

//Faça edições
func (m *Menu) EditMenuData() (err error) {
	Db.LogMode(true)
	return Db.Model(Menu{}).Save(&m).Error
}

//Obtenha todos os menus
func (m *Menu) AllMenu() (menu []Menu, err error) {
	return menu, Db.Model(Menu{}).Where(&m).Order("list_order asc").Find(&menu).Error
}

//Obtenha todos os menus
func (m *Menu) OneMenu() (err error) {
	return Db.Model(Menu{}).Where(&m).First(&m).Error
}

func (m *Menu) DelMenu() (err error) {
	return Db.Model(Menu{}).Delete(&m).Error
}
