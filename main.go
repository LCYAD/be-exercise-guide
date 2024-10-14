package main

import (
	"be-exerise-go-mod/repository"
	"be-exerise-go-mod/seeder"
	"be-exerise-go-mod/util"
	"database/sql"
	"fmt"
	"math/rand"
	"os"

	"github.com/brianvoe/gofakeit/v7"
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

	switch command {
	case "up":
		fmt.Println("----- Starting running seeders -----")
		departmentRepo := repository.NewDepartmentRepository(db)
		teacherRepo := repository.NewTeacherRepository(db)

		teacherSeeder := seeder.NewTeacherSeeder(teacherRepo, departmentRepo, gofakeit.New(0), rand.New(rand.NewSource(0)))
		departmentSeeder := seeder.NewDepartmentSeeder(departmentRepo)

		departmentSeeder.Seed()
		teacherSeeder.Seed(teacherSize)
		seeder.CourseSeeder(db)
		seeder.GradeSettingSeeder(db)
		seeder.StudentSeeder(db, studentSize)
		seeder.EnrollmentSeeder(db)
		seeder.AssignmentSeeder(db)
		seeder.ExamSeeder(db)
		seeder.SubmissionSeeder(db)
		seeder.ScoreSeeder(db)
		fmt.Println("----- Finished running seeder -----")
	case "down":
		fmt.Println("----- Starting running deseeder -----")
		seeder.DeseedAll(db)
		fmt.Println("----- Complete running deseeder -----")
	default:
		fmt.Println("Wrong command. Please use 'up' or 'down'")
	}
}
