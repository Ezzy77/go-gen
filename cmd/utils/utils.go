package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// GenerateProjectFolder generates the project folder structure
func GenerateProjectFolder(projectDirectory string) error {
	// Define folder names
	folders := []string{
		"bin",
		"cmd/api",
		"models",
		"migrations",
	}

	// Create folders
	for _, folder := range folders {
		err := os.MkdirAll(filepath.Join(projectDirectory, folder), 0755)
		if err != nil {
			return err
		}
	}

	// Create main.go
	mainFilePath := filepath.Join(projectDirectory, "cmd/api/main.go")
	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		return err
	}
	defer mainFile.Close()

	// Create routes.go
	routesFilePath := filepath.Join(projectDirectory, "cmd/api/routes.go")
	routesFile, err := os.Create(routesFilePath)
	if err != nil {
		return err
	}
	defer routesFile.Close()

	// Create handlers.go
	handlersFilePath := filepath.Join(projectDirectory, "cmd/api/todo.go")
	handlersFile, err := os.Create(handlersFilePath)
	if err != nil {
		return err
	}
	defer handlersFile.Close()

	// Create model.go
	modelFilePath := filepath.Join(projectDirectory, "models/todo.go")
	modelFile, err := os.Create(modelFilePath)
	if err != nil {
		return err
	}
	defer modelFile.Close()

	return nil
}

// GenerateCustomBoilerplate Generate boilerplate code based on selected options
func GenerateCustomBoilerplate(routingFramework, database, projectDirectory string) error {
	// Create main.go file with the selected routing framework
	if err := generateMainFile(routingFramework, projectDirectory); err != nil {
		return err
	}

	// Generate handlers, services, and models based on the selected database
	if err := generateHandlers(routingFramework, projectDirectory); err != nil {
		return err
	}
	if err := generateRoutes(routingFramework, projectDirectory); err != nil {
		return err
	}
	if err := generateModels(projectDirectory, database); err != nil {
		return err
	}

	return nil
}

// Function to generate main.go file
func generateMainFile(routingFramework, projectDirectory string) error {
	// Define main file content based on the selected routing framework
	var mainContent string
	switch routingFramework {
	case "gorilla/mux":
		mainContent = `

package main

import (
	"flag"
	"fmt"
	"github.com/username/your_project_name/models"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// holds application config
type config struct {
	port int
	env  string
}

// application struct holds dependencies for http handler and middleware
type application struct {
	config config
	logger *log.Logger
	todos  *models.TodoModel
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// init new logger that writes messages to the standard out stream
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
		todos:  &models.TodoModel{},
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, server.Addr)

	err := server.ListenAndServe()
	logger.Fatal(err)
}
`
	case "Gin":
		mainContent = `
package main

import (
	"flag"
	"fmt"
	"github.com/username/your_project_name/models"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// holds application config
type config struct {
	port int
	env  string
}

// application struct holds dependencies for http handler and middleware
type application struct {
	config config
	logger *log.Logger
	todos  *models.TodoModel
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// init new logger that writes messages to the standard out stream
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
		todos:  &models.TodoModel{},
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, server.Addr)

	err := server.ListenAndServe()
	logger.Fatal(err)
}
`
	}

	// Write main file content to main.go
	mainFilePath := filepath.Join(projectDirectory, "cmd/api/main.go")
	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		return err
	}
	defer mainFile.Close()

	_, err = mainFile.WriteString(mainContent)
	if err != nil {
		return err
	}
	return nil
}

// Function to generate handlers
func generateHandlers(routingFramework, projectDirectory string) error {
	// Define handlers file content based on the selected routing framework
	var handlersContent string
	switch routingFramework {
	case "gorilla/mux":
		handlersContent = `
package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/username/your_project_name/models"
	"net/http"
	"time"
)

func (app *application) getTodosHandler(w http.ResponseWriter, r *http.Request) {
	// Convert todos slice to JSON
	todos, err := app.todos.GetTodos()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if len(todos) == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize a new instance of the Todo struct
	var todo models.Todo

	// Decode JSON request body into todo struct
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Decode JSON request body into todo struct
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Assign ID and CreatedAt fields
	todo.ID = uuid.New()
	todo.CreatedAt = time.Now()

	// create
	err = app.todos.CreateTodo(&todo)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Encode the todo as JSON and send response
	jsonBytes, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonBytes)
	if err != nil {
		return
	}
}

func (app *application) getTodoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID, err := uuid.Parse(params["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	todo, err := app.todos.GetTodoById(todoID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if todo == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		return
	}
	return
}

func (app *application) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID, err := uuid.Parse(params["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	updatedTodo, err := app.todos.UpdateTodo(todoID, &todo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	jsonBytes, err := json.Marshal(updatedTodo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		return
	}
}

func (app *application) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID, err := uuid.Parse(params["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = app.todos.DeleteTodo(todoID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// todos deleted successfully
	w.WriteHeader(http.StatusOK)
}
`
	case "Gin":
		handlersContent = `

package main

import (
	"fmt"
	"github.com/username/your_project_name/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (app *application) getTodosHandler(c *gin.Context) {

	res, err := app.todos.GetTodos()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	c.JSON(http.StatusAccepted, res)

}

func (app *application) createTodosHandler(c *gin.Context) {

	var todo models.Todo

	err := c.ShouldBindJSON(&todo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse todo",
		})
		return
	}

	todo.ID = uuid.New()
	todo.CreatedAt = time.Now()

	validate := validator.New()
	err = validate.Struct(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "error validating",
		})
		fmt.Println(err)
		return
	}

	err = app.todos.CreateTodo(&todo)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "bad request",
		})
		return
	}

	c.JSON(http.StatusCreated, &todo)

}

func (app *application) getTodoHandler(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	todo, err := app.todos.GetTodoById(uuidID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "could not find todo",
		})
		return
	}

	c.JSON(http.StatusAccepted, todo)

}

func (app *application) updateTodoHandler(c *gin.Context) {
	id := c.Param("id")

	uuidID, err := uuid.Parse(id)
	if err != nil {
		// Respond with error
	}
	todo := models.Todo{}

	err = c.ShouldBindJSON(&todo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse todo",
		})
		return
	}

	res, err := app.todos.UpdateTodo(uuidID, &todo)
	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not update todo",
		})
		return
	}

	c.JSON(http.StatusAccepted, res)

}

func (app *application) deleteTodoHandler(c *gin.Context) {
	id := c.Param("id")

	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "error parsing id",
		})
	}

	err = app.todos.DeleteTodo(uuidID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "could not find todo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted successfuly",
	})

}

	`
	}

	handlersFilePath := filepath.Join(projectDirectory, "cmd/api/todo.go")
	mainFile, err := os.Create(handlersFilePath)
	if err != nil {
		return err
	}
	defer mainFile.Close()

	_, err = mainFile.WriteString(handlersContent)
	if err != nil {
		return err
	}
	return nil
}

