package accountsv1

type Account struct {
	Id               string
	UserId           string
	Label            string
	CreatedAtRfc3339 string
	UpdatedAtRfc3339 string
}

type CreateAccountRequest struct{ UserId, Label string }
type CreateAccountResponse struct{ Id string }

type GetAccountRequest struct{ Id string }
type GetAccountResponse struct{ Account *Account }

type ListAccountsByUserRequest struct{ UserId string; Limit, Offset int32 }
type ListAccountsByUserResponse struct{ Accounts []*Account }

type UpdateAccountRequest struct{ Id string; Label *string }
type UpdateAccountResponse struct{}

type DeleteAccountRequest struct{ Id string }
type DeleteAccountResponse struct{}
