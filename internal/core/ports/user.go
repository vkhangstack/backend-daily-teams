package ports

import (
	"github.com/vkhangstack/dlt/internal/core/domain/dto"
	"github.com/vkhangstack/dlt/internal/core/domain/model"
)

type UserService interface {
	Register(payload dto.RegisterDto) error
	Login(username, password string) (*dto.LoginResponse, error)
	ProfileMe(userId uint64) (*model.User, error)
	GetAccessToken(token string) (string, error)
	UpdateUser(payload *model.UserUpdate, userId uint64) error
	//LoginWithTeams(payload *model.UserLoginSocial) (*model.User, error)
	//DeleteProfileMe(userId string) error
}

type UserRepository interface {
	CreateUser(payload *model.User) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
	ProfileMe(userId uint64) (*model.User, error)
	HashPassword(password string) (string, error)
	VerifyPassword(hash, password string) error
	UpdateUser(payload *model.UserUpdate, userId uint64) error
}
