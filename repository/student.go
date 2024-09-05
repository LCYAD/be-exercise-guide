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

func GetStudentIDs(db *sql.DB) []int32 {
	stmt := SELECT(
		Student.ID,
	).FROM(
		Student,
	)

	var dest []struct {
		model.Student
	}

	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.Student.ID)
	}

	return ids
}

func InsertMultipleStudents(db *sql.DB, students []model.Student) {
	insertStmt := Student.INSERT(Student.FirstName, Student.LastName, Student.Dob, Student.Email, Student.DepartmentID, Student.CreatedAt, Student.UpdatedAt).MODELS(students)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllStudents(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE student RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating student table and reset auto increment")
}
