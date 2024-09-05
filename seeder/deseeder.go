package seeder

import (
	"database/sql"

	"be-exerise-go-mod/repository"

	_ "github.com/lib/pq"
)

func DeseedAll(db *sql.DB) {
	repository.ClearAllTeachers(db)
	repository.ClearAllStudents(db)
	repository.ClearAllDepartments(db)
}
