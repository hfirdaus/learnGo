package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"time"
)

// Reference: https://www.codementor.io/engineerapart/getting-started-with-postgresql-on-mac-osx-are8jcopb

func init() {
	db := openTodoDb()
	defer db.Close()

	if !db.HasTable(&Todo{}) {
		db.AutoMigrate(&Todo{})
		InsertTodo(Todo{Task: "Thing 1", Completed: false, DueDate: time.Now(), User: "TestUser"})
		InsertTodo(Todo{Task: "Thing 2", Completed: true, DueDate: time.Now(), User: "TestUser"})
	}
}

func openTodoDb() *gorm.DB {
//	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	db, err := gorm.Open("postgres", "host=tantor.db.elephantsql.com port=5432 user=yvrvqewb dbname=yvrvqewb password=sf95rtFPhXxvO8-Ag9T-fXJhZdRbRWL3 sslmode=disable")
//	url, err := pq.ParseURL("postgres://yvrvqewb:0AuSTPldgCKiL9m8sa8VSP46Et65noQP@tantor.db.elephantsql.com:5432/yvrvqew?sslmode=disable")
//	db, err := gorm.Open("postgres", url)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func InsertTodo(t Todo) Todo {
	db := openTodoDb()
	defer db.Close()
	t.ID = time.Now().Nanosecond()
	db.Create(&t)
	return t
}

func DeleteTodoById(Id int) {
	var todo Todo
	db := openTodoDb()
	defer db.Close()
	db.First(&todo, Id)
	db.Delete(todo)
}

func ToggleTodoCompletedValue(Id int) {
	var todo Todo
	db := openTodoDb()
	defer db.Close()
	db.First(&todo, Id)
	db.Model(todo).Update("Completed", !todo.Completed)
}

func UpdateTodo(t Todo) {
	db := openTodoDb()
	defer db.Close()
	db.Model(t).Update("Task", t.Task)
	db.Model(t).Update("DueDate", t.DueDate)
}

func FindTodoById(Id int) Todo {
	var todo Todo
	db := openTodoDb()
	defer db.Close()
	db.First(todo, Id)
	return todo
}

func FindAllTodos() []Todo {
	var todos []Todo
	db := openTodoDb()
	defer db.Close()
	db.Find(&todos)
	return todos
}
