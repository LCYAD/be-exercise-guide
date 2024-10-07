package seeder

import (
	"database/sql"

	"be-exerise-go-mod/repository"

	_ "github.com/lib/pq"
)

func DeseedAll(db *sql.DB) {
	repository.ClearAllScores(db)
	repository.ClearAllSubmissions(db)
	repository.ClearAllExams(db)
	// TODO maybe just pass along repostiroy struct pointer
	assignmentRepo := repository.NewAssignmentRepository(db)
	assignmentRepo.ClearAllAssignments()
	repository.ClearAllEnrollments(db)
	repository.ClearAllStudents(db)
	repository.ClearAllGradeSettings(db)
	repository.ClearAllCourses(db)
	repository.ClearAllTeachers(db)
	repository.ClearAllDepartments(db)
}
