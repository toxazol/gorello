package main

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
