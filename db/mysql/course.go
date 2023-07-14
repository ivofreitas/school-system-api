package mysql

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/leantech/school-system-api/api/v1/course"
	"github.com/leantech/school-system-api/api/v1/student"
	"github.com/leantech/school-system-api/model"
)

type courseRepository struct {
	db *sql.DB
	student.Repository
}

func NewCourseRepository(db *sql.DB, studentRepository student.Repository) course.Repository {
	return &courseRepository{db, studentRepository}
}

func (r *courseRepository) CreateCourse(ctx context.Context, course *model.Course) error {

	insert := `
	INSERT INTO school.courses(id, course_name, max_students_number, room_id, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(
		ctx,
		insert,
		course.ID,
		course.CourseName,
		course.MaxStudentsNumber,
		course.RoomID,
		course.CreatedAt,
		course.UpdatedAt)

	return err
}

func (r *courseRepository) GetCourseByRoomID(ctx context.Context, roomID string) (*model.Course, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, course_name, max_students_number, room_id, created_at, updated_at
		FROM school.courses
		WHERE room_id = ?
	`, roomID)

	course := new(model.Course)
	err := row.Scan(&course.ID, &course.CourseName, &course.MaxStudentsNumber, &course.RoomID, &course.CreatedAt, &course.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (r *courseRepository) EnrollStudents(ctx context.Context, enrolledStudents []string, courseID string) error {

	sqlStr := `INSERT INTO school.enrolled_students(id, course_id, student_id) VALUES `
	var values []interface{}

	for _, studentID := range enrolledStudents {
		sqlStr += "(?, ?, ?),"
		values = append(values, uuid.New().String(), courseID, studentID)
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]

	_, err := r.db.ExecContext(ctx, sqlStr, values...)

	return err
}
