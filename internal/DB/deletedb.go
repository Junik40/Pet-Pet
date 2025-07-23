package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func DeleteDB(id string){
	db,err := sql.Open("postgres","user=postgres password=123456 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	querry := fmt.Sprintf("DELETE FROM InfUnit WHERE UserId = '%s'", id)
	_, err  = db.Exec(querry)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

}