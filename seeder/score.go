package seeder

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

func ScoreSeeder(db *sql.DB) {
	submissions := repository.GetSubmissionIDsAndDepartmentIDs(db)
	teachers := repository.GetAllTeachers(db)

	// group teacher by department
	teachersByDepartment := make(map[int32][]int32)
	for _, teacher := range teachers {
		deptID := *teacher.DepartmentID
		teachersByDepartment[deptID] = append(teachersByDepartment[deptID], int32(teacher.ID))
	}

	now := time.Now()
	var scoreModelLinks []model.Score
	for _, submission := range submissions {
		modelLink := model.Score{
			SubmissionID: &submission.ID,
			TeacherID:    &teachersByDepartment[submission.DepartmentID][rand.Intn(len(teachersByDepartment[submission.DepartmentID]))],
			Value:        int32(rand.Intn(101)),
			CreatedAt:    &now,
			UpdatedAt:    &now,
		}
		scoreModelLinks = append(scoreModelLinks, modelLink)
	}
	// Define the batch size
	batchSize := 5000

	// Process submissions in batches
	for i := 0; i < len(scoreModelLinks); i += batchSize {
		end := i + batchSize
		if end > len(scoreModelLinks) {
			end = len(scoreModelLinks)
		}
		batch := scoreModelLinks[i:end]
		repository.InsertMultipleScores(db, batch)
	}

	fmt.Println("Finish seeding Score")
}
