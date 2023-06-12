package usecase

import "api_server/domain"

// ここがいわゆるDataAccessInterface
type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
