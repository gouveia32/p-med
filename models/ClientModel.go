/**
@author : CCB_f
@后台功能模板
*/
package models

import (
	"strings"
)

type Client struct {
	Id            int64   `form:"id"        		alias:"ID" `
	Nome          string  `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(120)"`
	ContatoFuncao string  `form:"contato_funcao"	alias:"ContatoFuncao "	valid:"MaxSize(20)"`
	ContatoNome   string  `form:"contato_nome"    	alias:"ContatoNome"  	valid:"MaxSize(40)"`
	CgcCpf        string  `form:"cgc_cpf"    		alias:"CgcCpf"  		valid:"MaxSize(20)"`
	RazaoSocial   string  `form:"razao_social"    	alias:"RazaoSocial"  	valid:"MaxSize(100)"`
	InscrEstadual string  `form:"inscr_estadual"   	alias:"InscrEstadual"  	valid:"MaxSize(20)"`
	Endereco      string  `form:"endereco"    		alias:"Endereco"  		valid:"MaxSize(100)"`
	Cidade        string  `form:"cidade"    		alias:"Cidade"  		valid:"MaxSize(40)"`
	Uf            string  `form:"uf"    			alias:"Uf"  			valid:"MaxSize(2)"`
	Cep           string  `form:"cep"    			alias:"Cep"  			valid:"MaxSize(9)"`
	Telefone1     string  `form:"telefone1"    		alias:"Telefone1"  		valid:"MaxSize(20)"`
	Telefone2     string  `form:"telefone2"    		alias:"Telefone2"  		valid:"MaxSize(20)"`
	Telefone3     string  `form:"telefone3"    		alias:"Telefone3"  		valid:"MaxSize(20)"`
	Email         string  `form:"email"    			alias:"email"  			valid:"MaxSize(60)"`
	Obs           string  `form:"obs"    			alias:"Obs"  			valid:"MaxSize(10240)"`
	Estado        int8    `form:"estado"    		alias:"Estado"`
	PrecoBase     float64 `form:"preco_base"    	alias:"PrecoBase"  		orm:"digits(12);decimals(2)"`
}

func (m *Client) TableName() string {
	return "clientes"
}

//Consulte as informações do usuário de acordo com as condições da estrutura, passe os campos que precisam ser consultados
func (m *Client) QueryClientLists(field ...interface{}) (client []Client, err error) {
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
	err = Db.Model(Client{}).Where(&m).
		Order("list_order asc").
		Select(fieldStr).
		Find(&client).Error
	return
}

//Adicionar
func (m *Client) AddClientData() (err error) {
	Db.LogMode(true)
	return Db.Model(Client{}).Create(&m).Error
}

//Faça edições
func (m *Client) EditClientData() (err error) {
	Db.LogMode(true)
	return Db.Model(Client{}).Save(&m).Error
}

//Obtenha todos os menus
func (m *Client) AllClient(offset int, limit int) (client []Client, err error) {
	return client, Db.Model(Client{}).Where(&m).Order("id asc").Offset(offset).Limit(limit).Find(&client).Error
}

//Obtenha apenas um cliente
func (m *Client) OneClient() (err error) {
	return Db.Model(Client{}).Where(&m).First(&m).Error
}

func (m *Client) DelClient() (err error) {
	return Db.Model(Client{}).Delete(&m).Error
}
