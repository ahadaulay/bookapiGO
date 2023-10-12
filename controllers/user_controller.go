package controllers

import (
	"NOMOR1/config"
	"NOMOR1/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
	"message": "success get all users",
	"users":   users,
	})
}

// get user by id
func GetUsersByIdController(c echo.Context) error {
	id := c.Param("id")

	var user models.User

	result := config.DB.First(&user,id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound,"User Not Found")
	}

	return c.JSON(http.StatusOK,map[string]interface{}{
		"status" : "success",
		"data" : user,
	})  
}

func CreateUserController(c echo.Context) error {
	var input models.User
	c.Bind(&input)

	
	var user models.User = models.User{
		Name: input.Name,
		Email: input.Email,
		Password: input.Password,
	}

	err := config.DB.Create(&user).Error

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

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here

	var user models.User

	id := c.Param("id")

	result := config.DB.First(&user,id)

	if result.Error != nil{
		return c.JSON(http.StatusNotFound,map[string]interface{}{
			"status" : "failed",
			"erorr" : "ID not found",
		})
	}

	err := config.DB.Delete(&user).Error

	if err != nil {
		return c.JSON(http.StatusNotFound,map[string]interface{}{
			"status" : "failed",
			"erorr" : "Failed Delete",
		})
	}

	return c.JSON(http.StatusAccepted,map[string]interface{}{
		"status" : "success",
		"data" : user,
	})


}

// Update user by ID
func UpdateUserController(c echo.Context) error {
    var user models.User

    id := c.Param("id")

    // Cari user berdasarkan ID
    if err := config.DB.First(&user, id).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "status": "failed",
            "error": "Internal server error",
        })
    }

    // Bind request body ke struct User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "status": "failed",
            "error": "Bad request",
        })
    }

    // Simpan perubahan ke database
    if err := config.DB.Save(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "status": "failed",
            "error": "Failed to update user",
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": "success",
        "data":   user,
    })
}