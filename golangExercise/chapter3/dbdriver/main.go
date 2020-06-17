package main

import (
	"database/sql"
	_ "goStudyCode/golangExercise/chapter3/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
}
