package repository

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetSubmissionIDsAndDepartmentIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	repo := NewSubmissionRepository(db)

	now := time.Now()

	mockRows1 := sqlmock.NewRows([]string{
		"submission.id",
		"course.department_id",
		"submission.submitted_at",
		"assignment.due_date",
	}).
		AddRow(1, 1, &now, &now)

	mockRows2 := sqlmock.NewRows([]string{
		"submission.id",
		"course.department_id",
		"submission.submitted_at",
	}).
		AddRow(2, 2, &now)

	mock.ExpectQuery(`SELECT submission\..*, assignment\..*, course\..* FROM public\.submission LEFT JOIN public\.assignment ON \(submission\.assignment_id = assignment\.id\) LEFT JOIN public\.course ON \(assignment\.course_id = course\.id\) WHERE submission\.assignment_id IS NOT NULL`).
		WillReturnRows(mockRows1)

	mock.ExpectQuery(`SELECT submission\..*, course\..* FROM public\.submission LEFT JOIN public\.exam ON \(submission\.exam_id = exam\.id\) LEFT JOIN public\.course ON \(exam\.course_id = course\.id\) WHERE submission\.exam_id IS NOT NULL`).
		WillReturnRows(mockRows2)

	res := repo.GetSubmissionIDsAndDepartmentIDs()

	expectedRes := []SubmissionRes{
		{ID: 1, DepartmentID: 1, SubmittedAt: now, AssignmentDueDate: now, IsAssignment: true},
		{ID: 2, DepartmentID: 2, SubmittedAt: now, AssignmentDueDate: time.Time{}, IsAssignment: false},
	}

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(res) != len(expectedRes) {
		t.Errorf("expected %d teachers, got %d", len(res), len(expectedRes))
	}

	if !reflect.DeepEqual(res, expectedRes) {
		t.Errorf("output doesn't match")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
