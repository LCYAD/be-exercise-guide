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

func CourseSeeder(db *sql.DB) {
	if repository.CourseExists(db) {
		fmt.Println("Already created Courses.  Skipping....")
	} else {
		var courseModelLinks []model.Course
		departments := repository.GetAllDepartments(db)
		teachers := repository.GetAllTeachers(db)

		// group teacher by department
		teachersByDepartment := make(map[int32][]int32)
		for _, teacher := range teachers {
			deptID := *teacher.DepartmentID
			teachersByDepartment[deptID] = append(teachersByDepartment[deptID], int32(teacher.ID))
		}

		now := time.Now().UTC()
		for _, d := range departments {
			courses := departmentCourses[d.Name]
			for _, c := range courses {
				description := gofakeit.Paragraph(1, 10, 100, " ")
				// might have cases where there is no teacher in the department that triggers course_teacher_id_fkey error
				// TODO: investigate about this issue later, can replicate using a small teacher size
				modelLink := model.Course{
					Name:         c,
					Description:  &description,
					DepartmentID: &d.ID,
					TeacherID:    &teachersByDepartment[d.ID][rand.Intn(len(teachersByDepartment[d.ID]))],
					CreatedAt:    &now,
					UpdatedAt:    &now,
				}
				courseModelLinks = append(courseModelLinks, modelLink)
			}

		}
		repository.InsertMultipleCourses(db, courseModelLinks)
		fmt.Println("Finish seeding Course")
	}
}

var departmentCourses = map[string][]string{
	"Computer Science": {
		"Introduction to Programming",
		"Data Structures and Algorithms",
		"Database Systems",
		"Computer Networks",
		"Artificial Intelligence",
		"Software Engineering",
		"Web Development",
	},
	"Biology": {
		"Cell Biology",
		"Genetics",
		"Ecology",
		"Microbiology",
		"Evolutionary Biology",
		"Molecular Biology",
		"Biochemistry",
	},
	"Chemistry": {
		"Organic Chemistry",
		"Inorganic Chemistry",
		"Physical Chemistry",
		"Analytical Chemistry",
		"Environmental Chemistry",
		"Quantum Chemistry",
	},
	"Physics": {
		"Classical Mechanics",
		"Electromagnetism",
		"Quantum Mechanics",
		"Thermodynamics",
		"Optics",
		"Astrophysics",
		"Nuclear Physics",
	},
	"Mathematics": {
		"Calculus",
		"Linear Algebra",
		"Differential Equations",
		"Number Theory",
		"Topology",
		"Abstract Algebra",
		"Numerical Analysis",
	},
	"Economics": {
		"Microeconomics",
		"Macroeconomics",
		"Econometrics",
		"International Economics",
		"Public Finance",
		"Development Economics",
		"Game Theory",
	},
	"English Literature": {
		"Shakespeare Studies",
		"Modern Poetry",
		"Victorian Literature",
		"American Literature",
		"Literary Theory",
		"Creative Writing",
		"World Literature",
	},
	"History": {
		"Ancient Civilizations",
		"Medieval History",
		"Modern European History",
		"American History",
		"World War II",
		"Cold War Era",
		"Historical Methodology",
	},
	"Psychology": {
		"Cognitive Psychology",
		"Social Psychology",
		"Developmental Psychology",
		"Abnormal Psychology",
		"Neuroscience",
		"Clinical Psychology",
		"Research Methods in Psychology",
	},
	"Political Science": {
		"Introduction to Political Theory",
		"Comparative Politics",
		"International Relations",
		"Public Policy",
		"Constitutional Law",
		"Political Economy",
		"Global Governance",
	},
}
