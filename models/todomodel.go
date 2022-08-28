package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/itopangala/go-todo/config"
	"github.com/itopangala/go-todo/entities"
)

type TodoModel struct {
	conn *sql.DB
}

func NewTodoModel() *TodoModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &TodoModel{
		conn: conn,
	}
}

func (p *TodoModel) FindAll() ([]entities.Todo, error) {
	rows, err := p.conn.Query("select * from todo_table")
	if err != nil {
		return []entities.Todo{}, err
	}
	defer rows.Close()

	var dataTodo []entities.Todo
	for rows.Next() {
		var todo entities.Todo
		rows.Scan(&todo.Id,
			&todo.Kegiatan,
			&todo.Catatan,
			&todo.Prioritas,
			&todo.TenggatWaktu)

		if todo.Prioritas == "1" {
			todo.Prioritas = "Penting"
		} else {
			todo.Prioritas = "Tidak Penting"
		}

		tenggat_waktu, _ := time.Parse("2006-01-02", todo.TenggatWaktu)
		todo.TenggatWaktu = tenggat_waktu.Format("02-01-2006")

		dataTodo = append(dataTodo, todo)
	}
	return dataTodo, nil
}

func (p *TodoModel) Create(todo entities.Todo) bool {
	result, err := p.conn.Exec("insert into todo_table (kegiatan, catatan, prioritas, tenggat_waktu) values(?,?,?,?)",
		todo.Kegiatan, todo.Catatan, todo.Prioritas, todo.TenggatWaktu)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *TodoModel) Find(id int64, todo *entities.Todo) error {

	return p.conn.QueryRow("select * from todo_table where id = ?", id).Scan(
		&todo.Id,
		&todo.Kegiatan,
		&todo.Catatan,
		&todo.Prioritas,
		&todo.TenggatWaktu)
}

func (p *TodoModel) Update(todo entities.Todo) error {

	_, err := p.conn.Exec(
		"update todo_table set kegiatan = ?, catatan = ?, prioritas = ?, tenggat_waktu = ? where id = ?",
		todo.Kegiatan, todo.Catatan, todo.Prioritas, todo.TenggatWaktu, todo.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *TodoModel) Delete(id int64) {
	p.conn.Exec("delete from todo_table where id = ?", id)
}
