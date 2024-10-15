package repository

import (
	"database/sql"
	"fmt"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	. "be-exerise-go-mod/.gen/be-exercise/public/table"
	"be-exerise-go-mod/util"

	_ "github.com/lib/pq"
)

type GradeSettingRepository interface {
	InsertMultipleGradeSettings(gradeSettings []model.GradeSetting)
	ClearAllGradeSettings()
}

type gradeSettingRepository struct {
	db *sql.DB
}

func NewGradeSettingRepository(db *sql.DB) *gradeSettingRepository {
	return &gradeSettingRepository{
		db: db,
	}
}

func (r *gradeSettingRepository) InsertMultipleGradeSettings(gradeSettings []model.GradeSetting) {
	insertStmt := GradeSetting.INSERT(
		GradeSetting.AssignmentPercent,
		GradeSetting.ExamPercent,
		GradeSetting.PassingGrade,
		GradeSetting.CourseID,
	).MODELS(gradeSettings)
	_, err := insertStmt.Exec(r.db)
	util.PanicOnError(err)
}

func (r *gradeSettingRepository) ClearAllGradeSettings() {
	_, err := r.db.Exec("TRUNCATE TABLE grade_setting RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating grade_setting table and reset auto increment")
}
