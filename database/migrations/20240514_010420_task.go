package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Task_20240514_010420 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Task_20240514_010420{}
	m.Created = "20240514_010420"

	migration.Register("Task_20240514_010420", m)
}

// Run the migrations
func (m *Task_20240514_010420) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Task_20240514_010420) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
