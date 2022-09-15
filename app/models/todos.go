package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// #todo作成 (テキスト)
func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos (
		content,
		user_id,
		created_at) VALUES (?,?,?)`
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// IDからtodoデータ取得
func GetTodo(id int) (todo Todo, err error) {
	cmd := `SELECT * FROM todos WHERE id = ?`
	todo = Todo{}
	// QueryRow 1レコード取得
	// Scan 取得したデータを入れる
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)
	return todo, err
}

// 全てのtodoデータ取得
func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT * FROM todos`
	// Query は条件に合うものを全て取得
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		// #todos = [{ID,Content,UserID,CreatedAt},{ID,Content,UserID,CreatedAt}]
		todos = append(todos, todo)
	}

	return todos, err
}

// UserIDからTodoのデータを取得
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `SELECT * FROM todos WHERE user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		// #todos = [{ID,Content,UserID,CreatedAt},{ID,Content,UserID,CreatedAt}]
		todos = append(todos, todo)
	}

	return todos, err
}

// #todoの更新
func (t *Todo) UpdateTodo() (err error) {
	cmd := `UPDATE todos SET content = ?, user_id = ? WHERE id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// #todoの削除
func (t *Todo) DeleteTodo() (err error) {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
