package user_errors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("usuário já existe")
	ErrUserNotExists     = errors.New("usuário não existe")
)
