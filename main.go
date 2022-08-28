package main

import (
	"net/http"

	"github.com/itopangala/go-todo/controllers/todocontroller"
)

func main() {

	http.HandleFunc("/", todocontroller.Index)
	http.HandleFunc("/todo", todocontroller.Index)
	http.HandleFunc("/todo/index", todocontroller.Index)
	http.HandleFunc("/todo/add", todocontroller.Add)
	http.HandleFunc("/todo/edit", todocontroller.Edit)
	http.HandleFunc("/todo/delete", todocontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
