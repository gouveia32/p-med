package init

import (
	"pm/models"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Concordar em registrar a configuração do banco de dados
func InitDatabase() {
	//Inicializar informações de configuração
	host := beego.AppConfig.DefaultString("::db.host", "127.0.0.1")
	user := beego.AppConfig.DefaultString("mysql::db.user", "root")
	password := beego.AppConfig.DefaultString("mysql::db.password", "")
	port := beego.AppConfig.DefaultString("mysql::db.port", "3306")
	dbname := beego.AppConfig.DefaultString("mysql::db.name", "ccb_bee")
	dsn := user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	models.Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Falha na conexão com o banco de dados:" + err.Error())
	}
	models.Db.SingularTable(true)
	//Defina o nome da tabela padrão
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		prefix := beego.AppConfig.DefaultString("db.prefix", "")
		return prefix + defaultTableName
	}

	//Se deve ativar o modo de depuração por padrão
	if beego.AppConfig.String("runmode") == "dev" {
		models.Db.LogMode(true)
	}
}
