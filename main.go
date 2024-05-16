package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"to-do/controllers"
	_ "to-do/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	SQLlink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", beego.AppConfig.String("dbUser"), beego.AppConfig.String("dbPass"), beego.AppConfig.String("dbHost"), beego.AppConfig.String("dbPort"), beego.AppConfig.String("dbName"))
	orm.RegisterDataBase("default", "mysql", SQLlink)

	orm.Debug = true
}

func main() {
	orm.RunCommand()
	beego.Router("/", &controllers.TaskController{})
	beego.Router("/addtask", &controllers.TaskController{}, "post:AddTask")
	beego.Router("/deletetask", &controllers.TaskController{}, "get:DeleteTask")
	beego.Router("/toggletaskdone", &controllers.TaskController{}, "get:ToggleTaskDone")

	beego.Router("/register", &controllers.UserController{}, "get,post:Register")
	beego.Router("/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/tasks", &controllers.TaskController{}, "get:GetTasks")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
