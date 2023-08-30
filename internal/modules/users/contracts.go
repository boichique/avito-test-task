package users

type CreateUserRequest struct {
	UserID int `param:"userID"`
}

type DeleteUserRequest struct {
	UserID int `param:"userID"`
}
