package domain_template

const GinTodoHandlerTemplate string = `
package handler

import "github.com/gin-gonic/gin"

type (
	TodoHandler interface {
		// your function definition
		AddNewTodo(ctx *gin.Context)
	}
	todoHandler struct {
		// your injected dependency
		// for example
		ts service.TodoService
	}
)

func NewTodoHandler(ts service.TodoService) TodoHandler {
	return &todoHandler{
		ts: ts,
	}
}

func (t *todoHandler) AddNewTodo(ctx *gin.Context){
	panic("unimplemented")
}
`

const ChiTodoHandlerTemplate string = `
package handler

type (
	TodoHandler interface {
		// your function definition
		AddNewTodo(w http.ResponseWriter, r *http.Request)
	}
	todoHandler struct {
		// your injected dependency
		// for example
		ts service.TodoService
	}
)

func NewTodoHandler(ts service.TodoService) TodoHandler {
	return &todoHandler{
		ts: ts,
	}
}

func (t *todoHandler) AddNewTodo(w http.ResponseWriter, r *http.Request){
	panic("unimplemented")
}
`

const EchoTodoHandlerTemplate string = `
package handler

import "github.com/labstack/echo/v4"

type (
	TodoHandler interface {
		// your function definition
		AddNewTodo(c echo.Context) error
	}
	todoHandler struct {
		// your injected dependency
		// for example
		ts service.TodoService
	}
)

func NewTodoHandler(ts service.TodoService) TodoHandler {
	return &todoHandler{
		ts: ts,
	}
}

func (t *todoHandler) AddNewTodo(c echo.Context)error{
	panic("unimplemented")
}
`

const FiberTodoHandlerTemplate string = `
package handler

import "github.com/gofiber/fiber/v2"

type (
	TodoHandler interface {
		// your function definition
		AddNewTodo(c *fiber.Ctx) error
	}
	todoHandler struct {
		// your injected dependency
		// for example
		ts service.TodoService
	}
)

func NewTodoHandler(ts service.TodoService) TodoHandler {
	return &todoHandler{
		ts: ts,
	}
}

func (t *todoHandler) AddNewTodo(c *fiber.Ctx)error{
	panic("unimplemented")
}
`

const GorillaTodoHandlerTemplate string = `
package handler

type (
	TodoHandler interface {
		// your function definition
		AddNewTodo(w http.ResponseWriter, r *http.Request)
	}
	todoHandler struct {
		// your injected dependency
		// for example
		ts service.TodoService
	}
)

func NewTodoHandler(ts service.TodoService) TodoHandler {
	return &todoHandler{
		ts: ts,
	}
}

func (t *todoHandler) AddNewTodo(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
`
