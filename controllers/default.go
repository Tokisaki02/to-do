package controllers

import (
	"errors"
	"fmt"
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
	_, err = o.QueryTable("task").Filter("User", user).All(&tasks)
	if err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}
	controller.Data["Tasks"] = tasks

	controller.TplName = "tasks.tpl"
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
	err = o.Read(&task)
	if err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}

	_, err = o.Delete(&task)
	if err != nil {
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
	err = o.Read(&task)
	if err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}

	task.Done = !task.Done
	_, err = o.Update(&task)
	if err != nil {
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
		// Проверка типа данных
		fmt.Printf("Тип данных пользователя: %T\n", u)
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
		controller.Data["ErrorRegister"] = "Пользователь уже существует"
		controller.TplName = "index.tpl"
		return
	}

	newUser := models.User{
		Username: username,
		Password: password,
	}
	_, err = o.Insert(&newUser)
	if err != nil {
		beego.Error(err)
		controller.TplName = "index.tpl"
		return
	}

	controller.Redirect("/login", 302)
}

func (controller *UserController) Login() {
	username := controller.GetString("username")
	password := controller.GetString("password")

	// Создаем объект пользователя вручную
	user := &models.User{
		Username: username,
		Password: password,
	}

	// Проверяем валидность пользователя
	if isValidUser(user) {
		// Устанавливаем сессию пользователя
		if controller.GetSession("user") != nil {
			beego.Error("Сессия уже установлена")
		} else {
			controller.SetSession("user", user)
			beego.Debug("Сессия установлена")
		}

		controller.Redirect("/tasks", 302)
	} else {
		beego.Error("Неверное имя пользователя или пароль")
		controller.Data["ErrorMessage"] = "Неверное имя пользователя или пароль"
		controller.TplName = "login.tpl"
	}
}

// Функция для проверки валидности пользователя
func isValidUser(user *models.User) bool {
	// Здесь можно добавить логику проверки пользователя, например, сравнение с данными в базе данных или другие проверки
	// В данном примере просто возвращаем true для демонстрации
	return true
}

func (controller *UserController) Logout() {
	// Удаляем сессию пользователя
	controller.DelSession("user")

	// Проверяем, что сессия удалена
	if controller.GetSession("user") == nil {
		beego.Debug("Сессия удалена корректно")
	} else {
		beego.Error("Ошибка удаления сессии")
	}

	// Перенаправляем на страницу входа
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
	_, err = o.QueryTable("task").Filter("User", user).All(&tasks)
	if err != nil {
		beego.Error(err)
		controller.Redirect("/", 302)
		return
	}
	controller.Data["Tasks"] = tasks
	controller.Data["User"] = user
	controller.TplName = "tasks.tpl"
}
