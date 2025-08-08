package database

import (
	"fmt"
	_ "github.com/lib/pq"
)

func DeleteDB(id int){
	db,err := Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	querry := fmt.Sprintf("DELETE FROM InfUnit WHERE ID = '%d'", id)
	_, err  = db.Exec(querry)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
}

func DeleteUuidDB(id string){
	db,err := Open()
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