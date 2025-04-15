package userdomain

import "context"

type Repository interface {
	FindByID(ctx context.Context, id int) (*User, error)
}
