package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct { // creating a struct (similar to a Python class) for the todo items
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{ // Storage for the todo list items
	{
		ID:        "1",
		Item:      "Buy milk",
		Completed: false,
	},
	{
		ID:        "2",
		Item:      "Buy eggs",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Buy bread",
		Completed: false,
	},
}

func getTodos(context *gin.Context) { // Contains information about the incoming HTTP request
	context.IndentedJSON(http.StatusOK, todos) // converts todo struct to JSON
}

func addTodos(context *gin.Context) { // Adds an item to the todo list
	var newTodo todo
	if err := context.ShouldBindJSON(&newTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(id string) (*todo, error) { // Gets an item from the todo list by ID
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Todo not found": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Todo not found": err.Error()})
		return
	}

	todo.Completed = !todo.Completed
}

func main() {
	router := gin.Default() // This creates my 'server'
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)
	router.Run("localhost:8080")
}
