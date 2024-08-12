CREATE TABLE department (
    id SERIAL PRIMARY KEY,
    dept_name VARCHAR(255) NOT NULL
);

CREATE TABLE student (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    dob DATE NOT NULL,
    department_id INT REFERENCES department(id) ON DELETE SET NULL
);

CREATE TABLE teacher (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    dob DATE NOT NULL,
    department_id INT REFERENCES department(id) ON DELETE SET NULL
);

CREATE TABLE course (
    id SERIAL PRIMARY KEY,
    course_name VARCHAR(255) NOT NULL,
    department_id INT REFERENCES department(id) ON DELETE CASCADE,
    teacher_id INT REFERENCES teacher(id) ON DELETE SET NULL
);

CREATE TABLE score (
    id SERIAL PRIMARY KEY,
    value INT NOT NULL,
    teacher_id INT REFERENCES teacher(id) ON DELETE SET NULL
);

CREATE TABLE assignment (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    type SMALLINT NOT NULL,
    assigned_at DATE NOT NULL,
    due_date DATE NOT NULL,
    course_id INT REFERENCES course(id) ON DELETE CASCADE
);

CREATE TABLE exam (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type SMALLINT NOT NULL,
    started_at TIMESTAMP NOT NULL,
    finished_at TIMESTAMP NOT NULL,
    course_id INT REFERENCES course(id) ON DELETE CASCADE
);

CREATE TABLE submission (
    id SERIAL PRIMARY KEY,
    submitted_at DATE NOT NULL,
    student_id INT REFERENCES student(id) ON DELETE CASCADE,
    assignment_id INT REFERENCES assignment(id) ON DELETE CASCADE,
    exam_id INT REFERENCES exam(id) ON DELETE CASCADE,
    score_id INT REFERENCES score(id) ON DELETE SET NULL
);


CREATE TABLE enrollment (
    enrollment_id SERIAL PRIMARY KEY,
    student_id INT REFERENCES student(id) ON DELETE CASCADE,
    course_id INT REFERENCES course(id) ON DELETE CASCADE,
    submitted_at DATE NOT NULL,
    approved BOOLEAN NOT NULL
);
