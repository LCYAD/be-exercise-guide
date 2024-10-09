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

type departmentRepository struct {
	db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) *departmentRepository {
	return &departmentRepository{
		db: db,
	}
}

func (r *departmentRepository) GetAllDepartments() []model.Department {
	stmt := SELECT(
		Department.AllColumns,
	).FROM(
		Department,
	)

	var dest []model.Department
	err := stmt.Query(r.db, &dest)
	util.PanicOnError(err)

	return dest
}

func (r *departmentRepository) GetDepartmentIDs() []int32 {
	var department = r.GetAllDepartments()

	ids := make([]int32, len(department))
	for i, d := range department {
		ids[i] = int32(d.ID)
	}

	return ids
}

func (r *departmentRepository) InsertMultipleDepartments(departments []model.Department) {
	insertStmt := Department.INSERT(Department.Name, Department.CreatedAt, Department.UpdatedAt).MODELS(departments)
	_, err := insertStmt.Exec(r.db)
	util.PanicOnError(err)
}

func (r *departmentRepository) ClearAllDepartments() {
	_, err := r.db.Exec("TRUNCATE TABLE department RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating department table and reset auto increment")
}
