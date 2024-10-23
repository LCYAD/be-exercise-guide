package seeder

import (
	"fmt"
	"math/rand"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

type examSeeder struct {
	examRepo   repository.ExamRepository
	courseRepo repository.CourseRepository
}

func NewExamSeeder(
	examRepo repository.ExamRepository,
	courseRepo repository.CourseRepository,
) *examSeeder {
	return &examSeeder{
		examRepo:   examRepo,
		courseRepo: courseRepo,
	}
}

func (s *examSeeder) Seed() {
	courseIDs := s.courseRepo.GetCourseIDs()
	examNames := []string{
		"Midterm Exam 1",
		"Midterm Exam 2",
		"Final Exam",
	}

	var examModelLinks []model.Exam
	for _, courseID := range courseIDs {
		now := time.Now().UTC()
		roundedNow := now.Round(time.Hour)
		nextTestDate := roundedNow.AddDate(0, 0, rand.Intn(50)+30)
		for _, name := range examNames {
			hoursToAdd := rand.Intn(2) + 1
			finishedTime := nextTestDate.Add(time.Duration(hoursToAdd) * time.Hour)
			examType := int16(0)
			if name == "Final Exam" {
				examType = 1
			}
			examStartAt := nextTestDate
			modelLink := model.Exam{
				Name:       name,
				Type:       examType, // 0: midterm, 1: final
				StartedAt:  &examStartAt,
				FinishedAt: &finishedTime,
				CourseID:   &courseID,
			}
			examModelLinks = append(examModelLinks, modelLink)
			nextTestDate = nextTestDate.AddDate(0, 0, rand.Intn(50)+30)
		}
	}
	s.examRepo.InsertMultipleExams(examModelLinks)
	fmt.Println("Finish seeding Exam")
}
