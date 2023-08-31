package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Bionic2113/avito/internal/models"
)

type SegmentRepository interface {
	FindByName(name string) (*models.Segment, error)
	FindAll() ([]models.Segment, error)
	Create(name string) (*models.Segment, error)
	Update(s *models.Segment) (*models.Segment, error)
	Delete(name string) error
}

type SegmentRepositoryDB struct {
	*sql.DB
}

func (sdb *SegmentRepositoryDB) FindByName(name string) (*models.Segment, error) {
	s := &models.Segment{}
	row := sdb.QueryRow("select * from segment where name = $1", name)
	err := row.Scan(&s.Id, &s.Name, &s.Active)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (sdb *SegmentRepositoryDB) FindAll() ([]models.Segment, error) {
	rows, err := sdb.Query("SELECT * FROM segment as s where s.active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	segments := []models.Segment{}

	for rows.Next() {
		segment := models.Segment{}
		err = rows.Scan(&segment.Id, &segment.Name, &segment.Active)
		if err != nil {
			return nil, err
		}
		segments = append(segments, segment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return segments, nil
}

func (sdb *SegmentRepositoryDB) Create(name string) (*models.Segment, error) {
	tx, err := sdb.Begin()
	if err != nil {
		log.Println("ошибка в создании транзакции")
		return nil, err
	}
	s := &models.Segment{}
	row := sdb.QueryRow("insert into segment (name) values ($1) returning *", name)
	err = row.Scan(&s.Id, &s.Name, &s.Active)
	if err != nil {
		log.Println("уже существует")
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("ошибка в коммите")
		return nil, err
	}
	return s, nil
}

func (sdb *SegmentRepositoryDB) Update(s *models.Segment) (*models.Segment, error) {
	tx, err := sdb.Begin()
	if err != nil {
		log.Println("ошибка в создании транзакции")
		return nil, err
	}
	row := sdb.QueryRow("update segment set id = $1, name = $2, active = $3 where name $2 returning *", s.Id, s.Name, s.Active)
	err = row.Scan(&s.Id, &s.Name, &s.Active)
	if err != nil {
		log.Println("уже существует")
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("ошибка в коммите")
		return nil, err
	}
	return s, nil
}

func (sdb *SegmentRepositoryDB) Delete(name string) error {
	// if _, err := sdb.FindByName(name); err != nil {
	// 	log.Println("Ошибка при выполнении запроса:", err)
	// 	return err
	// }
	log.Println("start Delete()")
	tx, err := sdb.Begin()
	if err != nil {
		log.Println("Ошибка при tx.Begin():", err)
		return err
	}
	result, err := sdb.Exec("update segment set active = false where name = $1 and active = true", name)
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		tx.Rollback()
		return err
	}
	count, err := result.RowsAffected()
	if count == 0 || err != nil {
		log.Println("Ошибка при RowsAffected:", err)
		tx.Rollback()
		return errors.New("not found")
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Ошибка при tx.Commit:", err)
		return err
	}
	log.Println("finish Delete()")
	return nil
}
