package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DataBase *sql.DB

func OpenDBConnection() {
	var err error
	host := "localhost"
	port := 5432
	username := "postgres"
	password := "Arul@123"
	dbName := "employeeDatabase"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)

	DataBase, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DataBase Connection Error:", err)
	}

	err = DataBase.Ping()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS employee (
    empId SERIAL PRIMARY KEY,
    name TEXT NOT NULL, 
    emailId TEXT NOT NULL,
    projectName TEXT NOT NULL,
    location TEXT NOT NULL,
    mobileNo TEXT NOT NULL
);`
	_, err = DataBase.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to Create a Table: ", err)
	}

	dataInsertToTable()

	log.Println("Connected to Database Successfully....")
}

func dataInsertToTable() {
	row := DataBase.QueryRow("SELECT COUNT(*) FROM employee")
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Fatal("Failed to Count Rows:", err)
	}
	if count == 0 {
		insertQuery := `INSERT INTO employee (name, emailId, projectName, location, mobileNo) 
						VALUES 
						('Arul Kumar', 'arul@example.com', 'Task Tracker', 'Chennai', '9876543210'),
						('Jane Doe', 'jane@example.com', 'CRM System', 'Bangalore', '9123456789');`
		_, err = DataBase.Exec(insertQuery)
		if err != nil {
			log.Fatal("Failed to Insert Initial Data: ", err)
		}
		log.Println("Inserted Initial Data Into the Employee Table.")
	}
}
