package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"studentmanagement/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/student", &controllers.StudentController{})
	beego.Router("/student/add", &controllers.StudentController{}, "post:Add")
	beego.Router("/student/search", &controllers.StudentController{}, "post:Search")
	beego.Router("/student/delete", &controllers.StudentController{}, "post:Delete")
	beego.Router("/student/update", &controllers.StudentController{}, "post:Update")
	beego.Router("/student/advancedSearch", &controllers.StudentController{}, "post:AdvancedSearch")

	beego.Router("/major", &controllers.MajorController{})
}
