package main

import (
	"database/sql"
	"fmt"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	. "be-exerise-go-mod/.gen/be-exercise/public/table"

	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "be-exercise"
)

func main() {
	gofakeit.Seed(0)
	var departmentNames = []string{"Economic", "Finance", "Computer Science", "Biology", "Chemistry"}

	fmt.Println("Starting uploading to DB")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	panicOnError(err)
	defer db.Close()

	var departmentModelLinks []model.Department
	for _, name := range departmentNames {
		now := time.Now()
		modelLink := model.Department{
			Name:      name,
			CreatedAt: &now,
			UpdatedAt: &now,
		}
		departmentModelLinks = append(departmentModelLinks, modelLink)
	}
	insertStmt := Department.INSERT(Department.Name, Department.CreatedAt, Department.UpdatedAt).MODELS(departmentModelLinks)
	_, err = insertStmt.Exec(db)
	panicOnError(err)
	fmt.Println("Finished uploading to DB")
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
