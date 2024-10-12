package repository

import (
	"reflect"
	"testing"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetExamIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	repo := NewExamRepository(db)

	mock.ExpectQuery(`SELECT exam\.id AS "exam.id" FROM public\.exam`).
		WillReturnRows(sqlmock.NewRows([]string{"exam.id"}).
			AddRow(1).
			AddRow(2).
			AddRow(3))

	ids := repo.GetExamIDs()

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

func TestGetExamsByCourseID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	repo := NewExamRepository(db)

	courseId := int32(1)

	now := time.Now()

	mockRows := sqlmock.NewRows([]string{
		"exam.id",
		"exam.name",
		"exam.type",
		"exam.started_at",
		"exam.finished_at",
		"exam.course_id",
		"exam.created_at",
		"exam.updated_at",
	}).
		AddRow(1, "ABC", 0, &now, &now, 1, &now, &now)

	mock.ExpectQuery(`SELECT exam\..* FROM public\.exam WHERE exam\.course_id = \$1`).
		WithArgs(courseId).
		WillReturnRows(mockRows)

	exams := repo.GetExamsByCourseID(courseId)

	expectedRes := []model.Exam{
		{ID: 1, Name: "ABC", Type: 0, StartedAt: &now, FinishedAt: &now, CourseID: &courseId, CreatedAt: &now, UpdatedAt: &now},
	}

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(exams) != len(expectedRes) {
		t.Errorf("expected %d exams, got %d", len(exams), len(expectedRes))
	}

	if !reflect.DeepEqual(exams, expectedRes) {
		t.Errorf("output doesn't match")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
