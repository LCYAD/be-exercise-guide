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

func GetScoreIDs(db *sql.DB) []int32 {
	stmt := SELECT(
		Score.ID,
	).FROM(
		Score,
	)

	var dest []model.Score

	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.ID)
	}

	return ids
}

func InsertMultipleScores(db *sql.DB, scores []model.Score) {
	insertStmt := Score.INSERT(Score.Value, Score.TeacherID, Score.SubmissionID, Score.CreatedAt, Score.UpdatedAt).MODELS(scores)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllScores(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE score RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating score table and reset auto increment")
}
