package domain_template

const TodoRepositoryTemplate string = `
package repository

import "{{.ModuleName}}/internal/domain/dto"

type (
	TodoRepository interface {
		// your function definition
		AddNewTodo(todo dto.Todo)(bool,error)
	}
	todoRepository struct {
	}
)

func NewTodoRepository() TodoRepository {
	return &todoRepository{
	}
}

func (t *todoRepository) AddNewTodo(todo dto.Todo)(bool,error){
	panic("unimplemented")
}
`
