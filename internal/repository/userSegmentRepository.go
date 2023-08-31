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
	FindAllById(user_id int, active bool) ([]models.UserSegment, error)
}

type UserSegmentDB struct {
	*sql.DB
}

func (usdb *UserSegmentDB) FindById(id int) (*models.UserSegment, error) {
	userSegment := &models.UserSegment{}
	row := usdb.QueryRow("SELECT * FROM user_segment AS us WHERE us.id = $1", id)
	err := scanOne(row, userSegment)
	if err != nil {
		log.Println("Ошибка при сканировании:", err)
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
	return scanForResult(rows)
}

func (usdb *UserSegmentDB) FindAllById(user_id int, active bool) ([]models.UserSegment, error) {
	rows, err := usdb.Query("SELECT * FROM user_segment as us WHERE us.User_id = $1 and us.active = $2", user_id, active)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanForResult(rows)
}

func scanOne(row *sql.Row, userSegment *models.UserSegment) error {
	return row.Scan(
		&userSegment.Id,
		&userSegment.User_id,
		&userSegment.Segment_name,
		&userSegment.CreationTime,
		&sql.NullInt64{Valid: true, Int64: int64(userSegment.DeletionTime)},
		&sql.NullInt64{Int64: int64(userSegment.Duration), Valid: true},
		&userSegment.Active,
	)
}

func scanForResult(rows *sql.Rows) ([]models.UserSegment, error) {
	arrayUS := []models.UserSegment{}
	for rows.Next() {
		userSegment := models.UserSegment{}
		err := rows.Scan(
			&userSegment.Id,
			&userSegment.User_id,
			&userSegment.Segment_name,
			&userSegment.CreationTime,
			&sql.NullInt64{Valid: true, Int64: int64(userSegment.DeletionTime)},
			&sql.NullInt64{Int64: int64(userSegment.Duration), Valid: true},
			&userSegment.Active,
		)
		if err != nil {
			return nil, err
		}
		arrayUS = append(arrayUS, userSegment)
	}

	if err := rows.Err(); err != nil {
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
	result := usdb.QueryRow(
		"update user_segment set id = $1, user_id = $2, segment_name = $3, creation_time = $4, deletion_time = $5, duration = $6, active = $7 where id = $1 returning *",
		us.Id,
		us.User_id,
		us.Segment_name,
		us.CreationTime,
		us.DeletionTime,
		us.Duration,
		us.Active,
	)
	err = scanOne(result, us)
	if err != nil {
		log.Println("Ошибка ", err)
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("ошибка в коммите")
		return nil, err
	}
	// us, err = usdb.FindById(int(us.Id))
	// if err != nil {
	// 	return nil, err
	// }
	return us, nil
}

func (usdb *UserSegmentDB) Create(us *models.UserSegment) (*models.UserSegment, error) {
	tx, err := usdb.Begin()
	if err != nil {
		log.Println("ошибка в создании транзакции")
		return nil, err
	}
	sdb := SegmentRepositoryDB{usdb.DB}
	if _, err := sdb.FindByName(us.Segment_name); err != nil {
		if _, err := sdb.Create(us.Segment_name); err != nil {
			return nil, err
		}
	}
	result := usdb.QueryRow(
		"insert into user_segment (user_id, segment_name) values ($1, $2) returning *",
		us.User_id,
		us.Segment_name,
	)
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		log.Println("уже существует")
		tx.Rollback()
		return nil, err
	}
	err = scanOne(result, us)
	if err != nil {
		log.Println("Ошибка при получении последнего добавленного айди:", err)
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("ошибка в коммите")
		return nil, err
	}
	// us, err = usdb.FindById(int(id))
	// if err != nil {
	// 	log.Println("Ошибка при получения чела:", err)
	// 	return nil, err
	// }
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
	result, err := usdb.Exec(
		"update user_segment set active = false, set deletion_time = EXTRACT(epoch from now()) where id = $1 and active = true",
		us.Id,
	)
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
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
