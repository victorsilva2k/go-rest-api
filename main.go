package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:id`
	Item      string `json:"title`
	Completed bool   `json:completed`
}

var todos = []todo{
	{ID: "1", Item: "Clean Socks", Completed: false},
	{ID: "2", Item: "Clean Bag", Completed: true},
	{ID: "3", Item: "Study for exam", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")
}
