package repository

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetCourseIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	repo := NewCourseRepository(db)

	mock.ExpectQuery(`SELECT course\.id AS "course.id" FROM public\.course`).
		WillReturnRows(sqlmock.NewRows([]string{"course.id"}).
			AddRow(1).
			AddRow(2).
			AddRow(3))

	ids := repo.GetCourseIDs()

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

func TestCourseExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()
	repo := NewCourseRepository(db)

	t.Run("course exist", func(t *testing.T) {
		mock.ExpectQuery(`SELECT course\.id AS "course.id" FROM public\.course LIMIT \$1`).
			WillReturnRows(sqlmock.NewRows([]string{"course.id"}).
				AddRow(1))

		res := repo.CourseExists()

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		expected := true

		if expected != res {
			t.Errorf("output doesn't match")
		}

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("course do not exist", func(t *testing.T) {
		mock.ExpectQuery(`SELECT course\.id AS "course.id" FROM public\.course LIMIT \$1`).
			WillReturnRows(sqlmock.NewRows([]string{"course.id"}))

		res := repo.CourseExists()

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		expected := false

		if expected != res {
			t.Errorf("output doesn't match")
		}

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
