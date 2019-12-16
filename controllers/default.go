package controllers

import (
	"github.com/astaxie/beego"
	"beegoapp/clockin/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	m := models.GetPageInfo()
	c.Data["Name"] = m.Name
	c.Data["Email"] = m.Email
	c.TplName = "index.tpl"
}
