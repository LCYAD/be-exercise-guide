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

type CourseRepository interface {
	GetCourseIDs() []int32
	CourseExists() bool
	InsertMultipleCourses(departments []model.Course)
	ClearAllCourses()
}

type courseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *courseRepository {
	return &courseRepository{
		db: db,
	}
}

func (r *courseRepository) GetCourseIDs() []int32 {
	stmt := SELECT(
		Course.ID,
	).FROM(
		Course,
	)

	var dest []model.Course

	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	ids := make([]int32, len(dest))
	for i, d := range dest {
		ids[i] = int32(d.ID)
	}

	return ids
}

func (r *courseRepository) CourseExists() bool {
	stmt := SELECT(
		Course.ID,
	).FROM(
		Course,
	).LIMIT(1)

	var dest []model.Course

	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	return len(dest) > 0
}

func (r *courseRepository) InsertMultipleCourses(departments []model.Course) {
	insertStmt := Course.INSERT(
		Course.Name,
		Course.Description,
		Course.DepartmentID,
		Course.TeacherID,
	).MODELS(departments)
	_, err := insertStmt.Exec(r.db)
	util.PanicOnError(err)
}

func (r *courseRepository) ClearAllCourses() {
	_, err := r.db.Exec("TRUNCATE TABLE course RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating course table and reset auto increment")
}
