package main

import (
	"encoding/json"
	"fmt"
	"go-backend/models"
	"log"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, res *http.Request) {
	fmt.Fprintln(w, "<h1>it works")
}

func createTask(w http.ResponseWriter, res *http.Request) {
	var task models.Task
	json.NewDecoder(res.Body).Decode(&task)
	models.Create(task)
}

func updateTask(w http.ResponseWriter, res *http.Request) {
	var task models.Task
	json.NewDecoder(res.Body).Decode(&task)
	models.Update(task)
}

func deleteTask(w http.ResponseWriter, res *http.Request) {
	muxVar := mux.Vars(res)
	id, err := strconv.Atoi(muxVar["id"])
	if err != nil {
		log.Panic(err)
	}
	models.Delete(id)
}

func tasks(w http.ResponseWriter, res *http.Request) {
	tasks := models.List()
	json.NewEncoder(w).Encode(tasks)
}
