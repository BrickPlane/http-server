package storage

import (
	"errors"
	
	"github.com/golang-jwt/jwt"
) 

type Credential struct {
	Id string 		 `json:"id"`
	Login string	 `json:"login"`
	Password string  `json:"password"`
	Phone int		 `json:"phone"`
	Types string	 `json:"type"`
}

func NewCredential(id, login, password, types string, phone int) (*Credential, error) {
	cred := &Credential{}
	cred.setID(id)
	cred.setLogin(login)
	cred.setPassword(password)
	cred.setPhone(phone)
	cred.setType(types)	
	
	return cred, nil
}

func (c *Credential) setID(id string) error {
	if len(id) == 0 {
		return  errors.New("ID field is required")
	}
	
	c.Id = id
	return nil
}

func (c *Credential) setLogin(login string) error {
	if len(login) == 0 {
		return errors.New("Login field is required")
	}

	c.Login = login

	return nil
}

func (c *Credential) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password field is required")
	}

	c.Password = password

	return nil
}

func (c *Credential) setPhone(phone int) error {
	if phone <= 0 {
		return errors.New("Phone field is required")
	}

	c.Phone = phone

	return nil
}

func (c *Credential) setType(types string) error {
	if len(types) == 0 {
		return errors.New("Type field is required")
	}

	c.Types = types

	return nil
}

type Claims struct{
	Login string	 `json:"login" binding:"required"`
	Password string  `json:"password" binding:"required"`
	jwt.StandardClaims
}

var Users = []Credential{
	{Id: "1", Login: "admin", Password: "admin1", Phone: 1234, Types: "Admin"},
	{Id: "2", Login: "john", Password: "qwe123", Phone: 4567, Types: "Player"},
}