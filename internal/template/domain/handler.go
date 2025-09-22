package domain_template

const HTTPHandlerTemplate string = `
package handler
	
import "github.com/gin-gonic/gin"

func RegisterTodoRoutes(r *gin.Engine, th TodoHandler) {
	r.POST("/todo", th.AddNewTodo)
}
`

const TodoHandlerTemplate string = `
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
