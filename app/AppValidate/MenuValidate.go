/**
* @author:CCB_F
* @introduce: 后台菜单的添加验证
* @time:2019/8/27
 */
package AppValidate

import (
	"errors"
	"reflect"

	"pm/models"

	"github.com/astaxie/beego/validation"
)

type MenuValidate struct {
	models.Menu
}

func (a *MenuValidate) ValidMenu() (err error) {
	//println("Estou aqui: MenuValidate1", a.Id)
	valid := validation.Validation{}
	b, _ := valid.Valid(a.Menu)
	if !b {
		st := reflect.TypeOf(MenuValidate{})
		for _, err := range valid.Errors {
			filed, _ := st.FieldByName(err.Field)
			var alias = filed.Tag.Get("alias")
			return errors.New(alias + err.Message)
		}
	}
	//Verifique se o nome do artigo atual existe
	var menu models.Menu
	menu.Controller = a.Controller
	menu.Action = a.Action
	menu.Param = a.Param
	menu.ParentId = a.ParentId
	_ = menu.OneMenu()
	if !reflect.DeepEqual(menu, models.Menu{}) {
		if menu.Id != a.Id {
			return errors.New("O controlador/método existem, por favor, preencha novamente")
		}
	}

	return nil
}
