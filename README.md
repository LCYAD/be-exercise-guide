# be-exercise-guide
A guide for creating BE exercise using the standard use case

## Introduction
* In order to get familiar with writing different programming language and its respective framework, a common case study is being created here
* Common tools will be provided in this repo
    * migration script for DB table generation
        * currently only doing it for Postgres
    * seeder for populating DB with data
* The BE server only needs to connect to the created DB then implement the business logic accordingly to its own design

## Scenario
* A scenario that is commonly used that involves an university where there are students, teachers, courses, exams etc.
* Rules:
    * a teacher works for a department
    * a student studies in a department
    * a department have many courses
    * a student can enroll in many courses from different departments
    * only one teacher can teach in a course
    * a course can contain many students
    * a course can have many assignments
    * a course can have many exams
    * a student will make a submission for an assignment or an exam
    * a submission will be scored by a teacher
    * a grade will be determined for each student that enrolled in a course
        * it will be determined by the assignment's score and also exam score, and the percentage for each partis controlled by the `grade_setting` table
            * the total of the assignment and exam percentage should have a total of 100
        * to keep it simple, there will only be value of the grade, which is determined by the assignment and exam score and if the student has passed or not will be determined by that value and the pass grade required by that course

### Note
* this is simpled version of the university course setup to keep backend exercise business logic thin
* things completely ignored in this design
    * no years of joining for students
        * so cannot determine which year they are in
    * no graduation related information and no date of acceptance
    * no university application related information
    * no student degree information
        * so do not know if a student is here as  PhD, MPhil, Master, Bachelor and the time limit before they can complete their degree
        * no limit on what course the student can join
            * for example, there is a list of courses that could only be joined with elective groups for certain degree
    * courses do not have levels (100, 200, 300, 400)
        * not to create too many courses and prevent the need to setup prerequisite relationship between courses for enrollment
    * courses do not have credit information
        * mainly for calculating if a student can graduate or not
* might consider down the road to create a separate exercise with more complex logic

## ER Diagram
```mermaid
erDiagram
    student {
        int id PK
        string first_name
        string last_name
        string email
        date dob 
        int department_id FK
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }
    
    department {
        int id PK
        string name
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }
    
    teacher {
        int id PK
        string first_name
        string last_name
        string email
        date dob 
        int department_id FK
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }
    
    course {
        int id PK
        string name
        text description
        int department_id FK
        int teacher_id FK
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    enrollment {
        int id PK
        int student_id FK
        int course_id FK
        bool approved
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }
    
    assignment {
        int id PK
        string title
        text description
        tinyint type
        date due_date
        int course_id FK
        bool graded
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    exam {
        int id PK
        string name
        tinyint type
        timestamp started_at
        timestamp finished_at
        int course_id FK
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    submission {
        int id PK
        int student_id FK
        int assignment_id FK
        int exam_id FK
        timestamp submitted_at
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    score {
        int id PK
        int value
        int teacher_id FK
        int submission_id FK
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    grade_setting {
        int id PK
        int assignment_percent "between 0 and 100"
        int exam_percent "between 0 and 100"
        int passing_grade "between 0 and 100"
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    grade {
        int id PK
        int enrollment_id FK
        int value "between 0 and 100"
        bool passed
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }
    
    student ||--o{ enrollment : "enrolls"
    course ||--o{ enrollment : "contains"
    
    student }o--o| department : "studies in"
    teacher }o--o| department : "works for"
    
    department |o--o{ course : "offers"
    
    course ||--o{ assignment : "contains"
    course ||--o{ exam : "contains"

    student ||--o{ submission : "submits"
    assignment ||--o{ submission : "contains"
    exam ||--o{ submission : "contains"
    
    submission ||--o| score : "receives"
    
    score }o--o| teacher : "is graded by"
    
    course }o--o| teacher : "is taught by"

    course ||--|| grade_setting : "contains"
    enrollment ||--o| grade : "contains"
```

## How to use
### Prerequisite
* install [go-migrate CLI](https://github.com/golang-migrate/migrate) (use install via brew if MacOS)
* have Docker installed so that we can use `docker-compose`

## Start up Postgres Server and create tables
* run `docker-compose up -d` to start up the postgres server (use `docker ps` to check)
* run the following command to create the tables
```bash
migrate -path ./db/postgres_migrations/ -database "postgres://<username>:<password>@<domain>:<port>/<DB_name>>?sslmode=disable" up
```

## Clean up Tables and shut down Postgres server
* run the following command to remove the tables
```bash
migrate -path ./db/postgres_migrations/ -database "postgres://<username>:<password>@<domain>:<port>/<DB_name>>?sslmode=disable" down
```
* run `docker-compose down` to shut down the postgres server (use `docker ps` to check)