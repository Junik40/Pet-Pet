package database

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)


var db *sql.DB

type InfUnit struct{
	ServiceName string `json:"ServiceName"`
	Price int	`json:"Price"`
	UserId string `json:"UserId"`
	StartDate string `json:"StartDate"`
	EndDate string `json:"EndDate"`
}

func CreateDB() {
	dbFile:= "./internal/DB/database.db"
	_,err := os.Stat(dbFile)
	if err == nil {
		return
	}
	db,err = sql.Open("postgres","user=postgres password=123456 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	
	res, err := db.Exec("CREATE TABLE IF NOT EXISTS InfUnit (ServiceName VARCHAR(256), Price INTEGER, UserId VARCHAR(256) , StartDate CHAR(7), EndDate CHAR(7))")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
	fmt.Println(res)
}


	


