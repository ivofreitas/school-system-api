package model

import "time"

type Course struct {
	ID                string    `json:"id"`
	RoomID            string    `json:"room_id"`
	MaxStudentsNumber int       `json:"max_students_number"`
	EnrolledStudents  []string  `json:"enrolled_students"`
	CourseName        string    `json:"course_name"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type EnrollStudentRequest struct {
	Students []string `json:"students"`
	RoomID   string   `json:"room_id"`
}
