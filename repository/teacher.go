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

func GetTeacherIDs(db *sql.DB) []int32 {
	stmt := SELECT(
		Teacher.ID,
	).FROM(
		Teacher,
	)

	var dest []struct {
		model.Teacher
	}

	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.Teacher.ID)
	}

	return ids
}

func InsertMultipleTeachers(db *sql.DB, teachers []model.Teacher) {
	insertStmt := Teacher.INSERT(Teacher.FirstName, Teacher.LastName, Teacher.Dob, Teacher.Email, Teacher.DepartmentID, Teacher.CreatedAt, Teacher.UpdatedAt).MODELS(teachers)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllTeachers(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE teacher RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating teacher table and reset auto increment")
}
