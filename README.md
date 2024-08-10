# be-exercise-guide
A guide for creating BE exercise using the standard use case

## Introduction
* In order to get familiar with writing different programming language and its respective framework, a common case study is being created here

## Scenario
* A simple 


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