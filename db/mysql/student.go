package mysql

import (
	"context"
	"database/sql"
	"strings"

	"github.com/leantech/school-system-api/api/v1/student"
	"github.com/leantech/school-system-api/model"
)

type studentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) student.Repository {
	return &studentRepository{db}
}

func (r *studentRepository) CreateStudent(ctx context.Context, student *model.Student) error {

	insert := `
	INSERT INTO school.students(id, first_name, last_name, ssid, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(
		ctx,
		insert,
		student.ID,
		student.FirstName,
		student.LastName,
		student.SSID,
		student.CreatedAt,
		student.UpdatedAt)

	return err
}

func (r *studentRepository) UpdateStudent(ctx context.Context, id string, student *model.Student) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE school.students
		SET first_name = ?, last_name = ?, ssid = ?, updated_at = NOW()
		WHERE id = ?
	`, student.FirstName, student.LastName, student.SSID, id)

	return err
}

func (r *studentRepository) DeleteStudent(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM school.students
		WHERE id = ?
	`, id)

	return err
}

func (r *studentRepository) GetStudentByID(ctx context.Context, id string) (*model.Student, error) {
	row := r.db.QueryRowContext(ctx, `
	SELECT id, first_name, last_name, ssid, created_at, updated_at
	FROM school.students
	WHERE id = ?`, id)

	student := new(model.Student)
	err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.SSID, &student.CreatedAt, &student.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (r *studentRepository) GetStudentsByIDs(ctx context.Context, ids []string) ([]*model.Student, error) {

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := r.db.QueryContext(ctx, `
	SELECT id, first_name, last_name, ssid, created_at, updated_at
	FROM school.students
	WHERE id IN (?`+strings.Repeat(",?", len(ids)-1)+`)`, args...)
	if err != nil {
		return nil, err
	}

	students := make([]*model.Student, 0)
	for rows.Next() {
		student := new(model.Student)
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.SSID, &student.CreatedAt, &student.UpdatedAt)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	if len(students) == 0 {
		return nil, sql.ErrNoRows
	}

	return students, nil
}
