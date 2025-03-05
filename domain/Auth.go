package domain

import "context"

type AuthService interface {
	Login(ctx context.Context)
}
