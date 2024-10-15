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

type ExamRepository interface {
	GetExamIDs() []int32
	GetExamsByCourseID(courseID int32) []model.Exam
	InsertMultipleExams(exams []model.Exam)
	ClearAllExams()
}

type examRepository struct {
	db *sql.DB
}

func NewExamRepository(db *sql.DB) *examRepository {
	return &examRepository{
		db: db,
	}
}

func (r *examRepository) GetExamIDs() []int32 {
	stmt := SELECT(
		Exam.ID,
	).FROM(
		Exam,
	)

	var dest []model.Exam

	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.ID)
	}

	return ids
}

func (r *examRepository) GetExamsByCourseID(courseID int32) []model.Exam {
	stmt := SELECT(
		Exam.AllColumns,
	).FROM(
		Exam,
	).WHERE(Exam.CourseID.EQ(Int32(courseID)))

	var dest []model.Exam

	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	return dest
}

func (r *examRepository) InsertMultipleExams(exams []model.Exam) {
	insertStmt := Exam.INSERT(
		Exam.Name,
		Exam.Type,
		Exam.StartedAt,
		Exam.FinishedAt,
		Exam.CourseID,
	).MODELS(exams)
	_, err := insertStmt.Exec(r.db)
	util.PanicOnError(err)
}

func (r *examRepository) ClearAllExams() {
	_, err := r.db.Exec("TRUNCATE TABLE exam RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating exam table and reset auto increment")
}
