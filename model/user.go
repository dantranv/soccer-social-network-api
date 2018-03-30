package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	*gorm.Model
	UserName string `gorm:"column:user_name;not null"`
	Email    string `gorm:"column:email;not null"`
	Password string `gorm:"column:password;not null"`
	FullName string `gorm:"column:full_name;not null"`
	Birthday string `gorm:"column:birthday;not null"`
}

// HashAndSaltPassword encrypt password
func (u User) HashAndSaltPassword() User {
	hashAndSaltPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashAndSaltPassword)
	return u
}

// CompareHashAndPassword map the password with current hash password of user
func (u User) CompareHashAndPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

// GetCustomClaims get customs claims
func (u User) GetCustomClaims() map[string]interface{} {
	claims := make(map[string]interface{})
	userclaim := struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		UserName string `json:"user_name"`
	}{
		ID:       u.ID,
		Email:    u.Email,
		UserName: u.UserName,
	}
	claims["user"] = userclaim
	return claims
}

// GetIdentifier get identifier function
func (u User) GetIdentifier() uint {
	return u.ID
}