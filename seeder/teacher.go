package seeder

import (
	"fmt"
	"math/rand"
	"time"

	"be-exerise-go-mod/repository"

	"be-exerise-go-mod/.gen/be-exercise/public/model"

	"github.com/brianvoe/gofakeit/v7"

	_ "github.com/lib/pq"
)

type teacherSeeder struct {
	teacherRepo    repository.TeacherRepository
	departmentRepo repository.DepartmentRepository
}

func NewTeacherSeeder(teacherRepo repository.TeacherRepository, departmentRepo repository.DepartmentRepository) *teacherSeeder {
	return &teacherSeeder{
		departmentRepo: departmentRepo,
		teacherRepo:    teacherRepo,
	}
}

func (s *teacherSeeder) Seed(num int32) {
	departmentIds := s.departmentRepo.GetDepartmentIDs()

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
	s.teacherRepo.InsertMultipleTeachers(teacherModelLinks)
	fmt.Println("Finish seeding Teachers")
}
