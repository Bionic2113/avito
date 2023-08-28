package repository

import (
	"database/sql"

	"github.com/Bionic2113/avito/pkg/data/models"
)

type UserRepository interface {
	FindById(id int) (*models.User, error)
	FindAll()([]models.User, error) 
	Update(u *models.User) error
	Create(u *models.User) error
	Delete(u *models.User) error
}

type UserDbRepoImpl struct{
	*sql.DB
}

func (uimpl *UserDbRepoImpl) FindById(id int) (*models.User, error){
	user := &models.User{}
	row := uimpl.QueryRow("SELECT * FROM USER AS u WHERE u.id = $1", id)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Segments, &user.Active)
	if err != nil{
	  return nil, err
	}
  return user, nil
}

func (uimpl *UserDbRepoImpl) FindAll() ([]models.User, error){
  return nil, nil
}

func (uimpl *UserDbRepoImpl) Update(u *models.User) error{
  return nil
}

func (uimpl *UserDbRepoImpl) Create(u *models.User) error{
  return nil
}

func (uimpl *UserDbRepoImpl) Delete(u *models.User) error{
  return nil
}

