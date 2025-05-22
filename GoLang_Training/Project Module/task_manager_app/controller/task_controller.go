package controller

import (
	"database/sql"
	"errors"
	config "task_manager_api/db"
	"task_manager_api/model"
)

func GetAllUsers() ([]model.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByID(id int) (model.User, error) {
	var user model.User
	row := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return user, errors.New("user not found")
	}
	return user, err
}

func CreateUser(user model.User) error {
	_, err := config.DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func UpdateUser(id int, user model.User) error {
	res, err := config.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("user not found")
	}
	return nil
}

func DeleteUser(id int) error {
	res, err := config.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("user not found")
	}
	return nil
}
