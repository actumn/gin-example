package data

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService interface {
	User(string, string) (*User, error)
	CreateUser(string, string) error
}

func (database *Database) User(userName, password string) (*User, error) {
	// search general user
	var user User
	if database.Where(&User{UserName: userName}).Find(&user).RecordNotFound() {
		log.Println("UserService: User Not found")
		return nil, ErrFailedAuth
	}

	// password check
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		log.Println("UserService: Password not matched")
		return nil, ErrFailedAuth
	}

	return &user, nil
}

func (database *Database) CreateUser(userName, password string) error {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// password error handling
		log.Println(err)
		return err
	}
	newUser := User{
		UserName: userName,
		Password: passwordBytes,
	}
	if err := database.Create(newUser).Error; err != nil {
		// database error handling
		log.Println(err)
		return err
	}

	return nil
}
