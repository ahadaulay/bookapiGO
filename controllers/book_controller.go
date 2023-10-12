package controllers

import (
	"NOMOR1/config"
	"NOMOR1/models"
	"net/http"

	"github.com/labstack/echo/v4"
)
func GetAllBook(c echo.Context) error {

	var books []models.Book

	err := config.DB.Find(&books).Error

	if err != nil {
		return c.JSON(500,map[string]interface{}{
			"status" : "failed",
			"message" : err,
		})
	}

	return c.JSON(http.StatusOK,map[string]interface{}{
		"status" : "success",
		"data" : books,
	})
}

func GetBookById(c echo.Context) error  {
	var books models.Book
	
	id := c.Param("id")

	err := config.DB.First(&books,id).Error

	if err != nil{
		return c.JSON(http.StatusNotFound,map[string]interface{}{
			"status" : "failed",
			"message" : "id not found",
		})
	}

	return c.JSON(http.StatusOK,map[string]interface{}{
		"status" : "success",
		"data" : books,
	})
}

func CreateBook(c echo.Context) error {

	var input models.Book
	c.Bind(&input)

	
	var user models.Book = models.Book{
		Title: input.Title,
		Author: input.Author,
		Publisher: input.Publisher,
	}

	err := config.DB.Create(&input).Error

	if err != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"status" : "Not Created",
		})
	}

	return c.JSON(201, map[string]interface{}{
		"status":"created",
		"data" : user,
	})
}


func UpdateBook(c echo.Context) error {
	var books models.Book

    id := c.Param("id")

    // Cari user berdasarkan ID
    if err := config.DB.First(&books, id).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "status": "failed",
            "error": "Internal server error",
        })
    }

    // Bind request body ke struct User
    if err := c.Bind(&books); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "status": "failed",
            "error": "Bad request",
        })
    }

    // Simpan perubahan ke database
    if err := config.DB.Save(&books).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "status": "failed",
            "error": "Failed to update user",
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": "success",
        "data":   books,
    })
}


func DeleteBook(c echo.Context) error {
	var books models.Book

	id := c.Param("id")

	if err := config.DB.First(&books,id).Error ; err != nil  {
		return c.JSON(http.StatusNotFound,map[string]interface{}{
			"status" : "failed",
			"message" : "id not found",
		})
	}

	if err := config.DB.Delete(&books).Error ; err != nil {
		return c.JSON(500,map[string]interface{}{
			"status" : "failed",
			"message" : err,
		})
	}

	return c.NoContent(http.StatusNoContent)
}