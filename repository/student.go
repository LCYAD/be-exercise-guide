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

type studentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *studentRepository {
	return &studentRepository{
		db: db,
	}
}

func (r *studentRepository) GetStudentIDs() []int32 {
	stmt := SELECT(
		Student.ID,
	).FROM(
		Student,
	)

	var dest []model.Student

	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.ID)
	}

	return ids
}

func (r *studentRepository) InsertMultipleStudents(students []model.Student) {
	insertStmt := Student.INSERT(
		Student.FirstName,
		Student.LastName,
		Student.Dob,
		Student.Email,
		Student.DepartmentID,
	).MODELS(students)
	_, err := insertStmt.Exec(r.db)
	util.PanicOnError(err)
}

func (r *studentRepository) ClearAllStudents() {
	_, err := r.db.Exec("TRUNCATE TABLE student RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating student table and reset auto increment")
}
