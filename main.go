package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

// ServerAddr defines the http host and port of the beer server
const ServerAddr = "0.0.0.0:8080"

var repo Storage
var router *httprouter.Router

// note: avoid using init
func init() {
	var err error

	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	db, err := sql.Open("mysql", mysqlUser+":"+mysqlPwd+"@tcp("+
		mysqlHost+":"+mysqlPort+")/gorello")
	// db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/gorello")
	// db, err := sql.Open("mysql", "gorello:password@/gorello")

	if err != nil {
		panic(err) //TO_DO: or log fatal?
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	// defer db.Close()

	repo = NewStorage(db)

	router = httprouter.New()
	router.GET("/projects", GetProjects)
	router.GET("/columns/:project_id", GetColumns)
	router.GET("/column/:column_id", GetColumn)
	router.GET("/tasks/:column_id", GetTasks)
	router.GET("/task/:task_id", GetTask)
	router.GET("/comments/:task_id", GetComments)
	router.GET("/comment/:comment_id", GetComment)

	router.POST("/project", CreateProject)
	router.POST("/column", CreateColumn)
	router.POST("/task", CreateTask)
	router.POST("/comment", CreateComment)

	router.PUT("/move/column/:column_id", MoveColumn) // ?new_pos=int
	router.PUT("/move/task/:task_id", MoveTask)       // ?new_pos=int

	router.DELETE("/project/:project_id", DeleteProject)
	router.DELETE("/column/:column_id", DeleteColumn)
	router.DELETE("/task/:task_id", DeleteTask)
	router.DELETE("/comment/:comment_id", DeleteComment)

}

func main() {
	fmt.Println("The gorello server is on tap at " + ServerAddr)
	log.Fatal(http.ListenAndServe(ServerAddr, router))
}
