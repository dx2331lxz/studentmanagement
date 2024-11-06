package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"studentmanagement/models"
)

type MajorController struct {
	beego.Controller
}

func (c *MajorController) Get() {
	major := models.Major{}
	majors, err := major.QueryAll()
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		c.Data["json"] = "查询失败"
		c.ServeJSON()
		return
	}
	c.Data["json"] = majors
	c.ServeJSON()
}
