package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thedummyuser/pastebin/models"
	"gorm.io/gorm"
	"net/http"
)

type PastePostReq struct {
	Content string `json:"content" validate:"required,min=5"`
}

type PasteResponse struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
	Content string `json:"content"`
}

func GetAllPastes(c echo.Context, db *gorm.DB) error {
	var pastes []models.Paste

	if err := db.Find(&pastes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch pastes",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully fetched all pastes",
		"count":   len(pastes),
		"data":    pastes,
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
