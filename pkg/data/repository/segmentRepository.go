package repository

import (
	"database/sql"

	"github.com/Bionic2113/avito/pkg/data/models"
)

type SegmentRepository interface{
  FindByName(name string)(*models.Segment, error)
  FindAll()([]models.Segment, error)
  Create(s *models.Segment) error
  Update(s *models.Segment) error
  Delete(s *models.Segment) error
}

type SegmentRepositoryDB struct{
  *sql.DB
}

func (sdb *SegmentRepositoryDB) FindByName(name string) (*models.Segment, error){
  return nil, nil
}

func (sdb *SegmentRepositoryDB) FindAll()([]models.Segment, error){return nil,nil}

func (sdb *SegmentRepositoryDB) Create(s *models.Segment) error{return nil}

func (sdb *SegmentRepositoryDB) Update(s *models.Segment) error{return nil}

func (sdb *SegmentRepositoryDB) Delete(s *models.Segment) error{return nil}


