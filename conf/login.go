package conf

// Conecte-se设置:后台Conecte-se相关设置
type Login struct {
	Token      string `mapstructure:"token" json:"token" yaml:"token"`                //Conecte-setoken验证
	Captcha    string `mapstructure:"captcha" json:"captcha" yaml:"captcha"`          //验证码
	Background string `mapstructure:"background" json:"background" yaml:"background"` //Conecte-se背景
}
