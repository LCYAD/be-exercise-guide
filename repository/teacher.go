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

type TeacherRepository interface {
	GetAllTeachers() []model.Teacher
	InsertMultipleTeachers(teachers []model.Teacher)
	ClearAllTeachers()
}

type teacherRepository struct {
	db *sql.DB
}

func NewTeacherRepository(db *sql.DB) *teacherRepository {
	return &teacherRepository{
		db: db,
	}
}

func (r *teacherRepository) GetAllTeachers() []model.Teacher {
	stmt := SELECT(
		Teacher.AllColumns,
	).FROM(
		Teacher,
	)

	var dest []model.Teacher
	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	return dest
}

func (r *teacherRepository) InsertMultipleTeachers(teachers []model.Teacher) {
	insertStmt := Teacher.INSERT(
		Teacher.FirstName,
		Teacher.LastName,
		Teacher.Dob,
		Teacher.Email,
		Teacher.DepartmentID,
	).MODELS(teachers)
	_, err := insertStmt.Exec(r.db)
	util.PanicOnError(err)
}

func (r *teacherRepository) ClearAllTeachers() {
	_, err := r.db.Exec("TRUNCATE TABLE teacher RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating teacher table and reset auto increment")
}
