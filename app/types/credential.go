package types

import (
	"http2/app/types/erors"
)

type Credential struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

func NewCredential(login, email, password string, id int) *Credential {
	cred := &Credential{}
	cred.setID(id)
	cred.setEmail(email)
	cred.setLogin(login)
	cred.setPassword(password)

	return cred
}

func LoginValidate(login, password string) error {
	if len(login) == 0 {
		return erors.Login
	}
	
	if len(password) == 0 {
		return erors.Pass
	}
	return nil
}


func (c *Credential) setID(ID int) error {
	if ID <= 0 {
		return erors.Id
	}

	return nil
}

func (c *Credential) setEmail(email string) error {
	if len(email) == 0 {
		return erors.Email
	}

	c.Email = email

	return nil
}

func (c *Credential) setLogin(login string) error {
	if len(login) == 0 {
		return erors.Login
	}

	c.Login = login

	return nil
}

func (c *Credential) setPassword(password string) error {
	if len(password) == 0 {
		return erors.Pass
	}

	c.Password = password

	return nil
}
