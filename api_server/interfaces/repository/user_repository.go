package repository

import (
	"api_server/domain"
)

// DataAccessInterfaceの実装
// Gatewayに該当する
// クエリを定義して抽象化したDB（InfraStructure）アクセスのインターフェースを使う
type UserRepository struct {
	SqlHandler //これはinfrastracture層のハンドラではない
}

// Repositoryは直接エンティティを扱って良い
func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	sql := "INSERT INTO users (first_name, last_name) VALUES (?, ?)"
	result, err := repo.Execute(sql, u.FirstName, u.LastName)

	if err != nil {
		return
	}

	id64, err := result.LastInsertId()

	if err != nil {
		return
	}

	id = int(id64)

	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	sql := "SELECT id, first_name, last_name FROM users WHERE id = ?"
	row, err := repo.Query(sql, identifier)
	defer row.Close()

	if err != nil {
		return
	}

	var id int
	var firstName string
	var lastName string

	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return
	}

	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName

	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	sql := "SELECT id, first_name, last_name FROM users"
	rows, err := repo.Query(sql)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, user)
	}
	return
}
