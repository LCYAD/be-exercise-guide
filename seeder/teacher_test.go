package seeder

import (
	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockTeacherRepository struct {
	mock.Mock
}

func (m *mockTeacherRepository) GetAllTeachers() []model.Teacher {
	args := m.Called()
	return args.Get(0).([]model.Teacher)
}

func (m *mockTeacherRepository) InsertMultipleTeachers(teacherModel []model.Teacher) {
	m.Called(teacherModel)
}

func (m *mockTeacherRepository) ClearAllTeachers() {
	m.Called()
}

// TODO: where to put these?
type mockFaker struct {
	mock.Mock
}

func (m *mockFaker) FirstName() string {
	args := m.Called()
	return args.String(0)
}

func (m *mockFaker) LastName() string {
	args := m.Called()
	return args.String(0)
}

func (m *mockFaker) DateRange(t1 time.Time, t2 time.Time) time.Time {
	args := m.Called(t1, t2)
	return args.Get(0).(time.Time)
}

func (m *mockFaker) Email() string {
	args := m.Called()
	return args.String(0)
}

type mockRandomizer struct {
	mock.Mock
}

func (m *mockRandomizer) Intn(n int) int {
	args := m.Called(n)
	return args.Int(0)
}

func TestTeacherSeed(t *testing.T) {
	mockDepartmentRepo := new(mockDepartmentRepository)
	mockTeacherRepo := new(mockTeacherRepository)
	mf := new(mockFaker)
	mr := new(mockRandomizer)

	now := time.Now().UTC()
	firstName := "ABC"
	lastName := "DEF"
	email := "abc@def.com"
	departmentId := int32(1)
	mockDepartmentRepo.On("GetDepartmentIDs").Return([]int32{departmentId})
	mockTeacherRepo.On("InsertMultipleTeachers", mock.Anything).Run(func(args mock.Arguments) {
		teacherModel := args[0].([]model.Teacher)
		expectedRes := []model.Teacher{
			{FirstName: firstName, LastName: lastName, Email: email, Dob: now, DepartmentID: &departmentId, CreatedAt: &now, UpdatedAt: &now},
		}
		if len(teacherModel) != len(expectedRes) {
			t.Errorf("Expected length of department model is %d, but got %d", len(expectedRes), len(teacherModel))
		}
		if reflect.DeepEqual(teacherModel, expectedRes) {
			t.Errorf("Input do not match")
		}
	})
	mf.On("FirstName").Return(firstName)
	mf.On("LastName").Return(lastName)
	mf.On("DateRange", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		t1 := args[0].(time.Time)
		t2 := args[1].(time.Time)
		t1ExpectedRes := now.AddDate(-70, 0, 0)
		t2ExpectedRes := now.AddDate(-25, 0, 0)
		if t1.Day() != t1ExpectedRes.Day() || t1.Month() != t1ExpectedRes.Month() || t1.Year() != t1ExpectedRes.Year() {
			t.Errorf("args %s do not match, expected %s, got %s", "t1", t1ExpectedRes.Format("2006-01-02"), t1.Format("2006-01-02"))
		}
		if t2.Day() != t2ExpectedRes.Day() || t2.Month() != t2ExpectedRes.Month() || t2.Year() != t2ExpectedRes.Year() {
			t.Errorf("args %s do not match, expected %s, got %s", "t2", t2ExpectedRes.Format("2006-01-02"), t2.Format("2006-01-02"))
		}
	}).Return(now)
	mf.On("Email").Return(email)
	mr.On("Intn").Return(departmentId)
	s := NewTeacherSeeder(mockTeacherRepo, mockDepartmentRepo, mf, mr)
	s.Seed(1)

	mockTeacherRepo.AssertExpectations(t)
	mf.AssertExpectations(t)
}
