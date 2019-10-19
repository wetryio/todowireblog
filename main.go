package main

import (
	"fmt"
	"os"
)

func main() {
	todo := Todo{Name: os.Args[1], Description: os.Args[2]}
	service := BuildTodoService()

	service.Save(todo)
	todoFromService := service.Get(todo.Name)

	fmt.Println(todoFromService)
}
