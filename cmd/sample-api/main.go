package main

import (
	"net/http"

	"github.com/RyutaKinoshitaSALTO/GoAPI/controller"
	"github.com/RyutaKinoshitaSALTO/GoAPI/model/repository"
)

var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todos/", ro.HandleTodosRequest)
	server.ListenAndServe()
}
