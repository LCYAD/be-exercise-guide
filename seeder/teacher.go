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

type teacherSeeder struct {
	db *sql.DB
}

func NewTeacherSeeder(db *sql.DB) *teacherSeeder {
	return &teacherSeeder{db: db}
}

func (s *teacherSeeder) TeacherSeeder(num int32) {
	var departmentIds = repository.GetDepartmentIDs(s.db)

	var teacherModelLinks []model.Teacher
	for range num {
		now := time.Now().UTC()
		modelLink := model.Teacher{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			Dob:          gofakeit.DateRange(now.AddDate(-70, 0, 0), now.AddDate(-25, 0, 0)),
			Email:        gofakeit.Email(),
			DepartmentID: &departmentIds[rand.Intn(len(departmentIds))],
			CreatedAt:    &now,
			UpdatedAt:    &now,
		}
		teacherModelLinks = append(teacherModelLinks, modelLink)
	}
	teacherRepository := repository.NewTeacherRepository(s.db)
	teacherRepository.InsertMultipleTeachers(teacherModelLinks)
	fmt.Println("Finish seeding Teachers")
}
