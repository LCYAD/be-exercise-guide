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
	teacherRepository := repository.NewTeacherRepository(db)
	scoreRepository := repository.NewScoreRepository(db)
	teachers := teacherRepository.GetAllTeachers()

	// group teacher by department
	teachersByDepartment := make(map[int32][]int32)
	for _, teacher := range teachers {
		deptID := *teacher.DepartmentID
		teachersByDepartment[deptID] = append(teachersByDepartment[deptID], int32(teacher.ID))
	}

	now := time.Now().UTC()
	var scoreModelLinks []model.Score
	for _, s := range submissions {
		// skipping assignment with submission time over due date
		// currently using UTC time as a cutoff, can review if this is a correct appraoch or not
		if !s.IsAssignment || (s.IsAssignment && s.AssignmentDueDate.AddDate(0, 0, 1).Before(s.SubmittedAt)) {
			modelLink := model.Score{
				SubmissionID: &s.ID,
				TeacherID:    &teachersByDepartment[s.DepartmentID][rand.Intn(len(teachersByDepartment[s.DepartmentID]))],
				Value:        int32(rand.Intn(101)),
				CreatedAt:    &now,
				UpdatedAt:    &now,
			}
			scoreModelLinks = append(scoreModelLinks, modelLink)
		}
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
		scoreRepository.InsertMultipleScores(batch)
	}

	fmt.Println("Finish seeding Score")
}
