package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"strconv"
	"studentmanagement/models"
)
import _ "studentmanagement/models"

type StudentController struct {
	beego.Controller
}

type Params struct {
	Sno   int
	Sname string
	Ssex  string
	Sage  int
	Sdept string
}

func (c *StudentController) Get() {
	//	如果参数为空，返回所有学生信息
	p, err := c.ValidParams()

	if err != nil {
		c.Data["json"] = "参数错误"
		c.ServeJSON()
		return
	}

	s := &models.Student{}
	if p.Sno == 0 && p.Sname == "" {
		fmt.Println("查询所有学生信息")
		students, err := s.QueryAll()
		if err != nil {
			c.Data["json"] = "查询失败"
			c.ServeJSON()
			return
		}
		c.Data["json"] = students
		c.ServeJSON()
	} else {
		students, err := s.QueryByParams(p.Sno, p.Sname)
		if err != nil {
			c.Data["json"] = "查询失败"
			c.ServeJSON()
			return
		}
		c.Data["json"] = students
		c.ServeJSON()
	}
}

// Add 添加学生信息
func (c *StudentController) Add() {
	var p, err = c.ValidParams()
	logs.Info(p)
	if err != nil {
		c.Data["json"] = "参数错误"
		c.ServeJSON()
		return
	}
	s := &models.Student{
		Sno:   p.Sno,
		Sname: p.Sname,
		Ssex:  p.Ssex,
		Sage:  p.Sage,
		Sdept: p.Sdept,
	}
	err = s.Add()
	if err != nil {
		c.Data["json"] = "添加失败" + err.Error()
		logs.Error("添加失败", err)
		c.ServeJSON()
		return
	}
	c.Data["json"] = "添加成功"
	c.ServeJSON()
}

func (c *StudentController) Search() {
	p, err := c.ValidParams()
	if err != nil {
		c.Data["json"] = "参数错误"
		c.ServeJSON()
		return
	}
	if p.Sno == 0 && p.Sname == "" {
		//	如果参数为空，返回Get方法
		c.Get()
	}
	s := &models.Student{
		Sno:   p.Sno,
		Sname: p.Sname,
		Ssex:  p.Ssex,
		Sage:  p.Sage,
		Sdept: p.Sdept,
	}
	students, err := s.QueryByParams(p.Sno, p.Sname)
	if err != nil {
		c.Data["json"] = "查询失败"
		c.ServeJSON()
		return
	}
	c.Data["json"] = students
	c.ServeJSON()
}

// Delete 删除学生信息
func (c *StudentController) Delete() {
	p, err := c.ValidParams()
	if err != nil {
		c.Data["json"] = "参数错误"
		c.ServeJSON()
		return
	}
	s := &models.Student{
		Sno: p.Sno,
	}
	err = s.Delete()
	if err != nil {
		c.Data["json"] = "删除失败"
		c.ServeJSON()
		return
	}
	c.Data["json"] = "删除成功"
	c.ServeJSON()
}

// Update 更新学生信息
func (c *StudentController) Update() {
	logs.Info("更新学生信息")
	p, err := c.ValidParams()
	if err != nil {
		c.Data["json"] = "参数错误"
		c.ServeJSON()
		return
	}
	s := &models.Student{
		Sno:   p.Sno,
		Sname: p.Sname,
		Ssex:  p.Ssex,
		Sage:  p.Sage,
		Sdept: p.Sdept,
	}
	err = s.Update()
	if err != nil {
		c.Data["json"] = "更新失败"
		c.ServeJSON()
		return
	}
	c.Data["json"] = "更新成功"
	c.ServeJSON()
}

// AdvancedSearch 高级查询
func (c *StudentController) AdvancedSearch() {
	var sage int
	var sageCondition string
	var err error
	p_sage := c.GetString("sage")
	if p_sage == "" {
		sage = 0
		sageCondition = ">"
	} else {
		sageCondition = c.GetString("sageCondition")
		sage, err = c.GetInt("sage")
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
			c.Data["json"] = "参数错误"
			c.ServeJSON()
			return
		}
	}
	sdept := c.GetString("sdept")
	ssex := c.GetString("ssex")

	s := &models.Student{}
	students, err := s.AdvancedSearch(sage, sdept, ssex, sageCondition)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
		c.Data["json"] = "查询失败"
		c.ServeJSON()
		return
	}
	c.Data["json"] = students
	c.ServeJSON()
}

// ValidParams 验证参数
func (c *StudentController) ValidParams() (p *Params, err error) {
	p = &Params{0, "", "", 0, ""}
	p.Sname = c.GetString("sname")
	p.Ssex = c.GetString("ssex")
	p.Sdept = c.GetString("sdept")
	p.Sage, err = c.GetInt("sage", 0)
	if err != nil {
		return p, err
	}
	p.Sno, err = c.GetInt("sno", 0)
	if err != nil {
		return p, err
	}
	// 查看search是否为空
	search := c.GetString("search")
	if search != "" {
		p.Sno, err = strconv.Atoi(search)
		if err != nil {
			p.Sno = -1
			p.Sname = search
			return p, nil
		}
		p.Sname = search
	}
	return
}
