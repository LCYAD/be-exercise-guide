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

// TODO: anyway to extract model.Student.ID type?
func StudentEnrolledInCourse(db *sql.DB, studentID int32, courseID int32) bool {
	stmt := SELECT(
		Enrollment.ID,
	).FROM(
		Enrollment,
	).WHERE(Enrollment.StudentID.EQ(Int32(studentID)).AND(Enrollment.CourseID.EQ(Int32(courseID)))).LIMIT(1)

	var dest []model.Enrollment

	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	return len(dest) > 0
}
func InsertMultipleEnrollments(db *sql.DB, enrollments []model.Enrollment) {
	insertStmt := Enrollment.INSERT(
		Enrollment.StudentID,
		Enrollment.CourseID,
		Enrollment.Approved,
		Enrollment.CreatedAt,
		Enrollment.UpdatedAt,
	).MODELS(enrollments)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllEnrollments(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE enrollment RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating enrollment table and reset auto increment")
}
