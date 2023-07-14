CREATE DATABASE IF NOT EXISTS school;


CREATE TABLE school.students (
    id          VARCHAR(40) NOT NULL,
    first_name  VARCHAR(80) NULL,
    last_name   VARCHAR(80) NULL,
    ssid        VARCHAR(40) UNIQUE NOT NULL,
    created_at  TIMESTAMP   NULL,
    updated_at  TIMESTAMP   NULL,
    PRIMARY KEY (id)
);

CREATE TABLE school.courses (
     id                  VARCHAR(40) NOT NULL,
     course_name         VARCHAR(80) NULL,
     max_students_number VARCHAR(40) NULL,
     room_id             VARCHAR(40) UNIQUE NOT NULL,
     created_at          TIMESTAMP   NULL,
     updated_at          TIMESTAMP   NULL,
     PRIMARY KEY (id)
);

CREATE TABLE school.enrolled_students (
      id         VARCHAR(40) NOT NULL,
      course_id  VARCHAR(40) NOT NULL,
      student_id VARCHAR(40) NOT NULL,
      PRIMARY KEY (id),
      FOREIGN KEY (course_id)   REFERENCES school.courses(id),
      FOREIGN KEY (student_id)  REFERENCES school.students(id)
);

CREATE TABLE school.roles (
    id          VARCHAR(40) NOT NULL,
    position    VARCHAR(80) UNIQUE NULL,
    created_at  TIMESTAMP   NULL,
    updated_at  TIMESTAMP   NULL,
    PRIMARY KEY (id)
);

CREATE TABLE school.users (
    id          VARCHAR(40)  NOT NULL,
    username    VARCHAR(80)  UNIQUE NOT NULL,
    password    VARCHAR(255) NOT NULL,
    role_id     VARCHAR(40)  NULL,
    created_at  TIMESTAMP    NULL,
    updated_at  TIMESTAMP    NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (role_id)   REFERENCES school.roles(id)
);

INSERT INTO school.roles
VALUES ('9c06de8d-865a-4ed4-88d0-a445f2bd4f11', 'student', NOW(), NOW()),
       ('59a2cc28-a219-47c2-945a-e520827bbfcd', 'teacher', NOW(), NOW()),
       ('ca99ac9f-f26d-4aeb-8966-234c277ff7b4', 'admin', NOW(), NOW());

INSERT INTO school.users
(id, username, password, role_id, created_at, updated_at)
VALUES('d70223f3-5738-4822-b412-b8633d187ff5', 'john.tech', '$2a$10$fNhHGDgvvu0UVjPjCEHPveIdn5Q4MdaUWZ0esySSRl9JjqaHhhSMy', 'ca99ac9f-f26d-4aeb-8966-234c277ff7b4', NOW(), NOW());

