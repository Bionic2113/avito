package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Bionic2113/avito/internal/models"
)

type UserRepository interface {
	FindById(id int) (*models.User, error)
	FindAll() ([]models.User, error)
	Update(u *models.User) (*models.User, error)
	Create(u *models.User) (*models.User, error)
	Delete(u *models.User) error
}

type UserRepoDB struct {
	*sql.DB
}

func (udb *UserRepoDB) FindById(id int) (*models.User, error) {
	user := &models.User{}
	row := udb.QueryRow("SELECT * FROM users AS u WHERE u.id = $1", id)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Active)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (udb *UserRepoDB) FindAll() ([]models.User, error) {
	rows, err := udb.Query("SELECT * FROM users as u WHERE u.active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Active)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (udb *UserRepoDB) Update(u *models.User) (*models.User, error) {
	tx, err := udb.Begin()
	if err != nil {
		log.Println("ошибка в создании транзакции")
		return nil, err
	}
	_, err = udb.Exec("update users set id = $1, first_name = $2, last_name = $3, active = $4 where id = $1", u.Id, u.FirstName, u.LastName, u.Active)
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
	u, err = udb.FindById(int(u.Id))
	if err != nil {
		return nil, err
	}
	return u, nil

}

func (udb *UserRepoDB) Create(u *models.User) (*models.User, error) {
	tx, err := udb.Begin()
	if err != nil {
		log.Println("ошибка в создании транзакции")
		return nil, err
	}
	_, err = udb.Exec("insert into users (first_name, last_name) values ($1, $2)", u.FirstName, u.LastName)
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
	u, err = udb.FindById(int(u.Id))
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (udb *UserRepoDB) Delete(u *models.User) error {
	if _, err := udb.FindById(int(u.Id)); err != nil {
		return err
	}
	tx, err := udb.Begin()
	if err != nil {
		return err
	}
	result, err := udb.Exec("update users set active = false where id = $1 and active = true", u.Id)
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
