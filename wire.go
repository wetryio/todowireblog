//+build wireinject

package main

import (
	"github.com/google/wire"
)

func BuildTodoService() TodoService {
	wire.Build(NewTodoRepository, NewTodoService)
	return TodoService{}
}
