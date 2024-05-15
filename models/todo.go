package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int    `orm:"column(id)"`
	Username string `orm:"column(username)"`
	Password string `orm:"column(password)"`
}

type Task struct {
	Id      int    `orm:"auto"`
	Content string `orm:"size(255)"`
	Done    bool
	User    *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(User), new(Task))
}

func (u *User) AddTask(content string) (*Task, error) {
	task := &Task{Content: content, User: u}
	_, err := orm.NewOrm().Insert(task)
	return task, err
}

func (u *User) GetTasks() ([]*Task, error) {
	tasks := make([]*Task, 0)
	_, err := orm.NewOrm().QueryTable("task").Filter("user_id", u.Id).All(&tasks)
	return tasks, err
}

func (u *User) DeleteTask(id int) error {
	_, err := orm.NewOrm().QueryTable("task").Filter("id", id).Filter("user_id", u.Id).Delete()
	return err
}
