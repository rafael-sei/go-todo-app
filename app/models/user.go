package models

import (
	"log"
	"time"
)

// ユーザーstruct
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	Todos     []Todo
}

// セッションstruct
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    string
	CreatedAt time.Time
}

// user作成
func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) VALUES (?,?,?,?,?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password), // paswwordをハッシュ化
		time.Now())          // 現在時刻

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// userの取得
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `SELECT * FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}

// userの更新
func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users SET name = ?, email = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// userの削除
func (u *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users WHERE id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ユーザテーブルからemailを使ってuserを取得
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `SELECT * FROM users WHERE email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

// session作成して取得する
func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	//sessionを作成
	cmd1 := `INSERT INTO sessions (
		uuid,
		email,
		user_id,
		created_at) values (?,?,?,?)`
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}
	//session取得
	cmd2 := `SELECT * FROM sessions WHERE user_id = ? AND email = ?`
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	)

	return session, err
}

// sessionがデータベースにあるかどうか
func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `SELECT * FROM sessions WHERE uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)
	// sessionがerrorであればboolの初期値のfalseを返す。
	if err != nil {
		return
	}
	// session.IDが初期値の0じゃ無ければ、trueを返す。
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

// sessionのUUIDを削除
func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `DELETE FROM sessions WHERE uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// sessions.user_id & users.idが一致するものを取得
func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select * from users where id = ?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)
	return user, err
}
