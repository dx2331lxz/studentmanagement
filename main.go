package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"studentmanagement/models"
	_ "studentmanagement/models"
	_ "studentmanagement/routers"
	"time"
)

func init() {
	c := models.NewConfig()
	err := models.InitDB(c)
	if err != nil {
		logs.Error("连接数据库失败：", err)
	}
}

func main() {
	serviceLogFile := "log/" + time.Now().Format("2006-01-02") + ".log"
	logs.SetLogger(logs.AdapterFile, `{
        "filename": "`+serviceLogFile+`",
        "daily": false,
        "maxlines": 1000,
        "maxsize": 0,
        "level": 7, 
        "perm": "0660"
    }`)
	logs.Info(serviceLogFile)
	logs.SetLogFuncCallDepth(3)
	logs.SetLogFuncCall(true) // 记录文件名和行号
	logs.Info("启动日志：项目启动成功")  // 启动时写入一条启动日志
	beego.Run()
}
