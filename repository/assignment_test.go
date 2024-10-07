package repository

import (
	"reflect"
	"testing"

	"be-exerise-go-mod/.gen/be-exercise/public/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAssignmentIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	repo := NewAssignmentRepository(db)

	mock.ExpectQuery(`SELECT assignment\.id AS "assignment.id" FROM public\.assignment`).
		WillReturnRows(sqlmock.NewRows([]string{"assignment.id"}).
			AddRow(1).
			AddRow(2).
			AddRow(3))

	ids := repo.GetAssignmentIDs()

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

func TestGetAssignmentsByCourseID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	repo := NewAssignmentRepository(db)

	mockRows := sqlmock.NewRows([]string{"assignment.id", "assignment.title", "assignment.type"}).
		AddRow(1, "test", 0)

	mock.ExpectQuery(`SELECT assignment\..*FROM public\.assignment WHERE assignment\.course_id = \$1`).
		WithArgs(1).
		WillReturnRows(mockRows)

	assignments := repo.GetAssignmentsByCourseID(1)

	expectedRes := []model.Assignment{
		{ID: 1, Title: "test", Type: 0},
	}

	if len(assignments) != len(expectedRes) {
		t.Errorf("expected %d assignments, got %d", len(assignments), len(expectedRes))
	}

	if !reflect.DeepEqual(assignments, expectedRes) {
		t.Errorf("output doesn't match")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
