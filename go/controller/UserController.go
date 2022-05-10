package controller

import (
	"github.com/gin-gonic/gin"
	"go-admin/go/entity"
	"go-admin/go/service"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	err := service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": user})
	}

}

func GetUserList(c *gin.Context) {
	userList, err := service.GetUserList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": userList})
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.BindJSON(&user)
	if err = service.UpdateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": user})
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteUserById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
	}
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": user})
	}
}
