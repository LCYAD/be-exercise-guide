package seeder

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

func SubmissionSeeder(db *sql.DB) {
	chanceOfSubmission := []bool{true, true, true, true, true, true, true, true, true, false}
	courseRepository := repository.NewCourseRepository(db)
	examRepository := repository.NewExamRepository(db)
	submissionRepository := repository.NewSubmissionRepository(db)

	courseIDs := courseRepository.GetCourseIDs()
	// Create repo struct here, u may refactor it later
	assignmentRepo := repository.NewAssignmentRepository(db)
	enrollmentRepository := repository.NewEnrollmentRepository(db)
	var submissionModelLinks []model.Submission
	now := time.Now().UTC()

	for _, courseId := range courseIDs {
		studentIDs := enrollmentRepository.GetStudentIDsEnrolledInCourse(courseId)
		assignments := assignmentRepo.GetAssignmentsByCourseID(courseId)
		exams := examRepository.GetExamsByCourseID(courseId)
		for _, assignment := range assignments {
			for _, studentId := range studentIDs {
				willSubmitAssignment := chanceOfSubmission[rand.Intn(len(chanceOfSubmission))]
				if willSubmitAssignment {
					submissionTime := assignment.CreatedAt.AddDate(0, 0, rand.Intn(15)).Add(time.Duration(rand.Intn(50)) * time.Hour)
					modelLink := model.Submission{
						StudentID:    &studentId,
						AssignmentID: &assignment.ID,
						SubmittedAt:  submissionTime,
						CreatedAt:    &now,
						UpdatedAt:    &now,
					}
					submissionModelLinks = append(submissionModelLinks, modelLink)
				}
			}
		}
		for _, exam := range exams {
			for _, studentId := range studentIDs {
				willSubmitAssignment := chanceOfSubmission[rand.Intn(len(chanceOfSubmission))]
				if willSubmitAssignment {
					modelLink := model.Submission{
						StudentID: &studentId,
						ExamID:    &exam.ID,
						// assumption is that most people at the exam hall will submit at the end of the exam
						SubmittedAt: *exam.FinishedAt,
						CreatedAt:   &now,
						UpdatedAt:   &now,
					}
					submissionModelLinks = append(submissionModelLinks, modelLink)
				}
			}
		}
	}
	// Define the batch size
	batchSize := 3000

	// Process submissions in batches
	for i := 0; i < len(submissionModelLinks); i += batchSize {
		end := i + batchSize
		if end > len(submissionModelLinks) {
			end = len(submissionModelLinks)
		}
		batch := submissionModelLinks[i:end]
		submissionRepository.InsertMultipleSubmissions(batch)
	}

	fmt.Println("Finish seeding Submission")
}
