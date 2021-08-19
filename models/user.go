package models

import (
    "golang.org/x/crypto/bcrypt"
)

type User struct {
  ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}


// CreateUserRecord creates a user record in the database
func (user *User) CreateUserRecord() error {
	result := DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}


// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
