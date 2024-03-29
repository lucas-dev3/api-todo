package models

import "api-postgres/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.Connect()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT id, title, description, done FROM todos`

	rows, err := conn.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	return
}
