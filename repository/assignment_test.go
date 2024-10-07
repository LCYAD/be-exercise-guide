package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
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

	for i, id := range expectedIDs {
		if ids[i] != id {
			t.Errorf("expected ID %d, got %d", id, ids[i])
		}
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
