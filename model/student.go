package model

import "time"

type Student struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	SSID      string    `json:"ssid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteStudentRequest struct {
	ID string `param:"id"`
}

type UpdateStudentRequest struct {
	ID string `param:"id"`
	*Student
}
