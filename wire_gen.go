// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func BuildTodoService() TodoService {
	todoRepository := NewTodoRepository()
	todoService := NewTodoService(todoRepository)
	return todoService
}