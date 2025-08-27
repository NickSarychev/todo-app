package main

import "github.com/NickSarychev/todo-app"

func main() {
	srv := new(todo.Server)
	srv.Run("8000")
}
