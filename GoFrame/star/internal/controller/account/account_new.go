package account

import "star/internal/logic/users"

type ControllerV1 struct {
	users *users.Users
}

func NewV1() *ControllerV1 {
	return &ControllerV1{
		users: users.New(),
	}
}
