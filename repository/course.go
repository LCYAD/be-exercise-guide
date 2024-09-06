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

func GetCourseIDs(db *sql.DB) []int32 {
	stmt := SELECT(
		Course.ID,
	).FROM(
		Course,
	)

	var dest []model.Course

	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.ID)
	}

	return ids
}

func CourseExists(db *sql.DB) bool {
	stmt := SELECT(
		Course.ID,
	).FROM(
		Course,
	).LIMIT(1)

	var dest []model.Course

	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	return len(dest) > 0
}

func InsertMultipleCourses(db *sql.DB, departments []model.Course) {
	insertStmt := Course.INSERT(
		Course.Name,
		Course.Description,
		Course.DepartmentID,
		Course.TeacherID,
		Course.CreatedAt,
		Course.UpdatedAt,
	).MODELS(departments)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllCourses(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE course RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating course table and reset auto increment")
}
