package main

import (
	"be-exerise-go-mod/internal/seeder"
	"be-exerise-go-mod/util"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host        = "localhost"
	port        = 5432
	user        = "root"
	password    = "root"
	dbname      = "be-exercise"
	teacherSize = 100
	studentSize = 1000
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide 'up' or 'down' as an argument")
		return
	}

	command := os.Args[1]

	fmt.Println("Starting connection to DB")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	util.PanicOnError(err)
	defer db.Close()

	s := seeder.NewSeeder(db)

	switch command {
	case "up":
		fmt.Println("----- Starting running seeders -----")
		s.SeedAll(teacherSize, studentSize)
		fmt.Println("----- Finished running seeder -----")
	case "down":
		fmt.Println("----- Starting running deseeder -----")
		s.DeseedAll()
		fmt.Println("----- Complete running deseeder -----")
	default:
		fmt.Println("Wrong command. Please use 'up' or 'down'")
	}
}
