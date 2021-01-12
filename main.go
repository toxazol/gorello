package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//User type
type User struct {
	Name     string     `json:"name"`
	Projects []*Project `json:"projects"`
}

//Project type
type Project struct {
	Name    string    `json:"name"`
	Columns []*Column `json:"columns"`
}

//Column type
type Column struct {
	Name  string  `json:"name"`
	Tasks []*Task `json:"tasks"`
}

//Task type
type Task struct {
	Name       string     `json:"name"`
	Comments   []*Comment `json:"comments"`
	HostColumn *Column    `json:"-"`
}

//Comment type
type Comment struct {
	Text string `json:"text"`
}

//NewUser func
func NewUser(name string) *User {

	return &User{
		Name:     name,
		Projects: []*Project{NewProject("defaultProjectName")},
	}
}

//NewProject func
func NewProject(name string) *Project {

	return &Project{
		Name:    name,
		Columns: []*Column{NewColumn("defaultColumnName")},
	}
}

//NewColumn func
func NewColumn(name string) *Column {

	return &Column{
		Name:  name,
		Tasks: []*Task{},
	}
}

//NewTask func
func NewTask(name string, columnPtr *Column) *Task {

	return &Task{
		Name:       name,
		Comments:   []*Comment{},
		HostColumn: columnPtr,
	}
}

//NewComment func
func NewComment(text string) *Comment {

	return &Comment{
		Text: text,
	}
}

// AddColumn func
func (p *Project) AddColumn(name string) *Column {
	newCol := NewColumn(name)
	p.Columns = append(p.Columns, newCol)
	return newCol
}

// AddTask func
func (c *Column) AddTask(name string) *Task {
	newT := NewTask(name, c)
	c.Tasks = append(c.Tasks, newT)
	return newT
}

// AddComment func
func (t *Task) AddComment(text string) *Comment {
	newCom := NewComment(text)
	t.Comments = append(t.Comments, newCom)
	return newCom
}

// ChangePriority swaps task with other task in column
func (t *Task) ChangePriority(newPriority int) int {
	taskCount := len(t.HostColumn.Tasks) // panic!
	if newPriority < 0 {
		newPriority = 0
	} else if newPriority >= taskCount {
		newPriority = taskCount - 1
	}
	swapWith := t.HostColumn.Tasks[newPriority]
	*swapWith, *t = *t, *swapWith // panic!
	return newPriority
}

// PrintObj prints user's board contents
func PrintObj(obj interface{}) {
	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
}

func main() {
	me := NewUser("Anton")
	me.Projects[0].AddColumn("TO_DO").AddTask("gorello").AddComment("gambare!")
	me.Projects[0].Columns[1].AddTask("NOT important task")
	me.Projects[0].AddColumn("Done")
	importantTask := me.Projects[0].Columns[1].AddTask("very important task")
	PrintObj(me)
	importantTask.ChangePriority(0)
	PrintObj(me)

}
