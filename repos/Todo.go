package repos

import (
	_ "github.com/lib/pq"
	m "nhaoday.com/models"
)

func TodoList() (m.Todos, error) {

	rows, err := db.Query("Select * from todos order by id desc")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	var todos m.Todos
	for rows.Next() {
		var todo m.Todo
		err = rows.Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.Due)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func FindTodoById(Id int) (m.Todo, error) {
	row := db.QueryRow("select * from todos where Id = $1", Id)
	var todo m.Todo
	err := row.Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.Due)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func AddTodo(t *m.Todo) error {
	var insertedId int
	strSql := "insert into todos (Name, Due) values ($1, $2) returning id"
	err := db.QueryRow(strSql, t.Name, t.Due).Scan(&insertedId)
	if err != nil {
		return err
	}
	t.Id = int(insertedId)
	return nil
}

func UpdateTodo(id int, t m.Todo) error {
	strSql := "update todos set name=$1, completed=$2, due=$3 where id=$4"
	_, err := db.Exec(strSql, t.Name, t.Completed, t.Due, id)
	if err != nil {
		return err
	}
	return nil
}
