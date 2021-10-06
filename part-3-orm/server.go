package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "db_alta_orm",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

func GetUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	user := User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	if err := DB.First(&user, id).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := DB.Save(&user).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	// user, err := strconv.Atoi(c.Param())
	user := User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	if err := DB.First(&user, id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := DB.Delete(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete by user id",
		"user":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	user := User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	if err := DB.First(&user, id).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id",
		"user":    user,
	})
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
