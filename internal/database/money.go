package database

import (
	"fmt"
	"strconv"
	"strings"
)

func money(startdate string, enddate string, price int) (int, error) {
	start := strings.Split(startdate, "-")
	end := strings.Split(enddate, "-")
	startMonth,err := strconv.Atoi(start[0])
	if err != nil {
		return 0, err
	}
	startYear,err := strconv.Atoi(start[1])
	if err != nil {
		return 0, err
	}
	endMonth,err := strconv.Atoi(end[0])
	if err != nil {
		return 0, err
	}
	endYear,err := strconv.Atoi(end[1])
	if err != nil {
		return 0, err
	}
	summ := 0
	for {
		if startMonth == endMonth && startYear == endYear {
			return summ, nil
		}
		summ = summ + price
		startMonth += 1
		if startMonth == 13 {
			startMonth = 1
			startYear += 1
		}
	}
}

func MoneyDB(userId string) (int, error) {
	db, err := Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT Price, StartDate, EndDate FROM InfUnit WHERE UserId = '%s'", userId)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()
	var price int
	var startDate string
	var endDate string
	var summ int
	for rows.Next() {
		var bufsum int
		err = rows.Scan(&price, &startDate, &endDate)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		bufsum, err := money(startDate, endDate, price)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		summ = summ + bufsum
		fmt.Println(summ)
	}
	return summ, nil


}