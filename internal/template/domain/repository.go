package domain_template

const TodoRepositoryTemplate string = `
package repository

import storage_infra "{{.ModuleName}}/internal/infra/storage"

type (
	TodoRepository interface {
		// your function definition
		AddNewTodo(todo dto.Todo)(bool,error)
	}
	todoRepository struct {
		// your injected dependency, like logger, cache, mq in infra
		// for example
		q storage_infra.Querier
	}
)

func NewTodoRepository(q storage_infra.Querier) TodoRepository {
	return &todoRepository{
		q:q,
	}
}

func (t *todoRepository) AddNewTodo(todo dto.Todo)(bool,error){
	panic("unimplemented")
}
`
