package seeder

import (
	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockStudentRepository struct {
	mock.Mock
}

func (m *mockStudentRepository) GetStudentIDs() []int32 {
	args := m.Called()
	return args.Get(0).([]int32)
}

func (m *mockStudentRepository) InsertMultipleStudents(studentModel []model.Student) {
	m.Called(studentModel)
}

func (m *mockStudentRepository) ClearAllStudents() {
	m.Called()
}

func TestStudentSeed(t *testing.T) {
	mockDepartmentRepo := new(mockDepartmentRepository)
	mockStudentRepo := new(mockStudentRepository)
	mf := new(mockFaker)
	mr := new(mockRandomizer)

	now := time.Now().UTC()
	firstName := "ABC"
	lastName := "DEF"
	email := "abc@def.com"
	departmentId := int32(1)
	expectedRes := []model.Student{
		{FirstName: firstName, LastName: lastName, Email: email, Dob: now, DepartmentID: &departmentId},
	}
	mockDepartmentRepo.On("GetDepartmentIDs").Return([]int32{departmentId})
	mockStudentRepo.On("InsertMultipleStudents", mock.Anything).Run(func(args mock.Arguments) {
		studentModel := args[0].([]model.Student)
		if len(studentModel) != len(expectedRes) {
			t.Errorf("Expected length of department model is %d, but got %d", len(expectedRes), len(studentModel))
		}
		if !reflect.DeepEqual(studentModel, expectedRes) {
			t.Errorf("Input do not match")
		}
	})
	mf.On("FirstName").Return(firstName)
	mf.On("LastName").Return(lastName)
	mf.On("DateRange", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		t1 := args[0].(time.Time)
		t2 := args[1].(time.Time)
		t1ExpectedRes := now.AddDate(-50, 0, 0)
		t2ExpectedRes := now.AddDate(-20, 0, 0)
		if t1.Day() != t1ExpectedRes.Day() || t1.Month() != t1ExpectedRes.Month() || t1.Year() != t1ExpectedRes.Year() {
			t.Errorf("args %s do not match, expected %s, got %s", "t1", t1ExpectedRes.Format("2006-01-02"), t1.Format("2006-01-02"))
		}
		if t2.Day() != t2ExpectedRes.Day() || t2.Month() != t2ExpectedRes.Month() || t2.Year() != t2ExpectedRes.Year() {
			t.Errorf("args %s do not match, expected %s, got %s", "t2", t2ExpectedRes.Format("2006-01-02"), t2.Format("2006-01-02"))
		}
	}).Return(now)
	mf.On("Email").Return(email)
	mr.On("Intn", mock.Anything).Run(func(args mock.Arguments) {
		n := args[0]
		if n != len(expectedRes) {
			t.Errorf("incorrect length passed, expected %d, got %d", len(expectedRes), n)
		}
	}).Return(0)
	s := NewStudentSeeder(mockStudentRepo, mockDepartmentRepo, mf, mr)
	s.Seed(1)

	mockStudentRepo.AssertExpectations(t)
	mf.AssertExpectations(t)
}
