package seeder

import (
	"fmt"
	"math/rand"
	"time"

	"be-exerise-go-mod/repository"

	"be-exerise-go-mod/.gen/be-exercise/public/model"

	_ "github.com/lib/pq"
)

type faker interface {
	FirstName() string
	LastName() string
	DateRange(time.Time, time.Time) time.Time
	Email() string
}

type randomizer interface {
	Intn(int) int
}

type teacherSeeder struct {
	teacherRepo    repository.TeacherRepository
	departmentRepo repository.DepartmentRepository
	faker          faker
	randomizer     randomizer
}

func NewTeacherSeeder(
	teacherRepo repository.TeacherRepository,
	departmentRepo repository.DepartmentRepository,
	faker faker,
	randomizer randomizer,
) *teacherSeeder {
	return &teacherSeeder{
		departmentRepo: departmentRepo,
		teacherRepo:    teacherRepo,
		faker:          faker,
		randomizer:     randomizer,
	}
}

func (s *teacherSeeder) Seed(num int32) {
	departmentIds := s.departmentRepo.GetDepartmentIDs()

	var teacherModelLinks []model.Teacher
	for range num {
		now := time.Now().UTC()
		modelLink := model.Teacher{
			FirstName:    s.faker.FirstName(),
			LastName:     s.faker.LastName(),
			Dob:          s.faker.DateRange(now.AddDate(-70, 0, 0), now.AddDate(-25, 0, 0)),
			Email:        s.faker.Email(),
			DepartmentID: &departmentIds[rand.Intn(len(departmentIds))],
		}
		teacherModelLinks = append(teacherModelLinks, modelLink)
	}
	s.teacherRepo.InsertMultipleTeachers(teacherModelLinks)
	fmt.Println("Finish seeding Teachers")
}
