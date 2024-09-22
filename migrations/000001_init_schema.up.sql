
CREATE TABLE student (
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE course (
    id SERIAL PRIMARY KEY,
    description TEXT
);

CREATE TABLE enrollment (
    id SERIAL PRIMARY KEY,
    student_id INT,
    course_id INT,

    FOREIGN KEY (student_id) REFERENCES student(id),
    FOREIGN KEY (course_id) REFERENCES course(id)
);
