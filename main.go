package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name     string    `json:"name"`
	Projects []Project `json:"projects"`
}

func NewUser(name string) User {

	return User{
		Name:     name,
		Projects: []Project{NewProject("defaultProjectName")},
	}
}

type Project struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

func NewProject(name string) Project {

	return Project{
		Name:    name,
		Columns: []Column{NewColumn("defaultColumnName")},
	}
}

type Column struct {
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

func NewColumn(name string) Column {

	return Column{
		Name:  name,
		Tasks: []Task{NewTask("defaultTaskName")},
	}
}

type Task struct {
	Name     string    `json:"name"`
	Comments []Comment `json:"comments"`
}

func NewTask(name string) Task {

	return Task{
		Name:     name,
		Comments: []Comment{NewComment("defaultCommentText")},
	}
}

type Comment struct {
	Text string `json:"text"`
}

func NewComment(text string) Comment {

	return Comment{
		Text: text,
	}
}

func main() {
	me := NewUser("Anton")
	jsonBytes, err := json.MarshalIndent(me, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
}
