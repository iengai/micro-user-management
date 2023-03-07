package entity

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
	"time"
)

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) IsValid() bool {
	return u.Username != "" && u.PasswordHash != ""
}

// NewUser creates a new user with the given username and password.
func NewUser(username string, password string) (*User, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}
	if password == "" {
		return nil, errors.New("password is required")
	}

	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	sha := sha256.New()
	sha.Write([]byte(password))
	sha.Write(salt)
	hash := sha.Sum(nil)

	encodedHash := base64.StdEncoding.EncodeToString(hash)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)

	return &User{
		Username:     username,
		PasswordHash: encodedSalt + ":" + encodedHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

// CheckPasswordHash checks if the given password matches the user's password hash.
func (u *User) CheckPasswordHash(password string) bool {
	parts := strings.Split(u.PasswordHash, ":")
	if len(parts) != 2 {
		return false
	}

	encodedSalt, encodedHash := parts[0], parts[1]
	salt, err := base64.StdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return false
	}

	sha := sha256.New()
	sha.Write([]byte(password))
	sha.Write(salt)
	hash := sha.Sum(nil)

	encodedHash2 := base64.StdEncoding.EncodeToString(hash)
	return encodedHash == encodedHash2
}
