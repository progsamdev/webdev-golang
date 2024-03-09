package models

import (
	"database/sql"
	"fmt"

	"github.com/progsamdev/coursescalhoun/rand"
)

const (
	//The minimum number of bytes to be used for each session token.
	MinBytesPerToken = 32
)

type Session struct {
	ID              string
	UserID          string
	NewSessionToken string
	TokenHash       string
}

type SessionService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// each session token. If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used.
	BytesPerToken int
}

func (ss *SessionService) Create(userID string) (*Session, error) {
	if ss.BytesPerToken < MinBytesPerToken {
		ss.BytesPerToken = MinBytesPerToken
	}
	token, err := rand.Strings(ss.BytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	//TODO: hash de session toekn
	session := Session{
		UserID:          userID,
		NewSessionToken: token,
	}
	//store the session in our DB
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	//1. Create the session token
	//2.

	return nil, nil

}