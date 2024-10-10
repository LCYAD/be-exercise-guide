package seeder

import (
	"database/sql"

	"be-exerise-go-mod/repository"

	_ "github.com/lib/pq"
)

func DeseedAll(db *sql.DB) {
	// TODO maybe just pass along repository struct pointer
	assignmentRepo := repository.NewAssignmentRepository(db)
	teacherRepository := repository.NewTeacherRepository(db)
	studentRepository := repository.NewStudentRepository(db)
	departmentRepository := repository.NewDepartmentRepository(db)
	courseRepository := repository.NewCourseRepository(db)
	enrollmentRepository := repository.NewEnrollmentRepository(db)
	gradeSettingRepository := repository.NewGradeSettingRepository(db)

	repository.ClearAllScores(db)
	repository.ClearAllSubmissions(db)
	repository.ClearAllExams(db)
	assignmentRepo.ClearAllAssignments()
	enrollmentRepository.ClearAllEnrollments()
	studentRepository.ClearAllStudents()
	gradeSettingRepository.ClearAllGradeSettings()
	courseRepository.ClearAllCourses()
	teacherRepository.ClearAllTeachers()
	departmentRepository.ClearAllDepartments()
}
