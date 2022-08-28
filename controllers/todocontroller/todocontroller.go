package todocontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/itopangala/go-todo/entities"
	"github.com/itopangala/go-todo/libraries"
	"github.com/itopangala/go-todo/models"
)

var validation = libraries.NewValidation()
var todoModel = models.NewTodoModel()

func Index(response http.ResponseWriter, request *http.Request) {

	todo, _ := todoModel.FindAll()
	data := map[string]interface{}{
		"todo": todo,
	}

	temp, err := template.ParseFiles("views/todo/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

}
func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/todo/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var todo entities.Todo
		todo.Kegiatan = request.Form.Get("kegiatan")
		todo.Catatan = request.Form.Get("catatan")
		todo.Prioritas = request.Form.Get("prioritas")
		todo.TenggatWaktu = request.Form.Get("tenggat_waktu")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(todo)

		if vErrors != nil {
			data["todo"] = todo
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data berhasil disimpan"
			todoModel.Create(todo)
		}

		temp, _ := template.ParseFiles("views/todo/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var todo entities.Todo
		todoModel.Find(id, &todo)

		data := map[string]interface{}{
			"todo": todo,
		}

		temp, err := template.ParseFiles("views/todo/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var todo entities.Todo
		todo.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 16)
		todo.Kegiatan = request.Form.Get("kegiatan")
		todo.Catatan = request.Form.Get("catatan")
		todo.Prioritas = request.Form.Get("prioritas")
		todo.TenggatWaktu = request.Form.Get("tenggat_waktu")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(todo)

		if vErrors != nil {
			data["todo"] = todo
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data berhasil diperbarui"
			todoModel.Update(todo)
		}

		temp, _ := template.ParseFiles("views/todo/edit.html")
		temp.Execute(response, data)
	}
}
func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	todoModel.Delete(id)

	http.Redirect(response, request, "/todo", http.StatusSeeOther)
}
