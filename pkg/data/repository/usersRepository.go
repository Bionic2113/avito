package repository

import(
  "github.com/Bionic2113/avito/pkg/models"
)

type UserRepository interface {
	FindById(id int) User
	FindAll()[]User 
	Update(u *User)
	Create(u *User)
	Delete(u *User)
}
