package seeder

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

func ExamSeeder(db *sql.DB) {
	courseIDs := repository.GetCourseIDs(db)
	examNames := []string{
		"Midterm Exam 1",
		"Midterm Exam 2",
		"Final Exam",
	}

	var examModelLinks []model.Exam
	for _, courseID := range courseIDs {
		now := time.Now()
		nextTestDate := time.Now().AddDate(0, 0, rand.Intn(50)+30)
		for _, name := range examNames {
			hoursToAdd := rand.Intn(2) + 1
			finishedTime := nextTestDate.Add(time.Duration(hoursToAdd) * time.Hour)
			examType := int16(0)
			if name == "Final Exam" {
				examType = 1
			}
			modelLink := model.Exam{
				Name:       name,
				Type:       examType, // 0: midterm, 1: final
				StartedAt:  &nextTestDate,
				FinishedAt: &finishedTime,
				CourseID:   &courseID,
				CreatedAt:  &now,
				UpdatedAt:  &now,
			}
			examModelLinks = append(examModelLinks, modelLink)
			nextTestDate = nextTestDate.AddDate(0, 0, rand.Intn(50)+30)
		}
	}
	repository.InsertMultipleExams(db, examModelLinks)
	fmt.Println("Finish seeding Exam")
}
