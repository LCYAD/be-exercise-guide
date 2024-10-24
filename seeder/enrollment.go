package seeder

import (
	"fmt"
	"math/rand"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

type EnrollmentSeeder interface {
	Seed()
	Deseed()
}

type enrollmentSeeder struct {
	enrollmentRepo repository.EnrollmentRepository
	studentRepo    repository.StudentRepository
	courseRepo     repository.CourseRepository
}

func NewEnrollmentRepository(
	enrollmentRepo repository.EnrollmentRepository,
	studentRepo repository.StudentRepository,
	courseRepo repository.CourseRepository,
) *enrollmentSeeder {
	return &enrollmentSeeder{
		enrollmentRepo: enrollmentRepo,
		studentRepo:    studentRepo,
		courseRepo:     courseRepo,
	}
}
func (s *enrollmentSeeder) Seed() {
	minCourseEnroll := 3
	studentIDs := s.studentRepo.GetStudentIDs()
	courseIDs := s.courseRepo.GetCourseIDs()
	// increasing the ratio to approved vs false to 4:1
	approvedOption := []bool{true, true, true, false}

	var enrollmentModelLinks []model.Enrollment
	for _, studentID := range studentIDs {
		coursesEnroll := rand.Intn(5) + minCourseEnroll
		pickedCourseIDs := pickRandomIDs(courseIDs, coursesEnroll)
		for _, cIDs := range pickedCourseIDs {
			if !s.enrollmentRepo.IsStudentEnrolledInCourse(studentID, cIDs) {
				modelLink := model.Enrollment{
					StudentID: &studentID,
					CourseID:  &cIDs,
					Approved:  &approvedOption[rand.Intn(len(approvedOption))],
				}
				enrollmentModelLinks = append(enrollmentModelLinks, modelLink)
			}
		}
	}
	s.enrollmentRepo.InsertMultipleEnrollments(enrollmentModelLinks)
	fmt.Println("Finish seeding Enrollment")
}

func pickRandomIDs(arr []int32, count int) []int32 {
	// Create a copy of the original array to avoid modifying it
	temp := make([]int32, len(arr))
	copy(temp, arr)

	// Shuffle the temporary array
	rand.Shuffle(len(temp), func(i, j int) {
		temp[i], temp[j] = temp[j], temp[i]
	})

	// Return the first 'count' elements
	if count > len(temp) {
		count = len(temp)
	}
	return temp[:count]
}

func (s *enrollmentSeeder) Deseed() {
	s.enrollmentRepo.ClearAllEnrollments()
}
