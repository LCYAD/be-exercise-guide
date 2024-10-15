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

type AssignmentRepository interface {
	GetAssignmentIDs() []int32
	GetAssignmentsByCourseID(courseID int32) []model.Assignment
	InsertMultipleAssignments(assignments []model.Assignment)
	ClearAllAssignments()
}

type assignmentRepository struct {
	db *sql.DB
}

func NewAssignmentRepository(db *sql.DB) *assignmentRepository {
	return &assignmentRepository{
		db: db,
	}
}

func (r *assignmentRepository) GetAssignmentIDs() []int32 {
	stmt := SELECT(
		Assignment.ID,
	).FROM(
		Assignment,
	)

	var dest []model.Assignment

	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.ID)
	}

	return ids
}

func (r *assignmentRepository) GetAssignmentsByCourseID(courseID int32) []model.Assignment {
	stmt := SELECT(
		Assignment.AllColumns,
	).FROM(
		Assignment,
	).WHERE(Assignment.CourseID.EQ(Int32(courseID)))

	var dest []model.Assignment

	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	return dest
}

func (r *assignmentRepository) InsertMultipleAssignments(assignments []model.Assignment) {
	insertStmt := Assignment.INSERT(
		Assignment.Title,
		Assignment.Description,
		Assignment.Type,
		Assignment.DueDate,
		Assignment.Graded,
		Assignment.CourseID,
	).MODELS(assignments)
	_, err := insertStmt.Exec(r.db)
	util.PanicOnError(err)
}

func (r *assignmentRepository) ClearAllAssignments() {
	_, err := r.db.Exec("TRUNCATE TABLE assignment RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating assignment table and reset auto increment")
}
