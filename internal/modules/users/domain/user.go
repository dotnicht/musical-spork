package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrInvalidEmail     = errors.New("invalid email")
	ErrInvalidUserName  = errors.New("invalid name")
	ErrEmailAlreadyUsed = errors.New("email already used")
)

type UserID string

func NewUserID() UserID { return UserID(uuid.NewString()) }

type User struct {
	id        UserID
	email     string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func CreateNewUser(email, name string) (*User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	name = strings.TrimSpace(name)

	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}
	if name == "" || len(name) > 120 {
		return nil, ErrInvalidUserName
	}

	now := time.Now().UTC()
	return &User{
		id:        NewUserID(),
		email:     email,
		name:      name,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func RehydrateUser(id UserID, email, name string, createdAt, updatedAt time.Time) *User {
	return &User{
		id:        id,
		email:     email,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) ID() UserID           { return u.id }
func (u *User) Email() string        { return u.email }
func (u *User) Name() string         { return u.name }
func (u *User) CreatedAt() time.Time { return u.createdAt }
func (u *User) UpdatedAt() time.Time { return u.updatedAt }

func (u *User) Rename(newName string) error {
	newName = strings.TrimSpace(newName)
	if newName == "" || len(newName) > 120 {
		return ErrInvalidUserName
	}
	u.name = newName
	u.updatedAt = time.Now().UTC()
	return nil
}

func (u *User) ChangeEmail(newEmail string) error {
	newEmail = strings.TrimSpace(strings.ToLower(newEmail))
	if !isValidEmail(newEmail) {
		return ErrInvalidEmail
	}
	u.email = newEmail
	u.updatedAt = time.Now().UTC()
	return nil
}

var emailRe = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

func isValidEmail(s string) bool { return emailRe.MatchString(s) }
