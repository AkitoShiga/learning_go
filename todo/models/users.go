package models

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos     []Todo
}

func (u *User) CreateUser() (err error) {

	query := `INSERT INTO
    			users (UUID, NAME, EMAIL, PASSWORD, CREATED_AT)
			  VALUES
			     ( ?, ?, ?, ?, ? )
			 `

	result, err := Db.Exec(
		query,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(result)
	}

	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)
	fmt.Println(user.ID)

	return user, err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	return
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	//fmt.Println(u.ID)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	//fmt.Println(u.ID)
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions(uuid, email, user_id, created_at) values (?,?,?,?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		fmt.Println(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	if err != nil {
		fmt.Println(err)
	}

	return
}
