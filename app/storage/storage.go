package storage

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Storage struct {
}

func NewStorage() *Storage {
	return new(Storage)
}
// ?? 
type Map map[string]string
// ??
var Users = []Credential{}

func (storage *Storage) StorageIn(c *gin.Context, data string) error {
	var Users = []Map{
		{"Id": "1", "Login": "admin", "Password": "admin1", "Phone": "1234", "Types": "Admin"},
		// {Id: "2", Login: "john", Password: "qwe123", Phone: 4567, Types: "Player"},
	}
	fmt.Println("users", Users)
	return nil
}

// credential


type Credential struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewCredential(id, login, password, types string, phone int) (*Credential, error) {
	cred := &Credential{}
	cred.setID(id)
	cred.setLogin(login)
	cred.setPassword(password)

	return cred, nil
}

func (c *Credential) MyTestFunc(data int) error {
	fmt.Println("... Hello data : ", data)

	return nil
}

func (c *Credential) setID(id string) error {
	if len(id) == 0 {
		return errors.New("ID field is required")
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

type Claims struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	jwt.StandardClaims
}


