package seeder

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"

	"github.com/brianvoe/gofakeit/v7"
)

func AssignmentSeeder(db *sql.DB) {
	minAssignmentCount := 3
	courseRepository := repository.NewCourseRepository(db)
	courseIDs := courseRepository.GetCourseIDs()
	// increasing the ratio to approved vs false to 6:1
	gradedOption := []bool{true, true, true, true, true, false}
	randomTitles := []string{
		"Midterm Research Paper",
		"Group Project Presentation",
		"Lab Report Analysis",
		"Literature Review Essay",
		"Case Study Evaluation",
		"Final Exam Preparation",
		"Coding Challenge Implementation",
		"Data Analysis and Visualization",
		"Argumentative Essay",
		"Practical Skills Assessment",
	}

	var assignmentModelLinks []model.Assignment
	for _, courseID := range courseIDs {
		numAssignmentToCreate := rand.Intn(2) + minAssignmentCount
		now := time.Now().UTC()
		for range numAssignmentToCreate {
			description := gofakeit.Sentence(50)
			daysToAdd := rand.Intn(10) + 3
			modelLink := model.Assignment{
				Title:       gofakeit.RandomString(randomTitles),
				Description: &description,
				Type:        int16(rand.Intn(3)), // 0: assignment, 1: quiz, 2: project, or maybe assign it to the randomTitles?
				DueDate:     now.AddDate(0, 0, daysToAdd),
				CourseID:    &courseID,
				Graded:      &gradedOption[rand.Intn(len(gradedOption))],
			}
			assignmentModelLinks = append(assignmentModelLinks, modelLink)
		}
	}
	// TODO clean up repository struct creation (perhaps from main then pass along the seeding flow?)
	assignmentRepo := repository.NewAssignmentRepository(db)
	assignmentRepo.InsertMultipleAssignments(assignmentModelLinks)
	fmt.Println("Finish seeding Assignment")
}
