package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vkhangstack/dlt/internal/adapters/repository"
	"github.com/vkhangstack/dlt/internal/adapters/utils"
	"github.com/vkhangstack/dlt/internal/core/domain/dto"
	"github.com/vkhangstack/dlt/internal/core/domain/enum"
	"github.com/vkhangstack/dlt/internal/core/domain/model"
	"github.com/vkhangstack/dlt/internal/core/ports"
	"time"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func generateAccessToken(userID, jwtSecret string) (string, error) {
	hash := sha256.Sum256([]byte(userID + jwtSecret))

	claims := jwt.RegisteredClaims{
		Issuer:    "dtl-access",
		Subject:   userID,
		ID:        hex.EncodeToString(hash[:]),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour).UTC()), // 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func generateRefreshToken(userID, jwtSecret string) (string, error) {
	hash := sha256.Sum256([]byte(userID + jwtSecret))
	claims := jwt.RegisteredClaims{
		Issuer:    "dtl-refresh",
		Subject:   userID,
		ID:        hex.EncodeToString(hash[:]),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour).UTC()), //7d
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func (u *UserService) ProfileMe(id uint64) (*model.User, error) {
	return u.repo.ProfileMe(id)
}

func (u *UserService) GetAccessToken(token string) (string, error) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		return "", err
	}

	userId, err := utils.ValidateRefreshToken(token, apiCfg.JWTSecret)

	if err != nil {
		return "", err
	}

	accessToken, err := generateAccessToken(userId, apiCfg.JWTSecret)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (u *UserService) Register(payload *dto.RegisterDto) error {
	_, err := u.repo.FindUserByUsername(payload.Username)

	if err == nil {
		return fmt.Errorf("user already exist")
	}

	password, err := u.repo.HashPassword(payload.Password)
	if err != nil {
		return fmt.Errorf("hash password")
	}

	ID := utils.GenerateID()
	_, err = u.repo.CreateUser(&model.User{
		SqlModel: &model.SqlModel{
			ID: ID,
		},
		Username:   payload.Username,
		Email:      "",
		FirstName:  payload.FirstName,
		LastName:   payload.LastName,
		Phone:      "",
		Password:   password,
		ProviderId: "",
		AvatarURL:  "",
		Address:    "",
		Status:     int8(enum.ActiveUser),
	})

	return nil
}

func (u *UserService) Login(username, password string) (*dto.LoginResponse, error) {
	user, err := u.repo.FindUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	err = u.repo.VerifyPassword(user.Password, password)
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		return nil, err
	}

	accessToken, err := generateAccessToken(utils.TransformUInt64ToString(user.ID), apiCfg.JWTSecret)
	if err != nil {
		return nil, err
	}
	refreshToken, err := generateRefreshToken(utils.TransformUInt64ToString(user.ID), apiCfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		ID:           user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
