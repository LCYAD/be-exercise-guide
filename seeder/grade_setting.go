package seeder

import (
	"fmt"
	"math/rand"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

type GradeSettingSeeder interface {
	Seed()
	Deseed()
}

type gradeSettingSeeder struct {
	gradeSettingRepo repository.GradeSettingRepository
	courseRepo       repository.CourseRepository
}

func NewGradeSettingSeeder(
	gradeSettingRepo repository.GradeSettingRepository,
	courseRepo repository.CourseRepository,
) *gradeSettingSeeder {
	return &gradeSettingSeeder{
		gradeSettingRepo: gradeSettingRepo,
		courseRepo:       courseRepo,
	}
}

func (s *gradeSettingSeeder) Seed() {
	courseIDs := s.courseRepo.GetCourseIDs()
	var gradeSettingModelLinks []model.GradeSetting

	assignmentPercentRandomChoice := []int32{20, 25, 30, 35, 40, 45}
	passingGradeRandomChoice := []int32{60, 65, 70, 75, 80}

	for _, courseID := range courseIDs {
		assignmentPercent := assignmentPercentRandomChoice[rand.Intn(len(assignmentPercentRandomChoice))]
		modelLink := model.GradeSetting{
			AssignmentPercent: assignmentPercent,
			ExamPercent:       100 - assignmentPercent,
			PassingGrade:      passingGradeRandomChoice[rand.Intn(len(passingGradeRandomChoice))],
			CourseID:          &courseID,
		}
		gradeSettingModelLinks = append(gradeSettingModelLinks, modelLink)
	}
	s.gradeSettingRepo.InsertMultipleGradeSettings(gradeSettingModelLinks)
	fmt.Println("Finish seeding GradeSetting")
}

func (s *gradeSettingSeeder) Deseed() {
	s.gradeSettingRepo.ClearAllGradeSettings()
}
