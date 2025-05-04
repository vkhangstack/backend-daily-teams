package repository

import (
	"errors"
	"fmt"
	"github.com/vkhangstack/dlt/internal/core/domain/enum"
	"github.com/vkhangstack/dlt/internal/core/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (u *DB) CreateUser(payload *model.User) (*model.User, error) {

	user := payload
	//req := u.db.First(&user, "email = ?", email)
	//if req.RowsAffected != 0 {
	//	return nil, errors.New("user already exists")
	//}

	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//if err != nil {w
	//	return nil, fmt.Errorf("password not hashed: %v", err)
	//}

	req := u.db.Create(&user)
	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("user not saved: %v", req.Error)
	}
	return user, nil
}

func (u *DB) ProfileMe(userId uint64) (*model.User, error) {
	user := &model.User{}
	//cachekey := user.ID
	//err := u.cache.Get(cachekey, &user)
	//if err == nil {
	//	return user, nil
	//}

	req := u.db.First(&user, "id = ? and status = ?", userId, enum.ActiveUser)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	//err = u.cache.Set(cachekey, user, time.Minute*10)
	//if err != nil {
	//	fmt.Printf("Error storing user in cache: %v", err)
	//}
	return user, nil
}

func (u *DB) ReadUsers() ([]*model.User, error) {
	var users []*model.User

	req := u.db.Find(&users)
	if req.Error != nil {
		return nil, fmt.Errorf("users not found: %v", req.Error)
	}

	return users, nil
}

func (u *DB) UpdateUser(id, email, password string) error {
	user := &model.User{}
	req := u.db.First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}

	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//if err != nil {
	//	return fmt.Errorf("password not hashed: %v", err)
	//}

	// user = &domain.User{
	// 	Email:    email,
	// 	Password: string(hashedPassword),
	// }

	// Update the email and password fields of the user
	user.Email = email
	//user.Password = string(hashedPassword)

	req = u.db.Model(&user).Where("id = ?", id).Update(user)
	if req.RowsAffected == 0 {
		return errors.New("unable to update user :(")
	}

	// delete user in the cache
	//err := u.cache.Delete(id)
	//if err != nil {
	//	fmt.Printf("Error deleting user in cache: %v", err)
	//}

	return nil

}

func (u *DB) DeleteUser(id string) error {
	user := &model.User{}
	req := u.db.Where("id = ?", id).Delete(&user)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}
	//err := u.cache.Delete(id)
	//if err != nil {
	//	fmt.Printf("Error deleting user in cache: %v", err)
	//}
	return nil
}

//func (u *DB) LoginUserSocial(email string) (*model.User, error) {
//	//apiCfg, err := LoadAPIConfig()
//	//if err != nil {
//	//	return nil, err
//	//}
//
//	user, err := u.findUserByEmail(email)
//	if err != nil {
//		return nil, err
//	}
//
//	//err = u.VerifyPassword(user.Password, password)
//	//if err != nil {
//	//	return nil, err
//	//}
//
//	//accessToken, err := u.generateAccessToken(user.ID, apiCfg.JWTSecret)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//
//	//refreshToken, err := u.generateRefreshToken(user.ID, apiCfg.JWTSecret)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//
//	//return &LoginResponse{
//	//	ID:           user.ID,
//	//	Email:        user.Email,
//	//	AccessToken:  accessToken,
//	//	RefreshToken: refreshToken,
//	//	//Membership:   user.Membership,
//	//}, nil
//	return user, nil
//}

func (u *DB) UpdateMembershipStatus(id string, membership bool) error {
	user := &model.User{}
	req := u.db.First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}

	//user = &model.User{
	//	Membership: membership,
	//}
	req = u.db.Model(&user).Where("id = ?", id).Update(user)
	if req.RowsAffected == 0 {
		return errors.New("unable to update membership status :(")
	}
	return nil
}

func (u *DB) FindUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	req := u.db.First(&user, "email = ?", email)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}
func (u *DB) FindUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	req := u.db.First(&user, "username = ?", username)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u *DB) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("password not hashed: %v", err)
	}

	return string(hashedPassword), nil
}

func (u *DB) VerifyPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password not matched")
	}
	return nil
}
