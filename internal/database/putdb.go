package database

func PutDB(inf InfUnit){
	DeleteDB(inf.ID)
	PostDB(inf)
}