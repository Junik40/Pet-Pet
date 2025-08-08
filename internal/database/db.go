package database

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)


var db *sql.DB

type InfUnit struct{
	ID int `json:"ID"`
	ServiceName string `json:"ServiceName"`
	Price int	`json:"Price"`
	UserId string `json:"UserId"`
	StartDate string `json:"StartDate"`
	EndDate string `json:"EndDate"`
}

func Open() (*sql.DB,error) {
	db, err := sql.Open("postgres","user=postgres password=123456 dbname=postgres sslmode=disable")
	return db, err
}

func CreateDB() {
	dbFile:= "./internal/DB/database.db"
	_,err := os.Stat(dbFile)
	if err == nil {
		return
	}

	db,err = Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	

	
	res, err := db.Exec("CREATE TABLE IF NOT EXISTS InfUnit (ID SERIAL PRIMARY KEY, ServiceName VARCHAR(256), Price INTEGER, UserId VARCHAR(256) , StartDate CHAR(7), EndDate CHAR(7))")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(res)
}


	


