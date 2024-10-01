package repository

import (
	"database/sql"
	"fmt"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	. "be-exerise-go-mod/.gen/be-exercise/public/table"
	"be-exerise-go-mod/util"

	_ "github.com/lib/pq"
)

func InsertMultipleGradeSettings(db *sql.DB, gradeSettings []model.GradeSetting) {
	insertStmt := GradeSetting.INSERT(
		GradeSetting.AssignmentPercent,
		GradeSetting.ExamPercent,
		GradeSetting.PassingGrade,
		GradeSetting.CourseID,
		GradeSetting.CreatedAt,
		GradeSetting.UpdatedAt,
	).MODELS(gradeSettings)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllGradeSettings(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE grade_setting RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating grade_setting table and reset auto increment")
}
