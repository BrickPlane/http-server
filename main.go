package main

import (
	// "fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserType struct{
	ID string `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Type string `json:"type" binding:"required"`
}

var users = []UserType{
    {ID: "1", Name: "John", Email: "admin@mail", Type: "Admin"},
    {ID: "2", Name: "Bob", Email: "player@mail", Type: "Player"},
}

func main(){
	router :=gin.Default()
	router.GET("/users", getUsers)
	router.GET("/id", getUserByID)
	router.POST("/users", addUser)
	router.PATCH("/ch", updateUser)
	router.DELETE("/delete", deleteUsers)
	router.Run("localhost:8083")
	
}
// TODO: create rute getUserByID 0.0.0.0:8083/users?id=32 ✓
// TODO: add error handler result ✓
// TODO: add validation input data ✓
// TODO: PATCH must update record not create ✓
// TODO: Delete user ✓

func getUsers(c *gin.Context) {
	fmt.Println("in users")
	c.IndentedJSON(http.StatusOK, users)
}

func addUser(c *gin.Context) {
	var addUser UserType
	if err:= c.BindJSON(&addUser); err!= nil {
		c.IndentedJSON(http.StatusInternalServerError, "Cant parse input data")
		return
	} 
	for _, a := range users{
		if a.ID == addUser.ID {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "user already exist"})
			return
		}
	}
	users = append(users, addUser)
	c.IndentedJSON(http.StatusCreated, addUser)
}

func getUserByID(c *gin.Context){
	id := c.Query("id")
	fmt.Println("in get")
	for _, user := range users{
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func updateUser(c *gin.Context) {
	id := c.Query("id")
	var updateData UserType
	if error := c.BindJSON(&updateData); error != nil{
		return
	}
	for ch, UserType := range users{
		if UserType.ID == id{
			users[ch] = updateData
			c.IndentedJSON(http.StatusOK, users[ch])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func RemoveIndex(s []UserType, index int) []UserType{
	return append(s[:index], s[index+1:]...)
}

func deleteUsers(c *gin.Context){
	id := c.Query("id")
	for del, UserType := range users{
		if UserType.ID == id{
			users = RemoveIndex(users, del)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "user exterminated"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

