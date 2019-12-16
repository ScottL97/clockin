package controllers

import (
	"github.com/astaxie/beego"
	"beegoapp/clockin/models"
	"github.com/skip2/go-qrcode"
)

type ClockinController struct {
	beego.Controller	
} 
type LoginController struct {
	beego.Controller
}
type ResultController struct {
	beego.Controller
}

func (c *ClockinController) Get() {
	qrcode.WriteFile("http://starguard.cn:8080/clockin/login",qrcode.Medium,256,"../static/img/qrcode.png")
	m := models.GetQrcodeInfo()
	c.Data["Imgurl"] = m.Url
	c.TplName = "clockin.tpl"
} 

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *ResultController) Post() {
	Name := c.GetString("Name")
	res := models.InsertRec(Name)
	switch res {
		case 1:
			c.Data["Info"] = "签到成功：" + Name
		case 2:
			c.Data["Info"] = "请勿重复签到"
		default:
			c.Data["Info"] = "签到失败"
	}
	UserInfo := models.GetUserInfo()
	users := ""
	for _, user := range UserInfo {
		users = users + "，" + user.Name
	}
	users = users[3:]
	c.Data["Users"] = users
	c.TplName = "result.tpl"
}
