package domain_template

const GinHTTPHandlerTemplate string = `
package handler
	
import "github.com/gin-gonic/gin"

func RegisterTodoRoutes(r *gin.Engine, th TodoHandler) {
	r.POST("/todo", th.AddNewTodo)
}
`

const ChiHTTPHandlerTemplate string = `
package handler

import "github.com/go-chi/chi/v5"

func RegisterTodoRoutes(r *chi.Mux, th TodoHandler) {
	r.Post("/todo",th.AddNewTodo)
}
`
const EchoHTTPHandlerTemplate string = `
package handler

import "github.com/labstack/echo/v4"

func RegisterTodoRoutes(r *echo.Echo, th TodoHandler) {
	r.POST("/todo",th.AddNewTodo)
}
`

const FiberHTTPHandlerTemplate string = `
package handler

import "github.com/gofiber/fiber/v2"

func RegisterTodoRoutes(r *fiber.App, th TodoHandler) {
	r.Post("/todo",th.AddNewTodo)
}
`
