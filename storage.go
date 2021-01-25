package main

import (
	"database/sql"
)

// Storage represents all possible actions available to deal with data
type Storage interface {
	ReadProjects() ([]Project, error)

	ReadColumns(projectID int) ([]Column, error)
	ReadColumn(columnID int) (Column, error)

	ReadTasks(columnID int) ([]Task, error)
	ReadTask(taskID int) (Task, error)

	ReadComments(taskID int) ([]Comment, error)
	ReadComment(commentID int) (Comment, error)

	// SaveProject(...Project) error // TO_DO: divide into upd and create?
	// SaveColumn(...Column) error
	// SaveTask(...Task) error
	// SaveComment(...Comment) error

	// DeleteProject(...Project) error
	// DeleteColumn(...Column) error
	// DeleteTask(...Task) error
	// DeleteComment(...Comment) error

}

type mysql struct {
	db *sql.DB
}

//NewStorage returns storage implementation that satisfies the Storage interface
func NewStorage(db *sql.DB) Storage {
	return &mysql{
		db: db,
	}
}

//ReadProjects returns all projects stored in db
func (m *mysql) ReadProjects() ([]Project, error) {
	q := `SELECT id, name, description FROM projects ORDER BY NAME`
	rows, err := m.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := []Project{}
	for rows.Next() {
		var p Project
		rows.Scan(&p.ID, &p.Name, &p.Description)
		projects = append(projects, p)
	}

	return projects, nil
}

func (m *mysql) ReadColumns(projectID int) ([]Column, error) {
	q := `SELECT id, name FROM columns WHERE project_id = ?`
	rows, err := m.db.Query(q, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := []Column{}
	for rows.Next() {
		var c Column
		rows.Scan(&c.ID, &c.Name)
		columns = append(columns, c)
	}
	return columns, nil
}

func (m *mysql) ReadColumn(columnID int) (Column, error) {
	q := `SELECT id, name FROM columns WHERE id = ?`
	var column Column
	return column, m.db.QueryRow(q, columnID).Scan(&column.ID, &column.Name)
}

func (m *mysql) ReadTasks(columnID int) ([]Task, error) {
	q := `SELECT id, name, description FROM tasks WHERE column_id = ?`
	rows, err := m.db.Query(q, columnID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Name, &t.Description)
		tasks = append(tasks, t)
	}
	return tasks, nil
}
func (m *mysql) ReadTask(taskID int) (Task, error) {
	q := `SELECT id, name, description FROM tasks WHERE id = ?`
	var task Task
	return task, m.db.QueryRow(q, taskID).Scan(&task.ID, &task.Name, &task.Description)
}

func (m *mysql) ReadComments(taskID int) ([]Comment, error) {
	q := `SELECT id, text FROM comments WHERE task_id = ?`
	rows, err := m.db.Query(q, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var c Comment
		rows.Scan(&c.ID, &c.Text)
		comments = append(comments, c)
	}
	return comments, nil
}
func (m *mysql) ReadComment(commentID int) (Comment, error) {
	q := `SELECT id, text FROM comments WHERE id = ?`
	var comment Comment
	return comment, m.db.QueryRow(q, commentID).Scan(&comment.ID, &comment.Text)
}
