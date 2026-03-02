package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrAccountNotFound  = errors.New("account not found")
	ErrInvalidAccountLabel = errors.New("invalid account label")
	ErrUserDoesNotExist = errors.New("user does not exist")
)

type AccountID string

func NewAccountID() AccountID { return AccountID(uuid.NewString()) }

type Account struct {
	id        AccountID
	userID    string
	label     string
	createdAt time.Time
	updatedAt time.Time
}

func CreateNewAccount(userID, label string) (*Account, error) {
	userID = strings.TrimSpace(userID)
	label = strings.TrimSpace(label)
	if userID == "" {
		return nil, errors.New("user_id is required")
	}
	if label == "" || len(label) > 120 {
		return nil, ErrInvalidAccountLabel
	}
	now := time.Now().UTC()
	return &Account{
		id:        NewAccountID(),
		userID:    userID,
		label:     label,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func RehydrateAccount(id AccountID, userID, label string, createdAt, updatedAt time.Time) *Account {
	return &Account{
		id:        id,
		userID:    userID,
		label:     label,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (a *Account) ID() AccountID           { return a.id }
func (a *Account) UserID() string          { return a.userID }
func (a *Account) Label() string           { return a.label }
func (a *Account) CreatedAt() time.Time    { return a.createdAt }
func (a *Account) UpdatedAt() time.Time    { return a.updatedAt }

func (a *Account) Relabel(label string) error {
	label = strings.TrimSpace(label)
	if label == "" || len(label) > 120 {
		return ErrInvalidAccountLabel
	}
	a.label = label
	a.updatedAt = time.Now().UTC()
	return nil
}
