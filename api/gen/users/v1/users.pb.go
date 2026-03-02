package usersv1

type User struct {
	Id               string
	Email            string
	Name             string
	CreatedAtRfc3339 string
	UpdatedAtRfc3339 string
}

type CreateUserRequest struct{ Email, Name string }
type CreateUserResponse struct{ Id string }

type GetUserRequest struct{ Id string }
type GetUserResponse struct{ User *User }

type ListUsersRequest struct{ Limit, Offset int32 }
type ListUsersResponse struct{ Users []*User }

type UpdateUserRequest struct {
	Id    string
	Email *string
	Name  *string
}
type UpdateUserResponse struct{}

type DeleteUserRequest struct{ Id string }
type DeleteUserResponse struct{}
