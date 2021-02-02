package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// respond responds with json
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

//GetColumn func
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

//GetTasks func
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

//GetTask func
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

//GetComments func
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

//GetComment func
func GetComment(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	commentID, err := strconv.Atoi(p.ByName("comment_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid comment ID, it must be a number",
			p.ByName("comment_id")), http.StatusBadRequest)
		return
	}
	comment, err := repo.ReadComments(commentID)
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, comment)
}

//CreateProject func
func CreateProject(resp http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var newProject Project
	err := decoder.Decode(&newProject)
	if err != nil {
		log.Fatal("CreateProject decode error: ", err)
	}
	savedProject, err := repo.SaveProject(newProject)
	newColumn := Column{Name: "New Column", ProjectID: savedProject.ID} //TO_DO: put defaults in a config
	if err != nil {
		log.Fatal(err)
	}
	_, err = repo.SaveColumn(newColumn) //TO_DO: it should be transaction
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, savedProject) // TO_DO: status not ok ??
}

//CreateColumn func
func CreateColumn(resp http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var newColumn Column
	err := decoder.Decode(&newColumn)
	if err != nil {
		log.Fatal("CreateColumn decode error: ", err)
	}
	// TO_DO: add independent verifier
	isUniq, err := repo.CheckColNameUniq(newColumn.ProjectID, newColumn.Name)
	if err != nil {
		log.Fatal(err) //TO_DO: rethink fatality
	}
	if !isUniq {
		http.Error(resp, "duplicate column name", http.StatusNotAcceptable)
		return
	}
	savedColumn, err := repo.SaveColumn(newColumn)
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, savedColumn) // TO_DO: status not ok ??
}

//CreateTask func
func CreateTask(resp http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var newTask Task
	err := decoder.Decode(&newTask)
	if err != nil {
		log.Fatal("CreateTask decode error: ", err)
	}
	savedTask, err := repo.SaveTask(newTask)
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, savedTask) // TO_DO: status not ok ??
}

//CreateComment func
func CreateComment(resp http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var newComment Comment
	err := decoder.Decode(&newComment)
	if err != nil {
		log.Fatal("CreateComment decode error: ", err)
	}
	savedComment, err := repo.SaveComment(newComment)
	if err != nil {
		log.Fatal(err)
	}
	respond(resp, http.StatusOK, savedComment) // TO_DO: status not ok ??
}

//DeleteProject func
func DeleteProject(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	projectID, err := strconv.Atoi(p.ByName("project_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid project ID, it must be a number",
			p.ByName("project_id")), http.StatusBadRequest)
		return
	}
	err = repo.RemoveProject(projectID)
	if err != nil {
		log.Fatal(err)
	}
	resp.WriteHeader(http.StatusOK) // TO_DO: not ok status?
}

//DeleteColumn func
func DeleteColumn(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	columnID, err := strconv.Atoi(p.ByName("column_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid column ID, it must be a number",
			p.ByName("column_id")), http.StatusBadRequest)
		return
	}
	colToDel, err := repo.ReadColumn(columnID)
	if err != nil {
		log.Fatal(err)
	}
	columns, err := repo.ReadColumns(colToDel.ProjectID)
	if err != nil {
		log.Fatal(err)
	}
	if len(columns) > 1 {
		err = repo.RemoveColumn(columnID)
		if err != nil {
			log.Fatal(err)
		}
		resp.WriteHeader(http.StatusOK)
		return
	}
	resp.WriteHeader(http.StatusNotModified)
}

//DeleteTask func
func DeleteTask(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	taskID, err := strconv.Atoi(p.ByName("task_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid task ID, it must be a number",
			p.ByName("task_id")), http.StatusBadRequest)
		return
	}
	err = repo.RemoveTask(taskID)
	if err != nil {
		log.Fatal(err)
	}
	resp.WriteHeader(http.StatusOK)
}

//DeleteComment func
func DeleteComment(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	commentID, err := strconv.Atoi(p.ByName("comment_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid comment ID, it must be a number",
			p.ByName("comment_id")), http.StatusBadRequest)
		return
	}
	err = repo.RemoveComment(commentID)
	if err != nil {
		log.Fatal(err)
	}
	resp.WriteHeader(http.StatusOK)
}

func getNewColPriority(columns []Column, curPriority float64, newPosition int) float64 {
	colCount := len(columns)
	if colCount <= 1 {
		return curPriority
	}
	newPosition = getBoundIndex(newPosition, colCount)
	newPriorityFrom := columns[newPosition].Priority
	var newPriorityTo float64

	if newPosition == 0 {
		newPriorityTo = newPriorityFrom - 1
	} else if newPosition == colCount-1 {
		newPriorityTo = newPriorityFrom + 1
	} else if newPriorityFrom > curPriority {
		newPriorityTo = columns[newPosition+1].Priority
	} else if newPriorityFrom < curPriority {
		newPriorityTo = columns[newPosition-1].Priority
	}
	return getBetweenPriority(newPriorityFrom, newPriorityTo)

}

//MoveColumn func
func MoveColumn(resp http.ResponseWriter, r *http.Request, p httprouter.Params) {
	columnID, err := strconv.Atoi(p.ByName("column_id"))
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid column ID, it must be a number",
			p.ByName("column_id")), http.StatusBadRequest)
		return
	}
	newPos, err := strconv.Atoi(r.URL.Query().Get("new_pos")) // TO_DO: validate!
	if err != nil {
		http.Error(resp, fmt.Sprintf("%s is not a valid position, it must be a number",
			r.URL.Query().Get("new_pos")), http.StatusBadRequest)
		return
	}
	column, err := repo.ReadColumn(columnID)
	if err != nil {
		log.Fatal(err)
	}
	columns, err := repo.ReadColumns(column.ProjectID)
	if err != nil {
		log.Fatal(err)
	}
	column.Priority = getNewColPriority(columns, column.Priority, newPos)
	_, err = repo.SaveColumn(column) //TO_DO: it should be transaction
	if err != nil {
		log.Fatal(err)
	}
	resp.WriteHeader(http.StatusOK)
	//log.Printf("columnID: %d, newPos: %s", columnID, newPos)

}
