//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var GradeSetting = newGradeSettingTable("public", "grade_setting", "")

type gradeSettingTable struct {
	postgres.Table

	// Columns
	ID                postgres.ColumnInteger
	AssignmentPercent postgres.ColumnInteger
	ExamPercent       postgres.ColumnInteger
	PassingGrade      postgres.ColumnInteger
	CreatedAt         postgres.ColumnTimestamp
	UpdatedAt         postgres.ColumnTimestamp
	DeletedAt         postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type GradeSettingTable struct {
	gradeSettingTable

	EXCLUDED gradeSettingTable
}

// AS creates new GradeSettingTable with assigned alias
func (a GradeSettingTable) AS(alias string) *GradeSettingTable {
	return newGradeSettingTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new GradeSettingTable with assigned schema name
func (a GradeSettingTable) FromSchema(schemaName string) *GradeSettingTable {
	return newGradeSettingTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new GradeSettingTable with assigned table prefix
func (a GradeSettingTable) WithPrefix(prefix string) *GradeSettingTable {
	return newGradeSettingTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new GradeSettingTable with assigned table suffix
func (a GradeSettingTable) WithSuffix(suffix string) *GradeSettingTable {
	return newGradeSettingTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newGradeSettingTable(schemaName, tableName, alias string) *GradeSettingTable {
	return &GradeSettingTable{
		gradeSettingTable: newGradeSettingTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newGradeSettingTableImpl("", "excluded", ""),
	}
}

func newGradeSettingTableImpl(schemaName, tableName, alias string) gradeSettingTable {
	var (
		IDColumn                = postgres.IntegerColumn("id")
		AssignmentPercentColumn = postgres.IntegerColumn("assignment_percent")
		ExamPercentColumn       = postgres.IntegerColumn("exam_percent")
		PassingGradeColumn      = postgres.IntegerColumn("passing_grade")
		CreatedAtColumn         = postgres.TimestampColumn("created_at")
		UpdatedAtColumn         = postgres.TimestampColumn("updated_at")
		DeletedAtColumn         = postgres.TimestampColumn("deleted_at")
		allColumns              = postgres.ColumnList{IDColumn, AssignmentPercentColumn, ExamPercentColumn, PassingGradeColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
		mutableColumns          = postgres.ColumnList{AssignmentPercentColumn, ExamPercentColumn, PassingGradeColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
	)

	return gradeSettingTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                IDColumn,
		AssignmentPercent: AssignmentPercentColumn,
		ExamPercent:       ExamPercentColumn,
		PassingGrade:      PassingGradeColumn,
		CreatedAt:         CreatedAtColumn,
		UpdatedAt:         UpdatedAtColumn,
		DeletedAt:         DeletedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
