package seeder

import (
	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockDepartmentRepository struct {
	mock.Mock
}

func (m *mockDepartmentRepository) GetDepartmentIDs() []int32 {
	args := m.Called()
	return args.Get(0).([]int32)
}

func (m *mockDepartmentRepository) InsertMultipleDepartments(departmentModel []model.Department) {
	m.Called(departmentModel)
}

func TestSeed(t *testing.T) {
	mockRepo := new(mockDepartmentRepository)

	t.Run("Will Seed Department", func(t *testing.T) {
		// TODO: look into time mocking in Go, currently the result time do not point to the same address
		now := time.Now().UTC()
		mockRepo.On("GetDepartmentIDs").Return([]int32{})
		mockRepo.On("InsertMultipleDepartments", mock.Anything).Run(func(args mock.Arguments) {
			departmentModel := args[0].([]model.Department)
			expectedRes := []model.Department{
				{Name: "Computer Science", CreatedAt: &now, UpdatedAt: &now},
				{Name: "Biology", CreatedAt: &now, UpdatedAt: &now},
				{Name: "Chemistry", CreatedAt: &now, UpdatedAt: &now},
				{Name: "Physics", CreatedAt: &now, UpdatedAt: &now},
				{Name: "Mathematics", CreatedAt: &now, UpdatedAt: &now},
				{Name: "Economics", CreatedAt: &now, UpdatedAt: &now},
				{Name: "English Literature", CreatedAt: &now, UpdatedAt: &now},
				{Name: "History", CreatedAt: &now, UpdatedAt: &now},
				{Name: "Psychology", CreatedAt: &now, UpdatedAt: &now},
				{Name: "Political Science", CreatedAt: &now, UpdatedAt: &now},
			}
			if len(departmentModel) != 10 {
				t.Errorf("Expected length of department model is %d, but got %d", 10, len(departmentModel))
			}
			if reflect.DeepEqual(departmentModel, expectedRes) {
				t.Errorf("Input do not match")
			}
		})
		s := NewDepartmentSeeder(mockRepo)
		s.Seed()

		mockRepo.AssertExpectations(t)
	})
	t.Run("Will Skip Seeding", func(t *testing.T) {
		mockRepo := new(mockDepartmentRepository)
		mockRepo.On("GetDepartmentIDs").Return([]int32{1})

		s := NewDepartmentSeeder(mockRepo)
		s.Seed()

		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "InsertMultipleDepartments")
	})
}
