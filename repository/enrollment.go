package repository

import (
	"database/sql"
	"fmt"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	. "be-exerise-go-mod/.gen/be-exercise/public/table"
	"be-exerise-go-mod/util"

	. "github.com/go-jet/jet/v2/postgres"

	_ "github.com/lib/pq"
)

type enrollmentRepository struct {
	db *sql.DB
}

func NewEnrollmentRepository(db *sql.DB) *enrollmentRepository {
	return &enrollmentRepository{
		db: db,
	}
}

// TODO: anyway to extract model.Student.ID type?
func (r *enrollmentRepository) IsStudentEnrolledInCourse(studentID int32, courseID int32) bool {
	stmt := SELECT(
		Enrollment.ID,
	).FROM(
		Enrollment,
	).WHERE(Enrollment.StudentID.EQ(Int32(studentID)).AND(Enrollment.CourseID.EQ(Int32(courseID)))).LIMIT(1)

	var dest []model.Enrollment

	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	return len(dest) > 0
}

func (r *enrollmentRepository) GetStudentIDsEnrolledInCourse(courseID int32) []int32 {
	stmt := SELECT(
		Enrollment.StudentID,
	).FROM(
		Enrollment,
	).WHERE(Enrollment.CourseID.EQ(Int32(courseID)).AND(Enrollment.Approved.EQ(Bool(true))))

	var dest []model.Enrollment

	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	studentIds := make([]int32, len(dest))
	for i, d := range dest {
		studentIds[i] = int32(*d.StudentID)
	}

	return studentIds
}

func (r *enrollmentRepository) InsertMultipleEnrollments(enrollments []model.Enrollment) {
	insertStmt := Enrollment.INSERT(
		Enrollment.StudentID,
		Enrollment.CourseID,
		Enrollment.Approved,
	).MODELS(enrollments)
	_, err := insertStmt.Exec(r.db)
	util.PanicOnError(err)
}

func (r *enrollmentRepository) ClearAllEnrollments() {
	_, err := r.db.Exec("TRUNCATE TABLE enrollment RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating enrollment table and reset auto increment")
}
