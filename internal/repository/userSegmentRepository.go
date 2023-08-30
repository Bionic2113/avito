package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Bionic2113/avito/internal/models"
)

type UserSegmentRepository interface {
	FindById(id int) (*models.UserSegment, error)
	FindAll(active bool) ([]models.UserSegment, error)
	Update(us *models.UserSegment) (*models.UserSegment, error)
	Create(us *models.UserSegment) (*models.UserSegment, error)
	Delete(us *models.UserSegment) error
}

type UserSegmentDB struct {
	*sql.DB
}

func (usdb *UserSegmentDB) FindById(id int) (*models.UserSegment, error) {
	userSegment := &models.UserSegment{}
	row := usdb.QueryRow("SELECT * FROM user_segment AS us WHERE us.id = $1", id)
	err := row.Scan(
		&userSegment.Id,
		&userSegment.User_id,
		&userSegment.Segment_id,
		&userSegment.CreationTime,
		&userSegment.DeletionTime,
		&userSegment.Duration,
		&userSegment.Active,
	)
	if err != nil {
		return nil, err
	}
	return userSegment, nil
}

func (usdb *UserSegmentDB) FindAll(active bool) ([]models.UserSegment, error) {
	rows, err := usdb.Query("SELECT * FROM user_segment as us WHERE us.active = $1", active)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arrayUS := []models.UserSegment{}

	for rows.Next() {
		userSegment := models.UserSegment{}
		err = rows.Scan(
			&userSegment.Id,
			&userSegment.User_id,
			&userSegment.Segment_id,
			&userSegment.CreationTime,
			&userSegment.DeletionTime,
			&userSegment.Duration,
			&userSegment.Active,
		)
		if err != nil {
			return nil, err
		}
		arrayUS = append(arrayUS, userSegment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return arrayUS, nil
}

func (usdb *UserSegmentDB) Update(us *models.UserSegment) (*models.UserSegment, error) {
	tx, err := usdb.Begin()
	if err != nil {
		log.Println("ошибка в создании транзакции")
		return nil, err
	}
	_, err = usdb.Exec(
		"update user_segment set id = $1, user_id = $2, segment_id = $3, creation_time = $4, deletion_time = $5, duration = $5, active = $6 where id = $1",
		us.Id,
		us.User_id,
		us.Segment_id,
		us.CreationTime,
		us.DeletionTime,
		us.Duration,
		us.Active,
	)
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
	us, err = usdb.FindById(int(us.Id))
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (usdb *UserSegmentDB) Create(us *models.UserSegment) (*models.UserSegment, error) {
	tx, err := usdb.Begin()
	if err != nil {
		log.Println("ошибка в создании транзакции")
		return nil, err
	}
	_, err = usdb.Exec(
		"insert into user_segment (user_id, segment_id, deletion_time, duration, active) values ($1, $2, $3, $4, $5)",
		us.User_id,
		us.Segment_id,
		us.DeletionTime,
		us.Duration,
		us.Active,
	)
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
	us, err = usdb.FindById(int(us.Id))
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (usdb *UserSegmentDB) Delete(us *models.UserSegment) error {
	if _, err := usdb.FindById(int(us.Id)); err != nil {
		return err
	}
	tx, err := usdb.Begin()
	if err != nil {
		return err
	}
	result, err := usdb.Exec("update user_segment set active = false where id = $1 and active = true", us.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	count, err := result.RowsAffected()
	if count == 0 || err != nil {
		tx.Rollback()
		return errors.New("not found")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
