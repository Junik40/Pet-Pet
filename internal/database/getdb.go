package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDB() []InfUnit {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM InfUnit")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var units []InfUnit
	for rows.Next() {
		unit := InfUnit{}
		err := rows.Scan(&unit.ServiceName, &unit.Price, &unit.UserId, &unit.StartDate, &unit.EndDate)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		units = append(units, unit)
	}
	return units
}

func GetIdDB(id string) []InfUnit {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	querry := fmt.Sprintf("SELECT * FROM InfUnit WHERE UserId = '%s'", id)
	rows, err := db.Query(querry)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
	var units []InfUnit
	for rows.Next() {
		unit := InfUnit{}
		err := rows.Scan(&unit.ServiceName, &unit.Price, &unit.UserId, &unit.StartDate, &unit.EndDate)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		units = append(units, unit)
	}
	return units
}

func GetServiceNameDB(name string) []InfUnit {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	querry := fmt.Sprintf("SELECT * FROM InfUnit WHERE ServiceName = '%s'", name)
	rows, err := db.Query(querry)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var units []InfUnit
	for rows.Next() {
		unit := InfUnit{}
		err := rows.Scan(&unit.ServiceName, &unit.Price, &unit.UserId, &unit.StartDate, &unit.EndDate)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		units = append(units, unit)
	}
	return units
}
