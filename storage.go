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

	SaveProject(p Project) (Project, error) // TO_DO: divide into upd and create?
	SaveColumn(c Column) (Column, error)
	SaveTask(t Task) (Task, error)
	SaveComment(c Comment) (Comment, error)

	RemoveProject(projectID int) error
	RemoveColumn(columnID int) error
	RemoveTask(taskID int) error
	RemoveComment(commentID int) error
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

func (m *mysql) ReadProject(projectID int) (Project, error) {
	q := `SELECT id, name, description FROM projects WHERE id = ?`
	var project Project
	return project, m.db.QueryRow(q, projectID).Scan(&project.ID, &project.Name, &project.Description)
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

func (m *mysql) UpdateProject(p Project) (Project, error) {
	q, err := m.db.Prepare(`UPDATE projects SET name = ?, description = ? WHERE id = ?`)
	if err != nil {
		return p, err
	}
	_, err = q.Exec(p.Name, p.Description, p.ID)
	if err != nil {
		return p, err
	}
	return m.ReadProject(p.ID)
}

func (m *mysql) SaveProject(p Project) (Project, error) {
	if p.ID != 0 {
		return m.UpdateProject(p)
	}
	q, err := m.db.Prepare(`INSERT INTO projects (name, description) VALUES (?,?)`)
	if err != nil {
		return p, err
	}
	res, err := q.Exec(p.Name, p.Description)

	if err != nil {
		return p, err
	}
	lastID64, err := res.LastInsertId()
	if err != nil {
		return p, err
	}
	return m.ReadProject(int(lastID64))
}

func (m *mysql) UpdateColumn(c Column) (Column, error) {
	q, err := m.db.Prepare(`UPDATE columns SET name = ?, project_id = ? WHERE id = ?`)
	if err != nil {
		return c, err
	}
	_, err = q.Exec(c.Name, c.ProjectID, c.ID)
	if err != nil {
		return c, err
	}
	return m.ReadColumn(c.ID)
}
func (m *mysql) SaveColumn(c Column) (Column, error) {
	if c.ID != 0 {
		return m.UpdateColumn(c)
	}
	q, err := m.db.Prepare(`INSERT INTO columns (name, project_id) VALUES (?,?)`)
	if err != nil {
		return c, err
	}
	res, err := q.Exec(c.Name, c.ProjectID)

	if err != nil {
		return c, err
	}
	lastID64, err := res.LastInsertId()
	if err != nil {
		return c, err
	}
	return m.ReadColumn(int(lastID64))
}
func (m *mysql) UpdateTask(t Task) (Task, error) {
	q, err := m.db.Prepare(
		`UPDATE tasks SET name = ?, description = ?, priority = ?, column_id = ? WHERE id = ?`)
	if err != nil {
		return t, err
	}
	_, err = q.Exec(t.Name, t.Description, t.Priority, t.ColumnID)
	if err != nil {
		return t, err
	}
	return m.ReadTask(t.ID)
}
func (m *mysql) SaveTask(t Task) (Task, error) {
	if t.ID != 0 {
		return m.UpdateTask(t)
	}
	q, err := m.db.Prepare(`INSERT INTO tasks (name, description, priority, column_id) VALUES (?,?,?,?)`)
	if err != nil {
		return t, err
	}
	res, err := q.Exec(t.Name, t.Description, t.Priority, t.ColumnID)

	if err != nil {
		return t, err
	}
	lastID64, err := res.LastInsertId()
	if err != nil {
		return t, err
	}
	return m.ReadTask(int(lastID64))
}
func (m *mysql) UpdateComment(c Comment) (Comment, error) {
	q, err := m.db.Prepare(
		`UPDATE comments SET text = ?, task_id = ? WHERE id = ?`)
	if err != nil {
		return c, err
	}
	_, err = q.Exec(c.Text, c.TaskID)
	if err != nil {
		return c, err
	}
	return m.ReadComment(c.ID)
}
func (m *mysql) SaveComment(c Comment) (Comment, error) {
	if c.ID != 0 {
		return m.UpdateComment(c)
	}
	q, err := m.db.Prepare(`INSERT INTO comments (text, task_id) VALUES (?,?)`)
	if err != nil {
		return c, err
	}
	res, err := q.Exec(c.Text, c.TaskID)

	if err != nil {
		return c, err
	}
	lastID64, err := res.LastInsertId()
	if err != nil {
		return c, err
	}
	return m.ReadComment(int(lastID64))
}

func (m *mysql) RemoveProject(projectID int) error {
	_, err := m.db.Exec(`DELETE FROM projects WHERE id = ?`, projectID)
	return err
}
func (m *mysql) RemoveColumn(columnID int) error {
	_, err := m.db.Exec(`DELETE FROM columns WHERE id = ?`, columnID)
	return err
}
func (m *mysql) RemoveTask(taskID int) error {
	_, err := m.db.Exec(`DELETE FROM tasks WHERE id = ?`, taskID)
	return err
}
func (m *mysql) RemoveComment(commentID int) error {
	_, err := m.db.Exec(`DELETE FROM comments WHERE id = ?`, commentID)
	return err
}
