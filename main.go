package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

//Project type
type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Columns     []*Column `json:"columns"`
}

//Column type
type Column struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Tasks []*Task `json:"tasks"`
}

//Task type
type Task struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Priority    float64    `json:"-"`
	Comments    []*Comment `json:"comments"`
	HostColumn  *Column    `json:"-"`
}

//Comment type
type Comment struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

//NewProject func
func NewProject(name, descr string) *Project {

	return &Project{
		Name:        name,
		Description: descr,
		Columns:     []*Column{NewColumn("defaultColumnName")},
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
func NewTask(name, descr string, pos float64, columnPtr *Column) *Task {

	return &Task{
		Name:        name,
		Description: descr,
		Priority:    pos,
		Comments:    []*Comment{},
		HostColumn:  columnPtr,
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
func (c *Column) AddTask(name, descr string) *Task {

	newT := NewTask(name, descr, float64(len(c.Tasks)), c)
	c.Tasks = append(c.Tasks, newT)
	return newT
}

// AddComment func
func (t *Task) AddComment(text string) *Comment {
	newCom := NewComment(text)
	t.Comments = append(t.Comments, newCom)
	return newCom
}

// ChangePosition changes task position inside the column
func (t *Task) ChangePosition(newPosition int) {
	allTasks := t.HostColumn.Tasks
	taskCount := len(allTasks)
	if taskCount <= 1 {
		return
	}
	newPosition = getBoundIndex(newPosition, taskCount)
	newPriorityFrom := allTasks[newPosition].Priority
	var newPriorityTo float64

	if newPosition == 0 {
		newPriorityTo = newPriorityFrom - 1
	} else if newPosition == taskCount-1 {
		newPriorityTo = newPriorityFrom + 1
	} else if newPriorityFrom > t.Priority {
		newPriorityTo = t.HostColumn.Tasks[newPosition+1].Priority
	} else if newPriorityFrom < t.Priority {
		newPriorityTo = t.HostColumn.Tasks[newPosition-1].Priority
	}
	t.Priority = getBetweenPriority(newPriorityFrom, newPriorityTo)

	sort.SliceStable(allTasks, func(i, j int) bool {
		return allTasks[i].Priority < allTasks[j].Priority
	})
}

func getBoundIndex(unboundIndex, len int) int {
	if unboundIndex >= len-1 {
		return len - 1
	}
	if unboundIndex < 0 {
		return 0
	}
	return unboundIndex
}

func getBetweenPriority(priorityFrom, priorityTo float64) float64 {
	return (priorityFrom + priorityTo) / 2
}

// PrintObj prints user's board contents
func printObj(obj interface{}) {
	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
}

func main() {
	testProject := NewProject("golang", "")
	testProject.AddColumn("TO_DO").AddTask("gorello", "").AddComment("gambare!")
	testProject.Columns[1].AddTask("NOT important task", "")
	testProject.AddColumn("Done")
	importantTask := testProject.Columns[1].AddTask("very important task", "")
	printObj(testProject)
	importantTask.ChangePosition(0)
	printObj(testProject)

}
