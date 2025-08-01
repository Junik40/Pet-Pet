package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


func PutDB(inf InfUnit){
	db,err := sql.Open("postgres","user=postgres password=123456 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	querry := fmt.Sprintf("UPDATE InfUnit SET ServiceName = '%s', Price = '%d', UserId = '%s', StartDate = '%s', EndDate = '%s' WHERE ID = %d", inf.ServiceName, inf.Price, inf.UserId, inf.StartDate, inf.EndDate, inf.ID)
	_, err = db.Exec(querry)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
}