// Function to generate services
func generateRoutes(routingFramework, projectDirectory string) error {
	// Define route file content based on the selected routing framework
	var routesContent string
	switch routingFramework {
	case "gorilla/mux":
		routesContent = `
package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {

	route := mux.NewRouter()

	// todo routes
	route.HandleFunc("/v1/api/todos", app.getTodosHandler).Methods("GET")
	route.HandleFunc("/v1/api/todos", app.createTodoHandler).Methods("POST")
	route.HandleFunc("/v1/api/todos/{id}", app.getTodoHandler).Methods("GET")
	route.HandleFunc("/v1/api/todos/{id}", app.updateTodoHandler).Methods("PUT")
	route.HandleFunc("/v1/api/todos/{id}", app.deleteTodoHandler).Methods("DELETE")

	return route
}
`
	case "Gin":
		routesContent = `

package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {

	route := gin.Default()

	//todo routes
	route.GET("/v1/api/todos", app.getTodosHandler)
	route.POST("/v1/api/todos", app.createTodosHandler)
	route.GET("/v1/api/todos/:id", app.getTodoHandler)
	route.DELETE("/v1/api/todos/:id", app.deleteTodoHandler)
	route.PATCH("/v1/api/todos/:id", app.updateTodoHandler)

	return route
}

`

	}
	handlersFilePath := filepath.Join(projectDirectory, "cmd/api/routes.go")
	mainFile, err := os.Create(handlersFilePath)
	if err != nil {
		return err
	}
	defer mainFile.Close()

	_, err = mainFile.WriteString(routesContent)
	if err != nil {
		return err
	}
	return nil
}

// Function to generate models based on selected database
func generateModels(projectDirectory, database string) error {
	// Define models file content based on the selected database
	var modelsContent string
	switch database {
	case "None":
		modelsContent = `

package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID
	Title     string
	Completed bool
	CreatedAt time.Time
}

type Todos []Todo

var todos = Todos{
	{ID: uuid.New(), Title: "Example Todo 1", Completed: false, CreatedAt: time.Now()},
	{ID: uuid.New(), Title: "Example Todo 2", Completed: true, CreatedAt: time.Now()},
}

type TodoModel struct {
}

func (t *TodoModel) GetTodos() ([]*Todo, error) {
	// Create a slice of pointers.
	todoList := make([]*Todo, len(todos))

	// Fill up the slice with pointers to the todos.
	for i := range todos {
		todoList[i] = &todos[i]
	}

	return todoList, nil
}

func (t *TodoModel) CreateTodo(todo *Todo) error {
	todos = append(todos, *todo)
	return nil
}

func (t *TodoModel) GetTodoById(id uuid.UUID) (*Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("error getting todo")
}

func (t *TodoModel) UpdateTodo(id uuid.UUID, todo *Todo) (*Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("error getting todo")
}

func (t *TodoModel) DeleteTodo(id uuid.UUID) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	return nil
}
`
	}

	modelsFilePath := filepath.Join(projectDirectory, "models/todo.go")
	mainFile, err := os.Create(modelsFilePath)
	if err != nil {
		return err
	}
	defer mainFile.Close()

	_, err = mainFile.WriteString(modelsContent)
	if err != nil {
		return err
	}
	return nil
}

// InitGoModule initialises a Go module for the project located in the given project directory.
// It creates a command to run the "go mod init" command and sets the project directory as the working directory.
func InitGoModule(projectDir string) error {
	cmd := exec.Command("go", "mod", "init",
		"github.com/username/your_project_name")

	cmd.Dir = projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return err
	}

	// get all packages
	cmd2 := exec.Command("go", "mod", "tidy")
	cmd2.Dir = projectDir
	output2, err := cmd2.CombinedOutput()
	if err != nil {
		fmt.Println(string(output2))
		return err
	}

	fmt.Println("Go module initialized successfully")
	return nil
}
