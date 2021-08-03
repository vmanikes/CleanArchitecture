package users

import "CleanArchitecture/entity"

type Repository interface {
	GetUser(key string) (entity.User, error)
	AddUser(item entity.User) (err error)
}
