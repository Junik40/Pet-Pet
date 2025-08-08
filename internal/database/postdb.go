package database

import (
	"fmt"
	_ "github.com/lib/pq"
)
func PostDB(unit InfUnit){
	db,err := Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	querry := fmt.Sprintf("INSERT INTO InfUnit (ServiceName, Price, UserId, StartDate, EndDate) VALUES ('%s', %d, '%s', '%s', '%s')", unit.ServiceName, unit.Price, unit.UserId, unit.StartDate, unit.EndDate)
	_, err = db.Exec(querry)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
}