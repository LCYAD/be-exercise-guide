package seeder

import (
	"database/sql"

	"be-exerise-go-mod/repository"

	_ "github.com/lib/pq"
)

func DeseedAll(db *sql.DB) {
	repository.ClearAllStudents(db)
	repository.ClearAllCourses(db)
	repository.ClearAllTeachers(db)
	repository.ClearAllDepartments(db)
}
