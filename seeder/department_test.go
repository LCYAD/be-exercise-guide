package seeder

import (
	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"reflect"
	"testing"

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

func (m *mockDepartmentRepository) ClearAllDepartments() {
	m.Called()
}

func TestDepartmentSeed(t *testing.T) {
	mockDepartmentRepo := new(mockDepartmentRepository)

	t.Run("Will Seed Department", func(t *testing.T) {
		// TODO: look into time mocking in Go, currently the result time do not point to the same address
		mockDepartmentRepo.On("GetDepartmentIDs").Return([]int32{})
		mockDepartmentRepo.On("InsertMultipleDepartments", mock.Anything).Run(func(args mock.Arguments) {
			departmentModel := args[0].([]model.Department)
			expectedRes := []model.Department{
				{Name: "Computer Science"},
				{Name: "Biology"},
				{Name: "Chemistry"},
				{Name: "Physics"},
				{Name: "Mathematics"},
				{Name: "Economics"},
				{Name: "English Literature"},
				{Name: "History"},
				{Name: "Psychology"},
				{Name: "Political Science"},
			}
			if len(departmentModel) != len(expectedRes) {
				t.Errorf("Expected length of department model is %d, but got %d", len(expectedRes), len(departmentModel))
			}
			if !reflect.DeepEqual(departmentModel, expectedRes) {
				t.Errorf("Input do not match")
			}
		})
		s := NewDepartmentSeeder(mockDepartmentRepo)
		s.Seed()

		mockDepartmentRepo.AssertExpectations(t)
	})
	t.Run("Will Skip Seeding", func(t *testing.T) {
		mockDepartmentRepo := new(mockDepartmentRepository)
		mockDepartmentRepo.On("GetDepartmentIDs").Return([]int32{1})

		s := NewDepartmentSeeder(mockDepartmentRepo)
		s.Seed()

		mockDepartmentRepo.AssertExpectations(t)
		mockDepartmentRepo.AssertNotCalled(t, "InsertMultipleDepartments")
	})
}
