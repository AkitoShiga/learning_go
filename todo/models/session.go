package models

import (
    _ "github.com/mattn/go-sqlite3"
    "log"
    "time"
)

type Session struct {
    ID        int
    UUID      string
    Email     string
    UserID    int
    CreatedAt time.Time
}

func (session *Session) CheckSession() (valid bool, err error) {
    cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = ?`
    err = Db.QueryRow(cmd, session.UUID).Scan(
        &session.ID,
        &session.UUID,
        &session.Email,
        &session.UserID,
        &session.CreatedAt)

    if err != nil {
        valid = false
        return
    }

    if session.ID != 0 {
        valid = true
    }

    return
}

func (session *Session) DeleteSessionByUUID() (err error) {
    cmd := `delete from sessions where uuid =?`
    _, err = Db.Exec(cmd, session.UUID)
    if err != nil {
        log.Fatalln(err)
    }
    return
}

func (session *Session) GetUserBySession() (user User, err error) {
    user = User{}
    cmd := `select id, uuid, name, email, created_at FROM users where id = ?`
    err = Db.QueryRow(cmd, session.UserID).Scan(
        &user.ID,
        &user.UUID,
        &user.Name,
        &user.Email,
        &user.CreatedAt)

    return
}
