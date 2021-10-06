package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Context) error {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	if id < len(users) && id > -1 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "success get user by id",
			"users":    users[id],
		})
	}
	return c.String(http.StatusNotFound, "Not Found")
}

func DeleteUsersController(c echo.Context) error {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	if id < len(users) && id > -1 {
		for i := id; i < len(users)-1; i++ {
			users[i] = users[i+1]
		}
		users = users[:len(users)-1]
		return c.String(http.StatusOK, "success delete user by id")
	}
	return c.String(http.StatusNotFound, "Not Found")
}

func UpdateUsersController(c echo.Context) error {
	ids := c.Param("id")
	user := User{}
	id, _ := strconv.Atoi(ids)
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if id < len(users) && id > -1 {
		users[id] = user
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success update user by id",
			"id":       users[id].Id,
			"name":     users[id].Name,
			"email":    users[id].Email,
			"password": users[id].Password,
		})
	}
	return c.String(http.StatusNotFound, "Not Found")
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	// Controller
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)

	// Server Posrt
	e.Logger.Fatal(e.Start(":1323"))
}
