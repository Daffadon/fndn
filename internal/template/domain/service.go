package domain_template

const TodoServiceTemplate string = `
package service

type (
	TodoService interface {
		// your function definition
		AddNewTodo(todo dto.Todo) (bool,error)
	}
	todoService struct {
		// your injected dependency, like logger, cache, mq in infra
		// for example
			tr repository.TodoRepository
	}
)

func NewTodoService(tr repository.TodoRepository) TodoService {
	return &todoService{
		tr:tr,
	}
}

func (t *todoService) AddNewTodo(todo dto.Todo) (bool,error){
	panic("unimplemented")
}
`
