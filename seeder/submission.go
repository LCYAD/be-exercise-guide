package seeder

import (
	"fmt"
	"math/rand"
	"time"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

type submissionSeeder struct {
	submissionRepo repository.SubmissionRepository
	courseRepo     repository.CourseRepository
	enrollmentRepo repository.EnrollmentRepository
	assignmentRepo repository.AssignmentRepository
	examRepo       repository.ExamRepository
}

func NewSubmissionSeeder(
	submissionRepo repository.SubmissionRepository,
	courseRepo repository.CourseRepository,
	enrollmentRepo repository.EnrollmentRepository,
	assignmentRepo repository.AssignmentRepository,
	examRepo repository.ExamRepository,
) *submissionSeeder {
	return &submissionSeeder{
		submissionRepo: submissionRepo,
		courseRepo:     courseRepo,
		enrollmentRepo: enrollmentRepo,
		assignmentRepo: assignmentRepo,
		examRepo:       examRepo,
	}
}

func (s *submissionSeeder) Seed() {
	chanceOfSubmission := []bool{true, true, true, true, true, true, true, true, true, false}

	courseIDs := s.courseRepo.GetCourseIDs()
	// Create repo struct here, u may refactor it later
	var submissionModelLinks []model.Submission

	for _, courseId := range courseIDs {
		studentIDs := s.enrollmentRepo.GetStudentIDsEnrolledInCourse(courseId)
		assignments := s.assignmentRepo.GetAssignmentsByCourseID(courseId)
		exams := s.examRepo.GetExamsByCourseID(courseId)
		for _, assignment := range assignments {
			for _, studentId := range studentIDs {
				willSubmitAssignment := chanceOfSubmission[rand.Intn(len(chanceOfSubmission))]
				if willSubmitAssignment {
					submissionTime := assignment.CreatedAt.AddDate(0, 0, rand.Intn(15)).Add(time.Duration(rand.Intn(50)) * time.Hour)
					modelLink := model.Submission{
						StudentID:    &studentId,
						AssignmentID: &assignment.ID,
						SubmittedAt:  submissionTime,
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
		s.submissionRepo.InsertMultipleSubmissions(batch)
	}

	fmt.Println("Finish seeding Submission")
}
