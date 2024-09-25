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

func GetExamIDs(db *sql.DB) []int32 {
	stmt := SELECT(
		Exam.ID,
	).FROM(
		Exam,
	)

	var dest []model.Exam

	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.ID)
	}

	return ids
}

func InsertMultipleExams(db *sql.DB, exams []model.Exam) {
	insertStmt := Exam.INSERT(
		Exam.Name,
		Exam.Type,
		Exam.StartedAt,
		Exam.FinishedAt,
		Exam.CourseID,
		Exam.CreatedAt,
		Exam.UpdatedAt,
	).MODELS(exams)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllExams(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE exam RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating exam table and reset auto increment")
}
