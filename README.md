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