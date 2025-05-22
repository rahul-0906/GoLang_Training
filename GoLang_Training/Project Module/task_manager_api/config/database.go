package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "saimuluk"
	password = "manasa"
	dbName   = "task_manager"
)

func DBConnection() *gorm.DB {

	sqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlConn), &gorm.Config{})
	if err != nil {
		fmt.Println(db)
	}

	return db
}
