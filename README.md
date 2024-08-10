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
    * a student will make a submission for an assignment
    * an assignment will be scored by a teacher
    * a course can have many exams
    * a exam will be scored by a single teacher


## ERP Diagram
```mermaid
erDiagram
    student {
        int id PK
        string first_name
        string last_name
        string email
        int dob 
        int department_id FK
    }
    
    department {
        int id PK
        string dept_name
    }
    
    teacher {
        int id PK
        string first_name
        string last_name
        string email
        int dob 
        int department_id FK
    }
    
    course {
        int id PK
        string course_name
        string description
        int department_id FK
        int teacher_id FK
    }
    
    assignment {
        int id PK
        string title
        string description
        date assigned_at
        date due_date
        int course_id FK
        int score_id FK
    }

    submission {
        int id PK
        date submitted_at
        int student_id FK
        int assignment_id FK
    }

    exam {
        int id PK
        string name
        tinyint type
        date started_at
        date finished_at
        int course_id FK
        int score_id FK
    }

    score {
        int id PK
        int value
        int teacher_id FK
    }
    
    enrollment {
        int enrollment_id PK
        int student_id FK
        int course_id FK
        date submitted_at
        bool approved
    }
    
    student ||--o{ enrollment : "enrolls"
    course ||--o{ enrollment : "contains"
    
    student }o--|| department : "studies in"
    teacher }o--|| department : "works for"
    
    department ||--o{ course : "offers"
    
    course ||--o{ assignment : "contains"
    course ||--o{ exam : "contains"

    student ||--o{ submission : "submits"
    assignment ||--o{ submission : "contains"
    
    submission ||--o| score : "receives"
    exam ||--|| score : "receives"
    
    score }o--|| teacher : "is graded by"
    
    course }o--|| teacher : "is taught by"

```