package services

import (
	"os"

	supabaseAuth "github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
	"github.com/supabase-community/supabase-go"
)

type IAuthService interface {
	SignInWithEmailPassword(string, string) (Token, RefreshToken, error)
	SignUpWithEmailPassword(string, string) (Token, RefreshToken, error)
}

type (
	Token        string
	RefreshToken string
)

type AuthError string

const (
	AuthErrorInvalidCredentials AuthError = "invalid credentials"
	AuthErrorUserNotFound       AuthError = "user not found"
	AuthErrorUserAlreadyExists  AuthError = "user already exists"
	AuthErrorUnauthorized       AuthError = "unauthorized"
	AuthSystemError             AuthError = "system error"
)

func (e AuthError) Error() string {
	return string(e)
}

type SupabaseAuthService struct {
	db         *supabase.Client
	authClient supabaseAuth.Client
}

func NewSupabaseAuthService(db *supabase.Client) *SupabaseAuthService {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_PUB_KEY")
	auth := supabaseAuth.New(url, key)
	return &SupabaseAuthService{db: db, authClient: auth}
}

func (a *SupabaseAuthService) SignInWithEmailPassword(email, password string) (Token, RefreshToken, error) {
	if email == "" || password == "" {
		return "", "", AuthErrorInvalidCredentials
	}

	resp, err := a.authClient.Token(types.TokenRequest{
		GrantType: "password",
		Email:     email,
		Password:  password,
	})
	if err != nil {
		return "", "", AuthErrorUserNotFound
	}
	return Token(resp.AccessToken), RefreshToken(resp.RefreshToken), nil
}

func (a *SupabaseAuthService) SignUpWithEmailPassword(email, password string) (Token, RefreshToken, error) {
	if email == "" || password == "" {
		return "", "", AuthErrorInvalidCredentials
	}
	resp, err := a.authClient.Signup(types.SignupRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", "", AuthSystemError
	}
	return Token(resp.AccessToken), RefreshToken(resp.RefreshToken), nil
}

func (a *SupabaseAuthService) RefreshToken(refreshToken string) (string, error) {
	a.authClient.Reauthenticate()
	return "", nil
}
