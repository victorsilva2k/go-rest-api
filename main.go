package main

import (
	"net/http"
	"errors"
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
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodobyID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Task not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}
func getTodobyID(id string) (*todo, error){
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")

}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodobyID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Task not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodos)
	router.Run("0.0.0.0:$PORT")
}
