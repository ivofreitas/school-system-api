package mysql

import (
	"context"
	"database/sql"

	"github.com/leantech/school-system-api/api/v1/user"
	"github.com/leantech/school-system-api/model"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {

	insert := `
	INSERT INTO school.users(id, username, password, role_id, created_at, updated_at) 
	VALUES (?, ?, ?, (SELECT id FROM school.roles WHERE position = ?), NOW(), NOW())`

	_, err := r.db.ExecContext(
		ctx,
		insert,
		user.ID,
		user.Username,
		user.Password,
		user.Role)

	return err
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT u.username, u.password, r.position 
		FROM school.users u
		JOIN school.roles r ON u.role_id = r.id 
		WHERE u.username = ?
	`, username)

	user := new(model.User)
	err := row.Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil
}
