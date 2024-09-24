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

func GetSubmissionIDs(db *sql.DB) []int32 {
	stmt := SELECT(
		Submission.ID,
	).FROM(
		Submission,
	)

	var dest []model.Submission

	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.ID)
	}

	return ids
}

func InsertMultipleSubmissions(db *sql.DB, submissions []model.Submission) {
	insertStmt := Submission.INSERT(Submission.StudentID, Submission.AssignmentID, Submission.ExamID, Submission.CreatedAt, Submission.UpdatedAt).MODELS(submissions)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllSubmissions(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE submission RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating submission table and reset auto increment")
}
