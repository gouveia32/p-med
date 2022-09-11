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

type ClientValidate struct {
	models.Client
}

func (a *ClientValidate) ValidClient() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(a.Client)
	if !b {
		st := reflect.TypeOf(ClientValidate{})
		for _, err := range valid.Errors {
			filed, _ := st.FieldByName(err.Field)
			var alias = filed.Tag.Get("alias")
			return errors.New(alias + err.Message)
		}
	}
	//Verifique se o nome do cliente atual existe
	var client models.Client
	client.Id = a.Id //campo não permitido repetição
	_ = client.OneClient()
	if !reflect.DeepEqual(client, models.Client{}) {
		if client.Id != a.Id {
			return errors.New("Este Nome já está cadastrado!")
		}
	}

	return nil
}
