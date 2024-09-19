CREATE TABLE department (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT department_unique_name UNIQUE(name)
);

CREATE TABLE student (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    dob DATE NOT NULL,
    department_id INT REFERENCES department(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE teacher (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    dob DATE NOT NULL,
    department_id INT REFERENCES department(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE course (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    department_id INT REFERENCES department(id) ON DELETE SET NULL,
    teacher_id INT REFERENCES teacher(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT course_unique_name UNIQUE(name)
);

CREATE TABLE enrollment (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES student(id) ON DELETE CASCADE,
    course_id INT REFERENCES course(id) ON DELETE CASCADE,
    approved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE assignment (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NULL,
    type SMALLINT NOT NULL,
    due_date DATE NOT NULL,
    graded BOOLEAN DEFAULT TRUE,
    course_id INT REFERENCES course(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE exam (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type SMALLINT NOT NULL,
    started_at TIMESTAMP,
    finished_at TIMESTAMP,
    course_id INT REFERENCES course(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE submission (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES student(id) ON DELETE CASCADE,
    assignment_id INT REFERENCES assignment(id) ON DELETE CASCADE,
    exam_id INT REFERENCES exam(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT check_only_one_foreign_key_can_exist
    CHECK (
        (assignment_id IS NOT NULL AND exam_id IS NULL) OR 
        (assignment_id IS NULL AND exam_id IS NOT NULL)
    )
);

CREATE TABLE score (
    id SERIAL PRIMARY KEY,
    value INT NOT NULL,
    teacher_id INT REFERENCES teacher(id) ON DELETE SET NULL,
    submission_id INT REFERENCES submission(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE grade_setting (
    id SERIAL PRIMARY KEY,
    assignment_percent INT NOT NULL CHECK (assignment_percent BETWEEN 0 AND 100),
    exam_percent INT NOT NULL CHECK (exam_percent BETWEEN 0 AND 100),
    passing_grade INT NOT NULL CHECK (passing_grade BETWEEN 0 AND 100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT total_percent_check CHECK (assignment_percent + exam_percent = 100)
);

CREATE TABLE grade (
    id SERIAL PRIMARY KEY,
    enrollment_id INT REFERENCES enrollment(id) ON DELETE CASCADE,
    value INT NOT NULL CHECK (value BETWEEN 0 AND 100),
    passed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);