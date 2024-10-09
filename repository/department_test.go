package repository

import (
	"reflect"
	"testing"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllDepartments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	repo := NewDepartmentRepository(db)

	now := time.Now()

	mockRows := sqlmock.NewRows([]string{
		"department.id",
		"department.name",
		"department.created_at",
		"department.updated_at",
	}).
		AddRow(1, "Computer Science", &now, &now)

	mock.ExpectQuery(`SELECT department\..* FROM public\.department`).
		WillReturnRows(mockRows)

	departments := repo.GetAllDepartments()

	expectedRes := []model.Department{
		{ID: 1, Name: "Computer Science", CreatedAt: &now, UpdatedAt: &now},
	}

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(departments) != len(expectedRes) {
		t.Errorf("expected %d departments, got %d", len(departments), len(expectedRes))
	}

	if !reflect.DeepEqual(departments, expectedRes) {
		t.Errorf("output doesn't match")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetDepartmentIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	repo := NewDepartmentRepository(db)

	mock.ExpectQuery(`SELECT department\..* FROM public\.department`).
		WillReturnRows(sqlmock.NewRows([]string{"department.id"}).
			AddRow(1).
			AddRow(2).
			AddRow(3))

	ids := repo.GetDepartmentIDs()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	expectedIDs := []int32{1, 2, 3}
	if len(ids) != len(expectedIDs) {
		t.Errorf("expected %d IDs, got %d", len(expectedIDs), len(ids))
	}

	if !reflect.DeepEqual(ids, expectedIDs) {
		t.Errorf("output doesn't match")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
