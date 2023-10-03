package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/middleware"
	"github.com/yoda/web/internal/storage"
	"time"
)

type IAuthStorage interface {
	GetPermissionByUserId(id int32) (*[]string, error)
	GetProleByLogin(login string) (*storage.UserProfile, error)
	UserExists(id int32) bool
}

type AuthService struct {
	storage IAuthStorage
}

func NewAuthService(storage IAuthStorage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) GetPermissionByUserId(id int32) (*[]string, error) {
	return s.storage.GetPermissionByUserId(id)
}

func (s *AuthService) CreateToken(login *api.LoginInfo) (string, time.Time, error) {
	profile, err := s.storage.GetProleByLogin(login.Email)
	if err != nil {
		return "", time.Time{}, errors.Errorf("failed to get profile by login: %s", err)
	}
	if profile.Password != login.Password {
		return "", time.Time{}, errors.Errorf("wrong password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &middleware.Claims{
		UserName:    profile.UserName,
		UserId:      profile.UserId,
		UserEmail:   profile.Email,
		Permissions: profile.Permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(middleware.JwtKey)
	return accessToken, expirationTime, err
}

func (s *AuthService) CheckUserById(id int32) bool {
	return s.storage.UserExists(id)
}

func (s *AuthService) GetProfile(auth string) (*api.Profile, error) {
	token, err := middleware.ParsToken(auth)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*middleware.Claims)
	permissions := make([]api.Permission, len(claims.Permissions))
	for i, permission := range claims.Permissions {
		permissions[i] = api.Permission(permission)
	}
	return &api.Profile{
		Email:       claims.UserEmail,
		Name:        claims.UserName,
		Permissions: permissions,
	}, nil
}
