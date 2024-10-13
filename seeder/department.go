package seeder

import (
	"fmt"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

type departmentSeeder struct {
	departmentRepo repository.DepartmentRepository
}

func NewDepartmentSeeder(departmentRepo repository.DepartmentRepository) *departmentSeeder {
	return &departmentSeeder{
		departmentRepo: departmentRepo,
	}
}

func (s *departmentSeeder) Seed() {
	departmentNames := []string{
		"Computer Science",
		"Biology",
		"Chemistry",
		"Physics",
		"Mathematics",
		"Economics",
		"English Literature",
		"History",
		"Psychology",
		"Political Science",
	}

	var departmentModelLinks []model.Department
	departmentIds := s.departmentRepo.GetDepartmentIDs()
	now := time.Now().UTC()
	if len(departmentIds) == 0 {
		for _, name := range departmentNames {
			modelLink := model.Department{
				Name:      name,
				CreatedAt: &now,
				UpdatedAt: &now,
			}
			departmentModelLinks = append(departmentModelLinks, modelLink)
		}
		s.departmentRepo.InsertMultipleDepartments(departmentModelLinks)
		fmt.Println("Finish seeding Department")
	} else {
		fmt.Println("Already created Departments.  Skipping....")
	}
}
