package opgolib2

import (
	"crypto/md5"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Db *sql.DB
}

type User struct {
	Id    int
	Name  string
	Pass  string
	Email string
}

func (u *Users) Create_db(fname string) {
	var (
		s string
	)
	u.Db, _ = sql.Open("sqlite3", fname)
	s = "CREATE TABLE Users(id integer primary key, name text, email text, pass text);"
	u.Db.Exec(s)
}

func (u *Users) Open_db(fname string) {
	u.Db, _ = sql.Open("sqlite3", fname)
}

func (u *Users) Add_user(name, email, pass string) error {
	var (
		stmt *sql.Stmt
		s    string
		p    string
	)
	p = Make_md5_hash(name, pass)
	s = "INSERT INTO Users(name,email,pass)\nVALUES(?,?,?);"
	stmt, _ = u.Db.Prepare(s)
	stmt.Exec(name, email, p)
	return nil
}

func (u *Users) Delete_user_id(id int) error {
	u.Db.Exec("DELETE FROM Users WHERE id=?;", id)
	return nil
}

func (u *Users) Delete_user_name(name string) error {
	u.Db.Exec("DELETE FROM Users WHERE name=?;", name)
	return nil
}

func (u *Users) Find_user(name string) *User {
	var (
		rows *sql.Rows
		s    string
		stmt *sql.Stmt
		user *User
	)
	user = &User{}
	s = "SELECT * FROM Users WHERE name=?;"
	stmt, _ = u.Db.Prepare(s)
	rows, _ = stmt.Query(name)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Pass)
		return user
	}
	rows.Close()
	return nil
}

func (u *Users) Find_user_id(id int) *User {
	var (
		rows *sql.Rows
		s    string
		user *User
	)
	s = "SELECT * FROM Users WHERE id=?;"
	rows, _ = u.Db.Query(s, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Pass)
		return user
	}
	rows.Close()
	return nil
}

func Make_md5_hash(name, pass string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(name+pass)))
}

func (u *Users) Check_pass(name, pass string) bool {
	var (
		user *User
	)
	user = u.Find_user(name)
	if user == nil {
		return false
	}
	if Make_md5_hash(name, pass) == user.Pass {
		return true
	}
	return false
}

func (u *Users) Update_user(name, newpass, newemail string) {
	var (
		p string
	)
	if newpass != "" {
		p = Make_md5_hash(name, newpass)
		u.Db.Exec("UPDATE Users\nSET pass=?\nWHERE name=?;", p, name)
	}
	if newemail != "" {
		u.Db.Exec("UPDATE Users\nSET email=?\nWHERE name=?;", newemail, name)
	}
}
