package seeder

import (
	"database/sql"
	"fmt"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

func DepartmentSeeder(db *sql.DB) {
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
	departmentRepository := repository.NewDepartmentRepository(db)
	var departmentModelLinks []model.Department
	departmentIds := departmentRepository.GetDepartmentIDs()
	if len(departmentIds) == 0 {
		for _, name := range departmentNames {
			now := time.Now().UTC()
			modelLink := model.Department{
				Name:      name,
				CreatedAt: &now,
				UpdatedAt: &now,
			}
			departmentModelLinks = append(departmentModelLinks, modelLink)
		}
		departmentRepository.InsertMultipleDepartments(departmentModelLinks)
		fmt.Println("Finish seeding Department")
	} else {
		fmt.Println("Already created Departments.  Skipping....")
	}
}
