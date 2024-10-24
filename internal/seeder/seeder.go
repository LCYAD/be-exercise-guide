package seeder

import (
	"be-exerise-go-mod/repository"
	"be-exerise-go-mod/seeder"
	"database/sql"
	"math/rand"

	"github.com/brianvoe/gofakeit/v7"
)

type internalSeeder struct {
	assignmentSeeder   seeder.AssignmentSeeder
	courseSeeder       seeder.CoursetSeeder
	departmentSeeder   seeder.DepartmentSeeder
	enrollmentSeeder   seeder.EnrollmentSeeder
	examSeeder         seeder.ExamSeeder
	gradeSettingSeeder seeder.GradeSettingSeeder
	scoreSeeder        seeder.ScoreSeeder
	studentSeeder      seeder.StudentSeeder
	submissionSeeder   seeder.SubmissionSeeder
	teacherSeeder      seeder.TeacherSeeder
}

func NewSeeder(db *sql.DB) *internalSeeder {
	faker := gofakeit.New(0)
	randomizer := rand.New(rand.NewSource(0))

	departmentRepo := repository.NewDepartmentRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	assignmentRepo := repository.NewAssignmentRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)
	examRepo := repository.NewExamRepository(db)
	submissionRepo := repository.NewSubmissionRepository(db)
	scoreRepo := repository.NewScoreRepository(db)
	gradeSettingRepo := repository.NewGradeSettingRepository(db)

	teacherSeeder := seeder.NewTeacherSeeder(teacherRepo, departmentRepo, faker, randomizer)
	departmentSeeder := seeder.NewDepartmentSeeder(departmentRepo)
	studentSeeder := seeder.NewStudentSeeder(studentRepo, departmentRepo, faker, randomizer)
	courseSeeder := seeder.NewCourseSeeder(courseRepo, departmentRepo, teacherRepo)
	enrollmentSeeder := seeder.NewEnrollmentRepository(enrollmentRepo, studentRepo, courseRepo)
	assignmentSeeder := seeder.NewAssignmentSeeder(assignmentRepo, courseRepo, faker, randomizer)
	examSeeder := seeder.NewExamSeeder(examRepo, courseRepo)
	submissionSeeder := seeder.NewSubmissionSeeder(submissionRepo, courseRepo, enrollmentRepo, assignmentRepo, examRepo)
	scoreSeeder := seeder.NewScoreSeeder(scoreRepo, teacherRepo, submissionRepo)
	gradeSettingSeeder := seeder.NewGradeSettingSeeder(gradeSettingRepo, courseRepo)

	return &internalSeeder{
		assignmentSeeder:   assignmentSeeder,
		courseSeeder:       courseSeeder,
		departmentSeeder:   departmentSeeder,
		enrollmentSeeder:   enrollmentSeeder,
		examSeeder:         examSeeder,
		gradeSettingSeeder: gradeSettingSeeder,
		scoreSeeder:        scoreSeeder,
		studentSeeder:      studentSeeder,
		submissionSeeder:   submissionSeeder,
		teacherSeeder:      teacherSeeder,
	}
}

func (s *internalSeeder) SeedAll(tSize int32, sSize int32) {
	s.departmentSeeder.Seed()
	s.teacherSeeder.Seed(tSize)
	s.courseSeeder.Seed()
	s.gradeSettingSeeder.Seed()
	s.studentSeeder.Seed(sSize)
	s.enrollmentSeeder.Seed()
	s.assignmentSeeder.Seed()
	s.examSeeder.Seed()
	s.submissionSeeder.Seed()
	s.scoreSeeder.Seed()
}

func (s *internalSeeder) DeseedAll() {
	s.scoreSeeder.Deseed()
	s.submissionSeeder.Deseed()
	s.examSeeder.Deseed()
	s.assignmentSeeder.Deseed()
	s.enrollmentSeeder.Deseed()
	s.studentSeeder.Deseed()
	s.gradeSettingSeeder.Deseed()
	s.courseSeeder.Deseed()
	s.teacherSeeder.Deseed()
	s.departmentSeeder.Deseed()
}
