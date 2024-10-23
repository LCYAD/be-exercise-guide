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

		faker := gofakeit.New(0)
		randomizer := rand.New(rand.NewSource(0))

		departmentRepo := repository.NewDepartmentRepository(db)
		teacherRepo := repository.NewTeacherRepository(db)
		studentRepo := repository.NewStudentRepository(db)
		assignmentRepo := repository.NewAssignmentRepository(db)
		courseRepo := repository.NewCourseRepository(db)
		enrollmentRepo := repository.NewEnrollmentRepository(db)
		examRepo := repository.NewExamRepository(db)
		submissionRepo := repository.NewSubmissionRepository(db)
		scoreRepo := repository.NewScoreRepository(db)
		gradeSettingRepo := repository.NewGradeSettingRepository(db)

		teacherSeeder := seeder.NewTeacherSeeder(teacherRepo, departmentRepo, faker, randomizer)
		departmentSeeder := seeder.NewDepartmentSeeder(departmentRepo)
		studentSeeder := seeder.NewStudentSeeder(studentRepo, departmentRepo, faker, randomizer)
		courseSeeder := seeder.NewCourseSeeder(courseRepo, departmentRepo, teacherRepo)
		enrollmentSeeder := seeder.NewEnrollmentRepository(enrollmentRepo, studentRepo, courseRepo)
		assignmentSeeder := seeder.NewAssignmentSeeder(assignmentRepo, courseRepo, faker, randomizer)
		examSeeder := seeder.NewExamSeeder(examRepo, courseRepo)
		submissionSeeder := seeder.NewSubmissionSeeder(submissionRepo, courseRepo, enrollmentRepo, assignmentRepo, examRepo)
		scoreSeeder := seeder.NewScoreSeeder(scoreRepo, teacherRepo, submissionRepo)
		gradeSettingSeeder := seeder.NewGradeSettingSeeder(gradeSettingRepo, courseRepo)

		departmentSeeder.Seed()
		teacherSeeder.Seed(teacherSize)
		courseSeeder.Seed()
		gradeSettingSeeder.Seed()
		studentSeeder.Seed(studentSize)
		enrollmentSeeder.Seed()
		assignmentSeeder.Seed()
		examSeeder.Seed()
		submissionSeeder.Seed()
		scoreSeeder.Seed()
		fmt.Println("----- Finished running seeder -----")
	case "down":
		fmt.Println("----- Starting running deseeder -----")
		seeder.DeseedAll(db)
		fmt.Println("----- Complete running deseeder -----")
	default:
		fmt.Println("Wrong command. Please use 'up' or 'down'")
	}
}
