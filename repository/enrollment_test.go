package repository

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestIsStudentEnrolledInCourse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()
	repo := NewEnrollmentRepository(db)

	t.Run("enrollment found for student id and course id", func(t *testing.T) {
		mock.ExpectQuery(`SELECT enrollment\.id AS "enrollment.id" FROM public\.enrollment WHERE \(enrollment\.student_id = \$1::integer\) AND \(enrollment\.course_id = \$2::integer\) LIMIT \$3`).
			WithArgs(1, 1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"enrollment.id"}).
				AddRow(1))

		res := repo.IsStudentEnrolledInCourse(1, 1)

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

	t.Run("enrollment not found for student id and course id", func(t *testing.T) {
		mock.ExpectQuery(`SELECT enrollment\.id AS "enrollment.id" FROM public\.enrollment WHERE \(enrollment\.student_id = \$1::integer\) AND \(enrollment\.course_id = \$2::integer\) LIMIT \$3`).
			WithArgs(1, 1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"enrollment.id"}))

		res := repo.IsStudentEnrolledInCourse(1, 1)

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

func TestGetStudentIDsEnrolledInCourse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()
	repo := NewEnrollmentRepository(db)

	t.Run("return list of student Ids", func(t *testing.T) {
		mock.ExpectQuery(`SELECT enrollment\.student_id AS "enrollment.student_id" FROM public\.enrollment WHERE \(enrollment\.course_id = \$1::integer\) AND \(enrollment\.approved = \$2::boolean\)`).
			WithArgs(1, true).
			WillReturnRows(sqlmock.NewRows([]string{"enrollment.student_id"}).
				AddRow(1))

		res := repo.GetStudentIDsEnrolledInCourse(1)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		expectedIDs := []int32{1}
		if len(res) != len(expectedIDs) {
			t.Errorf("expected %d IDs, got %d", len(expectedIDs), len(res))
		}

		if !reflect.DeepEqual(res, expectedIDs) {
			t.Errorf("output doesn't match")
		}

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("return empty list of student Ids", func(t *testing.T) {
		mock.ExpectQuery(`SELECT enrollment\.student_id AS "enrollment.student_id" FROM public\.enrollment WHERE \(enrollment\.course_id = \$1::integer\) AND \(enrollment\.approved = \$2::boolean\)`).
			WithArgs(1, true).
			WillReturnRows(sqlmock.NewRows([]string{"enrollment.id"}))

		res := repo.GetStudentIDsEnrolledInCourse(1)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		expectedIDs := []int32{}
		if len(res) != len(expectedIDs) {
			t.Errorf("expected %d IDs, got %d", len(expectedIDs), len(res))
		}

		if !reflect.DeepEqual(res, expectedIDs) {
			t.Errorf("output doesn't match")
		}

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
