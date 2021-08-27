package Controllers

import (
	"resource-api/Models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

//GetClients ... Get all clients
func GetClients(c *gin.Context) {
	var client []Models.Client
	err := Models.GetAllClients(&client)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, client)
	}
}
//CreateClient ... Create Client
func CreateClient(c *gin.Context) {
	var client Models.Client
	c.BindJSON(&client)
	err := Models.CreateClient(&client)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, client)
	}
}
//GetClientByID ... Get the client by id
func GetClientByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var client Models.Client
	err := Models.GetClientByID(&client, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, client)
	}
}
//UpdateClient ... Update the client information
func UpdateClient(c *gin.Context) {
	var client Models.Client
	id := c.Params.ByName("id")
	err := Models.GetClientByID(&client, id)
	if err != nil {
		c.JSON(http.StatusNotFound, client)
	}
	c.BindJSON(&client)
	err = Models.UpdateClient(&client, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, client)
	}
}
//DeleteClient ... Delete the client
func DeleteClient(c *gin.Context) {
	var client Models.Client
	id := c.Params.ByName("id")
	err := Models.DeleteClient(&client, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}