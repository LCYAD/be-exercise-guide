package seeder

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"be-exerise-go-mod/repository"

	"be-exerise-go-mod/.gen/be-exercise/public/model"

	"github.com/brianvoe/gofakeit/v7"

	_ "github.com/lib/pq"
)

func StudentSeeder(db *sql.DB, num int32) {
	var departmentIds = repository.GetDepartmentIDs(db)

	var studentModelLinks []model.Student
	for range num {
		now := time.Now().UTC()
		modelLink := model.Student{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			Dob:          gofakeit.DateRange(now.AddDate(-50, 0, 0), now.AddDate(-20, 0, 0)),
			Email:        gofakeit.Email(),
			DepartmentID: &departmentIds[rand.Intn(len(departmentIds))],
			CreatedAt:    &now,
			UpdatedAt:    &now,
		}
		studentModelLinks = append(studentModelLinks, modelLink)
	}
	studentRepository := repository.NewStudentRepository(db)
	studentRepository.InsertMultipleStudents(studentModelLinks)
	fmt.Println("Finish seeding Students")
}
