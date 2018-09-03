package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nathanburkett/nathanb-server/datamodels"
	"golang.org/x/crypto/bcrypt"
)

type AuthClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// AuthPayload authentication payload
type AuthPayload struct {
	Token string `json:"token"`
}

// AuthParams Authentication params
type AuthParams struct {
	Username string
	Password string
}

type TokenRecordFetcher interface {
	FindByUsername(username string) (datamodels.User, error)
}

// AuthService Authentication service
type AuthService struct {
	Fetcher TokenRecordFetcher
}

// Attempt Attempt to authenticate
func (s *AuthService) Attempt(params *AuthParams) (*AuthPayload, error) {
	u, authErr := s.authenticate(params)

	if authErr != nil {
		return nil, authErr
	}

	token := s.createToken(u)

	signedToken, signErr := s.signToken(token)

	if signErr != nil {
		return nil, signErr
	}

	payload := &AuthPayload{
		signedToken,
	}

	return payload, nil
}

// Validate validate a token
func (s *AuthService) Validate(token *jwt.Token) bool {
	claims, ok := token.Claims.(AuthClaims)

	if !ok || !token.Valid {
		return false
	}

	user, err := s.GetUser(claims.Username)

	if err != nil || user == nil {
		return false
	}

	return true
}

// GetUser get the user
func (s *AuthService) GetUser(username string) (*datamodels.User, error) {
	u, err := s.Fetcher.FindByUsername(username)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// KeyFn Get the tokenizing keying function
func (s *AuthService) KeyFn(token *jwt.Token) (interface{}, error) {
	signingKey, exists := os.LookupEnv("APP_KEY")
	if !exists {
		return nil, errors.New("`APP_KEY` does not exist")
	}

	return []byte(signingKey), nil
}

func (s *AuthService) authenticate(params *AuthParams) (*datamodels.User, error) {
	if params.Password == "" {
		return nil, errors.New("password cannot be empty")
	}

	user, err := s.GetUser(params.Username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("no user for '%s'", params.Username)
	}

	if pwErr := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(params.Password)); pwErr != nil {
		return nil, pwErr
	}

	return user, nil
}

func (s *AuthService) createToken(u *datamodels.User) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		u.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 7).UnixNano(),
		},
	})

	return token
}

func (s *AuthService) signToken(token *jwt.Token) (string, error) {
	key, err := s.KeyFn(token)

	if err != nil {
		return "", errors.New("could not sign token")
	}

	return token.SignedString(key)
}
