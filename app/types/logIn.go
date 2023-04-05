package types

import (
	"http2/app/types/erors"
)

type Credential struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

func NewCred(login, password string) *Credential {
	cred := &Credential{}
	cred.credLogin(login)
	cred.credPassword(password)
	return cred
}

func (c *Credential) credLogin(login string) error {
	if len(login) == 0 {
		return erors.Login
	}

	c.Login = login

	return nil
}

func (c *Credential) credPassword(password string) error {
	if len(password) == 0 {
		return erors.Pass
	}

	c.Password = password

	return nil
}