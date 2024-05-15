package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"to-do/models"
)

type TaskController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (controller *TaskController) Get() {
	user, err := controller.GetCurrentUser()
	if err != nil {
		beego.Error(err)
		controller.Redirect("/login", 302)
		return
	}

	if user == nil {
		beego.Error("Пользователь не найден")
		controller.Redirect("/login", 302)
		return
	}

	o := orm.NewOrm()
	tasks := make([]*models.Task, 0)
	o.QueryTable("task").Filter("User", user).All(&tasks)
	controller.Data["Tasks"] = tasks

	controller.TplName = "index.tpl"
}

func (controller *TaskController) AddTask() {
	content := controller.GetString("content")
	user, err := controller.GetCurrentUser()
	if err != nil {
		beego.Error(err)
		controller.Redirect("/login", 302)
		return
	}

	if user == nil {
		beego.Error("Пользователь не найден")
		controller.Redirect("/login", 302)
		return
	}

	task := models.Task{
		Content: content,
		Done:    false,
		User:    user,
	}
	o := orm.NewOrm()
	_, err = o.Insert(&task)
	if err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}
	controller.Redirect("/", 302)
}

func (controller *TaskController) DeleteTask() {
	idStr := controller.GetString("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}

	o := orm.NewOrm()
	task := models.Task{Id: id}
	if err := o.Read(&task); err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}

	if _, err := o.Delete(&task); err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}
	controller.Redirect("/", 302)
}

func (controller *TaskController) ToggleTaskDone() {
	idStr := controller.GetString("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}

	o := orm.NewOrm()
	task := models.Task{Id: id}
	if err := o.Read(&task); err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}

	task.Done = !task.Done
	if _, err := o.Update(&task); err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}
	controller.Redirect("/", 302)
}

func (controller *TaskController) GetCurrentUser() (*models.User, error) {
	user := controller.GetSession("user")
	if user == nil {
		return nil, errors.New("Пользователь не найден")
	}

	if u, ok := user.(*models.User); ok {
		return u, nil
	}

	return nil, errors.New("Ошибка преобразования пользователя")
}

func (controller *UserController) Register() {
	username := controller.GetString("username")
	password := controller.GetString("password")

	o := orm.NewOrm()
	existingUser := models.User{Username: username}
	err := o.Read(&existingUser, "Username")
	if err == nil {
		beego.Error("Пользователь уже существует")
		controller.Redirect("/register", 302)
		return
	}

	newUser := models.User{
		Username: username,
		Password: password,
	}
	_, err = o.Insert(&newUser)
	if err != nil {
		beego.Error(err)
		controller.Redirect("/register", 302)
		return
	}

	controller.Redirect("/login", 302)
}

func (controller *UserController) Login() {
	username := controller.GetString("username")
	password := controller.GetString("password")

	o := orm.NewOrm()
	user := models.User{Username: username, Password: password}
	err := o.Read(&user, "Username", "Password")
	if err != nil {
		beego.Error("Неверное имя пользователя или пароль")
		controller.Data["ErrorMessage"] = "Неверное имя пользователя или пароль"
		controller.TplName = "login.tpl"
		return
	}

	controller.SetSession("user", &user)
	controller.Redirect("/tasks", 302)
}

func (controller *UserController) Logout() {
	controller.DelSession("user")
	controller.Redirect("/login", 302)
}

func (controller *TaskController) GetTasks() {
	user, err := controller.GetCurrentUser()
	if err != nil {
		beego.Error(err)
		controller.Redirect("/login", 302)
		return
	}

	if user == nil {
		beego.Error("Пользователь не найден")
		controller.Redirect("/login", 302)
		return
	}

	o := orm.NewOrm()
	tasks := make([]*models.Task, 0)
	o.QueryTable("task").Filter("User", user).All(&tasks)
	controller.Data["Tasks"] = tasks
	controller.Data["User"] = user
	controller.TplName = "tasks.tpl"
}
