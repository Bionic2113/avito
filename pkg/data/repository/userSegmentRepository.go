package repository

import (
	"database/sql"

	"github.com/Bionic2113/avito/pkg/data/models"
)

type UserSegmentRepository interface{
  	FindById(id int) (*models.UserSegment, error)
	FindAll()([]models.UserSegment, error) 
	Update(us *models.UserSegment) error
	Create(us *models.UserSegment) error
	Delete(us *models.UserSegment) error
}

type UserSegmentDB struct{
  *sql.DB
}
func (usdb *UserSegmentDB) FindById(id int) (*models.UserSegment, error){
return nil, nil
}

func (usdb *UserSegmentDB) FindAll()([]models.UserSegment, error){
  return nil, nil
}

func (usdb *UserSegmentDB) Update(us *models.UserSegment) error{
  return nil
}

func (usdb *UserSegmentDB) Create(us *models.UserSegment) error{
  return nil
}

func (usdb *UserSegmentDB) Delete(us *models.UserSegment) error{
  return nil
}
