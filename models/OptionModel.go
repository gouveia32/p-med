/**
* @author:CCB_F
* @introduce: Camada de modelo de tabela de configuração em todo o site
* @time:2019/8/29
 */

package models

//Obtenha informações de opções
type Option struct {
	Id            int
	Autoload      int    `alias:"Para carregar automaticamente"`
	OptionName    string `alias:"Nome da configuração"`
	OptionValue   string `alias:"Valor de configuração"`
	OptionComment string `alias:"Comentários"`
}

func (m *Option) OneOption() (err error) {
	Db.LogMode(true)
	return Db.Model(Option{}).Where(&m).First(&m).Error
}

//Adicionar
func (m *Option) AddOptionData() (err error) {
	Db.LogMode(true)
	return Db.Model(Option{}).Create(&m).Error
}

//Faça edições
func (m *Option) EditOptionData() (err error) {
	Db.LogMode(true)
	return Db.Model(Option{}).Save(&m).Error
}
