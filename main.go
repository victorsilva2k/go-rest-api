package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:id`
	Item      string `json:"item`
	Completed bool   `json:completed`
}

var todos = []todo{
	{ID: "1", Item: "Cleans Socks", Completed: false},
	{ID: "2", Item: "Clean Bag", Completed: true},
	{ID: "3", Item: "Study for exam", Completed: false},
}

func addTodos(context *gin.Context)  {
	var newtodo todo

	if err := context.BindJSON(&newtodo); err != nil {
		return
	}

	todos = append(todos, newtodo)

	context.IndentedJSON(http.StatusCreated, newtodo)

}

func getTodos(cont *gin.Context) {
	cont.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)
	router.Run("localhost:9090")
}
