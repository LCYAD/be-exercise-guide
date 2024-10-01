package seeder

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

func GradeSettingSeeder(db *sql.DB) {
	courseIDs := repository.GetCourseIDs(db)
	var gradeSettingModelLinks []model.GradeSetting
	now := time.Now().UTC()

	assignmentPercentRandomChoice := []int32{20, 25, 30, 35, 40, 45}
	passingGradeRandomChoice := []int32{60, 65, 70, 75, 80}

	for _, courseID := range courseIDs {
		assignmentPercent := assignmentPercentRandomChoice[rand.Intn(len(assignmentPercentRandomChoice))]
		modelLink := model.GradeSetting{
			AssignmentPercent: assignmentPercent,
			ExamPercent:       100 - assignmentPercent,
			PassingGrade:      passingGradeRandomChoice[rand.Intn(len(passingGradeRandomChoice))],
			CourseID:          &courseID,
			CreatedAt:         &now,
			UpdatedAt:         &now,
		}
		gradeSettingModelLinks = append(gradeSettingModelLinks, modelLink)
	}
	repository.InsertMultipleGradeSettings(db, gradeSettingModelLinks)
	fmt.Println("Finish seeding GradeSetting")
}
