package seeder

import (
	"database/sql"
	"time"
	"fmt"

	"be-exerise-go-mod/repository"
	"be-exerise-go-mod/.gen/be-exercise/public/model"
)

func DepartmentSeeder(db *sql.DB) {
	var departmentNames = []string{"Economic", "Finance", "Computer Science", "Biology", "Chemistry"}
	var departmentModelLinks []model.Department
	var departementIds = repository.GetDepartmentIDs(db)
	if len(departementIds) == 0 {
		for _, name := range departmentNames {
			now := time.Now()
			modelLink := model.Department{
				Name:      name,
				CreatedAt: &now,
				UpdatedAt: &now,
			}
			departmentModelLinks = append(departmentModelLinks, modelLink)
		}
		repository.InsertMultipleDepartments(db, departmentModelLinks)
		fmt.Println("Finish seeding Department")
	} else {
		fmt.Println("Already created Departments.  Skipping....")
	}
}