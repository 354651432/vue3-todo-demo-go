package models

import (
	"log"
	_ "log"
)

type Task struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Status   int    `json:"status"`
	CreateAt string `json:"create_at"`
}

var db = GetDb()

func Create(task Task) {
	log.Printf("create %v", task)
	db.Exec("insert into tasks(name,status) values(?,?)", task.Name, task.Status)
}

func Update(task Task) {
	db.Exec("update tasks set status=? where id=?", task.Status, task.Id)
}

func Delete(id int) {
	db.Exec("delete from tasks where id=?", id)
}

func List() (tasks []Task) {
	rows, err := db.Query("select id,name,status,create_at from tasks")
	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.Name, &task.Status, &task.CreateAt)
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		log.Panic(err)
	}

	defer func() {
		if tasks == nil {
			tasks = []Task{}
		}
	}()
	
	return
}
