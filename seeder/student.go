package seeder

import (
	"fmt"
	"time"

	"be-exerise-go-mod/repository"

	"be-exerise-go-mod/.gen/be-exercise/public/model"

	_ "github.com/lib/pq"
)

type studentSeeder struct {
	studentRepo    repository.StudentRepository
	departmentRepo repository.DepartmentRepository
	faker          faker
	randomizer     randomizer
}

func NewStudentSeeder(
	studentRepo repository.StudentRepository,
	departmentRepo repository.DepartmentRepository,
	faker faker,
	randomizer randomizer,
) *studentSeeder {
	return &studentSeeder{
		studentRepo:    studentRepo,
		departmentRepo: departmentRepo,
		faker:          faker,
		randomizer:     randomizer,
	}
}

func (s *studentSeeder) Seed(num int32) {
	var departmentIds = s.departmentRepo.GetDepartmentIDs()

	var studentModelLinks []model.Student
	for range num {
		now := time.Now().UTC()
		modelLink := model.Student{
			FirstName:    s.faker.FirstName(),
			LastName:     s.faker.LastName(),
			Dob:          s.faker.DateRange(now.AddDate(-50, 0, 0), now.AddDate(-20, 0, 0)),
			Email:        s.faker.Email(),
			DepartmentID: &departmentIds[s.randomizer.Intn(len(departmentIds))],
		}
		studentModelLinks = append(studentModelLinks, modelLink)
	}

	s.studentRepo.InsertMultipleStudents(studentModelLinks)
	fmt.Println("Finish seeding Students")
}
