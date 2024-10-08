package repository

import (
	"reflect"
	"testing"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllTeachers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	repo := NewTeacherRepository(db)

	dob, _ := time.Parse("2006-01-02", "1990-01-01")

	now := time.Now()

	mockRows := sqlmock.NewRows([]string{"teacher.id",
		"teacher.first_name",
		"teacher.last_name",
		"teacher.email",
		"teacher.dob",
		"teacher.department_id",
		"teacher.created_at",
		"teacher.updated_at",
	}).
		AddRow(1, "ABC", "DEF", "abc@gmail.com", dob, 1, &now, &now)

	mock.ExpectQuery(`SELECT teacher\..* FROM public\.teacher`).
		WillReturnRows(mockRows)

	teachers := repo.GetAllTeachers()

	departmentID := int32(1)
	expectedRes := []model.Teacher{
		{ID: 1, FirstName: "ABC", LastName: "DEF", Email: "abc@gmail.com", Dob: dob, DepartmentID: &departmentID, CreatedAt: &now, UpdatedAt: &now},
	}

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(teachers) != len(expectedRes) {
		t.Errorf("expected %d teachers, got %d", len(teachers), len(expectedRes))
	}

	if !reflect.DeepEqual(teachers, expectedRes) {
		t.Errorf("output doesn't match")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
