package main

import (
	"database/sql"
)

// Storage represents all possible actions available to deal with data
type Storage interface {
	ReadProject(projectID int) (Project, error)
	ReadProjects() ([]Project, error)

	ReadColumn(columnID int) (Column, error)
	ReadColumns(projectID int) ([]Column, error)
	CheckColNameUniq(projectID int, name string) (bool, error)

	ReadTask(taskID int) (Task, error)
	ReadTasks(columnID int) ([]Task, error)

	ReadComment(commentID int) (Comment, error)
	ReadComments(taskID int) ([]Comment, error)

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
	q := `SELECT id, name, priority, project_id FROM columns WHERE project_id = ? ORDER by priority`
	rows, err := m.db.Query(q, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := []Column{}
	for rows.Next() {
		var c Column
		rows.Scan(&c.ID, &c.Name, &c.Priority, &c.ProjectID)
		columns = append(columns, c)
	}
	return columns, nil
}

func (m *mysql) ReadColumn(columnID int) (Column, error) {
	q := `SELECT id, name, priority, project_id FROM columns WHERE id = ?`
	var column Column
	return column, m.db.QueryRow(q, columnID).Scan(
		&column.ID, &column.Name, &column.Priority, &column.ProjectID)
}

func (m *mysql) CheckColNameUniq(projectID int, name string) (bool, error) {
	q := `SELECT count(1) FROM columns WHERE project_id = ? AND name = ?`
	var duplicateNamesCount int
	err := m.db.QueryRow(q, projectID, name).Scan(&duplicateNamesCount)
	return duplicateNamesCount == 0, err
}

func (m *mysql) ReadTasks(columnID int) ([]Task, error) {
	q := `SELECT id, name, description, priority, column_id FROM tasks WHERE column_id = ? ORDER BY priority`
	rows, err := m.db.Query(q, columnID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Name, &t.Description, &t.Priority, &t.ColumnID)
		tasks = append(tasks, t)
	}
	return tasks, nil
}
func (m *mysql) ReadTask(taskID int) (Task, error) {
	q := `SELECT id, name, description, priority, column_id FROM tasks WHERE id = ?`
	var task Task
	return task, m.db.QueryRow(q, taskID).Scan(
		&task.ID, &task.Name, &task.Description, &task.Priority, &task.ColumnID)
}

func (m *mysql) ReadComments(taskID int) ([]Comment, error) {
	q := `SELECT id, text, task_id, createTs FROM comments WHERE task_id = ? ORDER BY createTs desc`
	rows, err := m.db.Query(q, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var c Comment
		rows.Scan(&c.ID, &c.Text, &c.TaskID, &c.TimeStamp)
		comments = append(comments, c)
	}
	return comments, nil
}
func (m *mysql) ReadComment(commentID int) (Comment, error) {
	q := `SELECT id, text, task_id, createTs FROM comments WHERE id = ?`
	var comment Comment
	return comment, m.db.QueryRow(q, commentID).Scan(
		&comment.ID, &comment.Text, &comment.TaskID, &comment.TimeStamp)
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
	q, err := m.db.Prepare(`UPDATE columns SET name = ?, project_id = ?, priority = ? WHERE id = ?`)
	if err != nil {
		return c, err
	}
	_, err = q.Exec(c.Name, c.ProjectID, c.Priority, c.ID)
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
	c.ID = int(lastID64)
	c.Priority = float64(c.ID)

	return m.UpdateColumn(c)
}
func (m *mysql) UpdateTask(t Task) (Task, error) {
	q, err := m.db.Prepare(
		`UPDATE tasks SET name = ?, description = ?, priority = ?, column_id = ? WHERE id = ?`)
	if err != nil {
		return t, err
	}
	_, err = q.Exec(t.Name, t.Description, t.Priority, t.ColumnID, t.ID)
	if err != nil {
		return t, err
	}
	return m.ReadTask(t.ID)
}
func (m *mysql) SaveTask(t Task) (Task, error) {
	if t.ID != 0 {
		return m.UpdateTask(t)
	}
	q, err := m.db.Prepare(`INSERT INTO tasks (name, description, column_id) VALUES (?,?,?)`)
	if err != nil {
		return t, err
	}
	res, err := q.Exec(t.Name, t.Description, t.ColumnID)

	if err != nil {
		return t, err
	}
	lastID64, err := res.LastInsertId()
	if err != nil {
		return t, err
	}
	t.ID = int(lastID64)
	t.Priority = float64(t.ID)

	return m.UpdateTask(t)
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
