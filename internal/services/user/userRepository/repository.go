package userrepository

import (
	"context"
	"database/sql"

	userdomain "github.com/lunatictiol/clean-architecture-in-go/internal/services/user/userDomain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*userdomain.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id = $1", id)

	var u userdomain.User
	if err := row.Scan(&u.Id, &u.Name, &u.Email); err != nil {
		return nil, err
	}
	return &u, nil
}
