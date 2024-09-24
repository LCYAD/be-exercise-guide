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

func GetAssignmentIDs(db *sql.DB) []int32 {
	stmt := SELECT(
		Assignment.ID,
	).FROM(
		Assignment,
	)

	var dest []model.Assignment

	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.ID)
	}

	return ids
}

func InsertMultipleAssignments(db *sql.DB, assignments []model.Assignment) {
	insertStmt := Assignment.INSERT(
		Assignment.Title,
		Assignment.Description,
		Assignment.Type,
		Assignment.DueDate,
		Assignment.Graded,
		Assignment.CourseID,
		Assignment.CreatedAt,
		Assignment.UpdatedAt,
	).MODELS(assignments)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllAssignments(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE assignment RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating assignment table and reset auto increment")
}
