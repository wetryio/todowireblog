package main

import (
	"encoding/json"
	"io/ioutil"
)

type TodoRepository interface {
	Save(todo Todo) error
	Get(todoName string) (Todo, error)
}

type TodoFileRepository struct {
}

func NewTodoRepository() TodoRepository {
	return TodoFileRepository{}
}

func (repo TodoFileRepository) Save(todo Todo) error {
	marsh, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(getFileName(todo.Name), marsh, 0644)
	return err
}

func (repo TodoFileRepository) Get(todoName string) (Todo, error) {
	fileContent, err := ioutil.ReadFile(getFileName(todoName))
	if err != nil {
		return Todo{}, err
	}
	target := Todo{}
	err = json.Unmarshal(fileContent, &target)
	return target, err
}

func getFileName(todoName string) string {
	return "./" + todoName + ".todo.json"
}
