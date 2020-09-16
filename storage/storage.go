package storage

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"

	"github.com/maragudk/rest-auth-server/model"
)

// Storer stores users in-memory.
type Storer struct {
	users map[string]*model.User
}

func New() *Storer {
	return &Storer{users: map[string]*model.User{}}
}

const (
	minPasswordLength = 10
	maxPasswordLength = 64
)

// Signup a user. Note that the given password is in cleartext and hashed here.
func (s *Storer) Signup(name, password string) error {
	passwordLength := utf8.RuneCountInString(password)
	if passwordLength < minPasswordLength || passwordLength > maxPasswordLength {
		return fmt.Errorf("%v is outside the password length range of [%v,%v]", passwordLength, minPasswordLength, maxPasswordLength)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("could not hash password: %w", err)
	}

	user := model.User{
		Name:     name,
		Password: hashedPassword,
	}

	s.users[name] = &user

	return nil
}

// Login with a given name and password. The password is cleartext and hashed here.
// Returns the user if succesful, without the password.
func (s *Storer) Login(name, password string) (*model.User, error) {
	user, ok := s.users[name]
	if !ok {
		return nil, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, nil
		}
		return nil, err
	}

	user.Password = nil
	return user, nil
}
