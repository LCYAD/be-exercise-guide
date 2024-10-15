package repository

import (
	"database/sql"
	"fmt"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	. "be-exerise-go-mod/.gen/be-exercise/public/table"
	"be-exerise-go-mod/util"

	. "github.com/go-jet/jet/v2/postgres"

	_ "github.com/lib/pq"
)

type SubmissionRepository interface {
	GetSubmissionIDsAndDepartmentIDs() []SubmissionRes
	InsertMultipleSubmissions(submissions []model.Submission)
	ClearAllSubmissions()
}

type submissionRepository struct {
	db *sql.DB
}

type SubmissionRes struct {
	ID                int32
	DepartmentID      int32
	SubmittedAt       time.Time
	AssignmentDueDate time.Time
	IsAssignment      bool
}

func NewSubmissionRepository(db *sql.DB) *submissionRepository {
	return &submissionRepository{
		db: db,
	}
}

func (r *submissionRepository) GetSubmissionIDsAndDepartmentIDs() []SubmissionRes {
	var res []SubmissionRes
	var dest []struct {
		model.Submission
		Assignment *model.Assignment `json:"Assignment"`
		Course     *model.Course     `json:"Course"`
	}
	// couldn't get the 2 left join on both assignment and exam with course to work.  Maybe have another look at it when have time
	stmt := SELECT(
		Submission.AllColumns,
		Assignment.AllColumns,
		Course.AllColumns,
	).FROM(
		Submission.
			LEFT_JOIN(Assignment, Submission.AssignmentID.EQ(Assignment.ID)).
			LEFT_JOIN(Course, Assignment.CourseID.EQ(Course.ID)),
	).WHERE(Submission.AssignmentID.IS_NOT_NULL())

	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	for _, c := range dest {
		res = append(res, SubmissionRes{
			ID:                int32(c.ID),
			DepartmentID:      *c.Course.DepartmentID,
			SubmittedAt:       c.SubmittedAt,
			AssignmentDueDate: c.Assignment.DueDate,
			IsAssignment:      true,
		})
	}

	// empty array
	dest = []struct {
		model.Submission
		Assignment *model.Assignment `json:"Assignment"`
		Course     *model.Course     `json:"Course"`
	}{}

	stmt = SELECT(
		Submission.AllColumns,
		Course.AllColumns,
	).FROM(
		Submission.
			LEFT_JOIN(Exam, Submission.ExamID.EQ(Exam.ID)).
			LEFT_JOIN(Course, Exam.CourseID.EQ(Course.ID)),
	).WHERE(Submission.ExamID.IS_NOT_NULL())

	err = stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	for _, c := range dest {
		res = append(res, SubmissionRes{
			ID:                int32(c.ID),
			DepartmentID:      *c.Course.DepartmentID,
			SubmittedAt:       c.SubmittedAt,
			AssignmentDueDate: time.Time{},
			IsAssignment:      false,
		})
	}

	return res
}

func (r *submissionRepository) InsertMultipleSubmissions(submissions []model.Submission) {
	insertStmt := Submission.INSERT(
		Submission.StudentID,
		Submission.AssignmentID,
		Submission.ExamID,
		Submission.SubmittedAt,
	).MODELS(submissions)
	_, err := insertStmt.Exec(r.db)
	util.PanicOnError(err)
}

func (r *submissionRepository) ClearAllSubmissions() {
	_, err := r.db.Exec("TRUNCATE TABLE submission RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating submission table and reset auto increment")
}
