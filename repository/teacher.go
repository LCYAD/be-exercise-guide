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

func GetAllTeachers(db *sql.DB) []model.Teacher {
	stmt := SELECT(
		Teacher.AllColumns,
	).FROM(
		Teacher,
	)

	var dest []model.Teacher
	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	return dest
}

func InsertMultipleTeachers(db *sql.DB, teachers []model.Teacher) {
	insertStmt := Teacher.INSERT(
		Teacher.FirstName,
		Teacher.LastName,
		Teacher.Dob,
		Teacher.Email,
		Teacher.DepartmentID,
		Teacher.CreatedAt,
		Teacher.UpdatedAt,
	).MODELS(teachers)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllTeachers(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE teacher RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating teacher table and reset auto increment")
}
