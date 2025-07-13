package handlers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thedummyuser/pastebin/models"
	"gorm.io/gorm"
)

type PastePostReq struct {
	Content string `json:"content" validate:"required,min=5"`
}

type PasteResponse struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
	Content string `json:"content"`
}

func GetSinglePost(c echo.Context, db *gorm.DB) error {
	uuid := c.Param("uuid")
	var paste models.Paste

	if uuid == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please provide uuid before calling the api",
		})
	}

	result := db.First(&paste, "uuid = ?", uuid)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Paste not found",
		})
	} else if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Something went wrong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Fetched the paste",
		"data":    paste,
	})
}

func Paste(c echo.Context, db *gorm.DB) error {
	var body PastePostReq

	// Bind JSON request body
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid content input"})
	}

	// Validate the request
	if err := c.Validate(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Create new paste instance
	paste := models.Paste{
		UUID:    uuid.New().String(), // Generate UUID
		Content: body.Content,
	}

	// Save to database
	if err := db.Create(&paste).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to save paste"})
	}

	// Return success response
	response := PasteResponse{
		Message: "paste created successfully",
		UUID:    paste.UUID,
		Content: paste.Content,
	}

	return c.JSON(http.StatusCreated, response)
}
