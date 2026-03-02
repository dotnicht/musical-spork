package accountsv1

type Account struct {
	Id               string `json:"id,omitempty"`
	UserId           string `json:"user_id,omitempty"`
	Label            string `json:"label,omitempty"`
	CreatedAtRfc3339 string `json:"created_at_rfc3339,omitempty"`
	UpdatedAtRfc3339 string `json:"updated_at_rfc3339,omitempty"`
}

type CreateAccountRequest struct{ UserId, Label string }
type CreateAccountResponse struct{ Id string }

type GetAccountRequest struct{ Id string }
type GetAccountResponse struct{ Account *Account }

type ListAccountsByUserRequest struct {
	UserId  string
	Limit   int32
	Offset  int32
}
type ListAccountsByUserResponse struct{ Accounts []*Account }

type UpdateAccountRequest struct {
	Id    string
	Label *string
}
type UpdateAccountResponse struct{}

type DeleteAccountRequest struct{ Id string }
type DeleteAccountResponse struct{}
