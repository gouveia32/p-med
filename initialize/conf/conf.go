package conf

import (
	"fmt"
	"os"
	"p-med/global"
	"p-med/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//加载配置文件
func init() {
	var config string
	if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
		config = utils.ConfigFile
		fmt.Printf("Você está usando o valor padrão de config, %v\n", utils.ConfigFile)
	} else {
		config = configEnv
		fmt.Printf("Você está usando a variável de ambiente BA_CONFIG,%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.BA_CONFIG); err != nil {
			panic(err)
		}
	})

	if err := v.Unmarshal(&global.BA_CONFIG); err != nil {
		panic(err)
	}

}
