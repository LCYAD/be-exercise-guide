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

func GetAllDepartments(db *sql.DB) []model.Department {
	stmt := SELECT(
		Department.AllColumns,
	).FROM(
		Department,
	)

	var dest []model.Department
	err := stmt.Query(db, &dest)
	util.PanicOnError(err)

	return dest
}

func GetDepartmentIDs(db *sql.DB) []int32 {
	var department = GetAllDepartments(db)

	ids := make([]int32, len(department))
	for i, d := range department {
		ids[i] = int32(d.ID)
	}

	return ids
}

func InsertMultipleDepartments(db *sql.DB, departments []model.Department) {
	insertStmt := Department.INSERT(Department.Name, Department.CreatedAt, Department.UpdatedAt).MODELS(departments)
	_, err := insertStmt.Exec(db)
	util.PanicOnError(err)
}

func ClearAllDepartments(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE department RESTART IDENTITY CASCADE")
	util.PanicOnError(err)
	fmt.Println("Complete truncating department table and reset auto increment")
}
