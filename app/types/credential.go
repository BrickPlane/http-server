package types

import (
	"http2/app/types/erors"
)

type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

func NewUser(login, email, password string, id int) *User {
	user := &User{}
	user.setID(id)
	user.setEmail(email)
	user.setLogin(login)
	user.setPassword(password)

	return user
}

func UserValidate(login, password, email string) error {
	if len(login) == 0 {
		return erors.Login
	}
	
	if len(password) == 0 {
		return erors.Pass
	}

	if len(email) == 0 {
		return erors.Email
	}
	return nil
}


func (c *User) setID(ID int) error {
	if ID <= 0 {
		return erors.Id
	}

	return nil
}

func (c *User) setEmail(email string) error {
	if len(email) == 0 {
		return erors.Email
	}

	c.Email = email

	return nil
}

func (c *User) setLogin(login string) error {
	if len(login) == 0 {
		return erors.Login
	}

	c.Login = login

	return nil
}

func (c *User) setPassword(password string) error {
	if len(password) == 0 {
		return erors.Pass
	}

	c.Password = password

	return nil
}
