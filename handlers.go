package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// PrintObj prints user's board contents
func respond(resp http.ResponseWriter, status int, obj interface{}) {
	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(status)
	resp.Write(jsonBytes)
}

//GetProjects handles /columns get request
func GetProjects(resp http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	projects, err := repo.ReadProjects()
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, projects)
}

//GetColumns handles /columns get request
func GetColumns(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	projID, err := strconv.Atoi(p.ByName("project_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid project ID, it must be a number",
			p.ByName("project_id")), http.StatusBadRequest)
		return
	}
	columns, err := repo.ReadColumns(projID)
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, columns)
}

func GetColumn(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	colID, err := strconv.Atoi(p.ByName("column_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid column ID, it must be a number",
			p.ByName("column_id")), http.StatusBadRequest)
		return
	}
	column, err := repo.ReadColumn(colID)
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, column)
}
func GetTasks(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	colID, err := strconv.Atoi(p.ByName("column_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid column ID, it must be a number",
			p.ByName("column_id")), http.StatusBadRequest)
		return
	}
	tasks, err := repo.ReadTasks(colID)
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, tasks)
}
func GetTask(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	taskID, err := strconv.Atoi(p.ByName("task_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid task ID, it must be a number",
			p.ByName("task_id")), http.StatusBadRequest)
		return
	}
	task, err := repo.ReadTask(taskID)
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, task)
}
func GetComments(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	taskID, err := strconv.Atoi(p.ByName("task_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid task ID, it must be a number",
			p.ByName("task_id")), http.StatusBadRequest)
		return
	}
	comments, err := repo.ReadComments(taskID)
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, comments)
}
func GetComment(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	commentID, err := strconv.Atoi(p.ByName("comment_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid task ID, it must be a number",
			p.ByName("comment_id")), http.StatusBadRequest)
		return
	}
	comment, err := repo.ReadComments(commentID)
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, comment)
}
