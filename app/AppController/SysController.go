/**
* @author:CCB_F
* @introduce: 系统相关配置控制器
* @time:2019/8/28
 */

package AppController

import (
	"pm/app/AppService"

	"github.com/astaxie/beego"
)

type SysController struct {
	AppBaseController
}

//网站相关配置修改
func (c *SysController) Config() {
	var site AppService.SiteConfig
	if c.IsHtml() {
		//读取option配置
		var site AppService.SiteConfig
		_ = site.GetSiteConfig()
		c.Data["data"] = site
		c.Display()
	} else {
		if err := c.ParseForm(&site); err != nil {
			c.Abort500("Erro de parâmetro de entrada:"+err.Error(), "")
		}
		err := site.AddSiteConfig()
		if err != nil {
			c.Abort500(err.Error(), "")
		}
		c.Abort200("", "Adicionado com sucesso", "")
	}
}

//获取icon
func (c *SysController) Icon() {
	c.Display()
}

//Interface de upload de imagens
func (c *SysController) UpImg() {
	//Formato restrito de imagens enviadas
	AllowImgExt := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	PIC_PATH := beego.AppConfig.String("pic_path")
	fpath, err := c.SaveFile(PIC_PATH, AllowImgExt)
	if err != nil {
		c.Abort500(err.Error(), "")
		return
	}
	c.Abort200(fpath, "Carregado com sucesso", "")
}
