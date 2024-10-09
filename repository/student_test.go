package repository

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetStudentIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	repo := NewStudentRepository(db)

	mock.ExpectQuery(`SELECT student\.id AS "student.id" FROM public\.student`).
		WillReturnRows(sqlmock.NewRows([]string{"student.id"}).
			AddRow(1).
			AddRow(2).
			AddRow(3))

	ids := repo.GetStudentIDs()

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
