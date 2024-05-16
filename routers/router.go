package routers

import (
	"github.com/astaxie/beego"
	"to-do/controllers"
)

func init() {
	beego.Router("/", &controllers.TaskController{})
	beego.Router("/addtask", &controllers.TaskController{}, "post:AddTask")
	beego.Router("/deletetask", &controllers.TaskController{}, "get:DeleteTask")
	beego.Router("/toggletaskdone", &controllers.TaskController{}, "get:ToggleTaskDone")

	beego.Router("/register", &controllers.UserController{}, "get,post:Register")
	beego.Router("/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/tasks", &controllers.TaskController{}, "get:GetTasks")

}
