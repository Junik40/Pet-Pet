package main

import (
	database "Pet-Pet/internal/database"
	server "Pet-Pet/internal/server"
)

func main(){
	database.CreateDB()
	server.Init()
}