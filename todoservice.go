package main

import "fmt"

type TodoService struct {
	repository TodoRepository
}

func (service TodoService) Save(todo Todo) {
	fmt.Printf("Saving %s todo\n", todo.Name)
	err := service.repository.Save(todo)
	if err != nil {
		panic(err)
	}
}

func (service TodoService) Get(todoName string) Todo {
	fmt.Printf("Getting %s todo\n", todoName)
	todo, err := service.repository.Get(todoName)
	if err != nil {
		panic(err)
	}
	return todo
}

func NewTodoService(repo TodoRepository) TodoService {
	return TodoService{repository: repo}
}
