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

func TeacherSeeder(db *sql.DB, num int32) {
	var departmentIds = repository.GetDepartmentIDs(db)

	var teacherModelLinks []model.Teacher
	for range num {
		now := time.Now()
		modelLink := model.Teacher{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			Dob:          gofakeit.DateRange(time.Now().AddDate(-70, 0, 0), time.Now().AddDate(-25, 0, 0)),
			Email:        gofakeit.Email(),
			DepartmentID: &departmentIds[rand.Intn(len(departmentIds))],
			CreatedAt:    &now,
			UpdatedAt:    &now,
		}
		teacherModelLinks = append(teacherModelLinks, modelLink)
	}
	repository.InsertMultipleTeachers(db, teacherModelLinks)
	fmt.Println("Finish seeding Teachers")
}
