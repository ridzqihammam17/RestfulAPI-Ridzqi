package controllers

import (
	"net/http"
	"restfulapi/ridzqi/lib/database"
	"restfulapi/ridzqi/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, e := database.GetBooks()

	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	book, e := database.GetBook(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {
	var book models.Book

	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	book, e := database.CreateBook(book)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	var newBook models.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	if err := c.Bind(&newBook); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	book, e := database.UpdateBook(id, newBook)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book by id",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	book, e := database.DeleteBook(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book by id",
		"book":    book,
	})
}
