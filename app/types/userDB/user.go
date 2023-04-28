package user_types

import (
	"http2/app/types/errors"
)

type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
	Wallet   float64 `json:"wallet" db:"wallet"`
}

func NewUser(login, email, password string, id int, wallet float64) *User {
	user := &User{}
	user.setID(id)
	user.setEmail(email)
	user.setLogin(login)
	user.setPassword(password)
	user.setWallet(wallet)

	return user
}

func UserValidate(login, password, email string) error {
	if len(login) == 0 {
		return errors.Login
	}

	if len(password) == 0 {
		return errors.Pass
	}

	if len(email) == 0 {
		return errors.Email
	}
	return nil
}

func (c *User) setID(ID int) error {
	if ID <= 0 {
		return errors.Id
	}

	return nil
}

func (c *User) setEmail(email string) error {
	if len(email) == 0 {
		return errors.Email
	}

	c.Email = email

	return nil
}

func (c *User) setLogin(login string) error {
	if len(login) == 0 {
		return errors.Login
	}

	c.Login = login

	return nil
}

func (c *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.Pass
	}

	c.Password = password

	return nil
}

func (c *User) setWallet(wallet float64) error {
	if wallet <= 0 {
		return errors.Wallet
	}
	c.Wallet = wallet

	return nil
}
