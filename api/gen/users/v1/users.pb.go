// Code generated (scaffold).
// Regenerate with protoc in real projects.

package usersv1

// NOTE: This file intentionally omits full protobuf descriptors.
// It is only here so the scaffold compiles once you `go mod tidy`.

type User struct {
	Id               string `json:"id,omitempty"`
	Email            string `json:"email,omitempty"`
	Name             string `json:"name,omitempty"`
	CreatedAtRfc3339 string `json:"created_at_rfc3339,omitempty"`
	UpdatedAtRfc3339 string `json:"updated_at_rfc3339,omitempty"`
}

type CreateUserRequest struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}
type CreateUserResponse struct{ Id string `json:"id,omitempty"` }

type GetUserRequest struct{ Id string `json:"id,omitempty"` }
type GetUserResponse struct{ User *User `json:"user,omitempty"` }

type ListUsersRequest struct {
	Limit  int32 `json:"limit,omitempty"`
	Offset int32 `json:"offset,omitempty"`
}
type ListUsersResponse struct{ Users []*User `json:"users,omitempty"` }

type UpdateUserRequest struct {
	Id    string  `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}
type UpdateUserResponse struct{}

type DeleteUserRequest struct{ Id string `json:"id,omitempty"` }
type DeleteUserResponse struct{}
