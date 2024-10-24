package seeder

import (
	"fmt"
	"math/rand"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

type ScoreSeeder interface {
	Seed()
	Deseed()
}

type scoreSeeder struct {
	scoreRepo      repository.ScoreRepository
	teacherRepo    repository.TeacherRepository
	submissionRepo repository.SubmissionRepository
}

func NewScoreSeeder(
	scoreRepo repository.ScoreRepository,
	teacherRepo repository.TeacherRepository,
	submissionRepo repository.SubmissionRepository,
) *scoreSeeder {
	return &scoreSeeder{
		scoreRepo:      scoreRepo,
		teacherRepo:    teacherRepo,
		submissionRepo: submissionRepo,
	}
}

func (s *scoreSeeder) Seed() {
	teachers := s.teacherRepo.GetAllTeachers()
	submissions := s.submissionRepo.GetSubmissionIDsAndDepartmentIDs()

	// group teacher by department
	teachersByDepartment := make(map[int32][]int32)
	for _, teacher := range teachers {
		deptID := *teacher.DepartmentID
		teachersByDepartment[deptID] = append(teachersByDepartment[deptID], int32(teacher.ID))
	}

	var scoreModelLinks []model.Score
	for _, s := range submissions {
		// skipping assignment with submission time over due date
		// currently using UTC time as a cutoff, can review if this is a correct appraoch or not
		if !s.IsAssignment || (s.IsAssignment && s.AssignmentDueDate.AddDate(0, 0, 1).Before(s.SubmittedAt)) {
			modelLink := model.Score{
				SubmissionID: &s.ID,
				TeacherID:    &teachersByDepartment[s.DepartmentID][rand.Intn(len(teachersByDepartment[s.DepartmentID]))],
				Value:        int32(rand.Intn(101)),
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
		s.scoreRepo.InsertMultipleScores(batch)
	}

	fmt.Println("Finish seeding Score")
}

func (s *scoreSeeder) Deseed() {
	s.scoreRepo.ClearAllScores()
}
