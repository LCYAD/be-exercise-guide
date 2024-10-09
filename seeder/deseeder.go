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
	// TODO maybe just pass along repository struct pointer
	assignmentRepo := repository.NewAssignmentRepository(db)
	teacherRepository := repository.NewTeacherRepository(db)
	studentRepository := repository.NewStudentRepository(db)
	assignmentRepo.ClearAllAssignments()
	repository.ClearAllEnrollments(db)
	studentRepository.ClearAllStudents()
	repository.ClearAllGradeSettings(db)
	repository.ClearAllCourses(db)
	teacherRepository.ClearAllTeachers()
	repository.ClearAllDepartments(db)
}
