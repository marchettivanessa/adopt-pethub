package handler

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
	"adopt-pethub/backend/logging"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type feedbackHandler struct {
	feedbackDbMethods RepositoryFeedback
}

type RepositoryFeedback interface {
	GetFeedbacks(db *database.Database) ([]domain.Feedback, error)
	InsertFeedback(feedback domain.Feedback, db *database.Database) error
}

func NewFeedbackHandler(repo RepositoryFeedback) *feedbackHandler {
	return &feedbackHandler{
		feedbackDbMethods: repo,
	}
}

func (h *feedbackHandler) GetFeedbacks(c echo.Context, ) error {
	db := c.Get("db").(*database.Database)
	feedbacks, err := h.feedbackDbMethods.GetFeedbacks(db)
	if err != nil {
		return errors.New("failed to get feedbacks from database")
	}
	return c.JSON(http.StatusOK, feedbacks)
}

func (h *feedbackHandler) CreateFeedback(c echo.Context) error {
	var feedback domain.Feedback
	db := c.Get("db").(*database.Database)

	log := logging.Logger(map[string]interface{}{
		"project": "adopt-pethub",
		"package": "feedback_handler",
	})

	err := c.Bind(&feedback)
	if err != nil {
		log.WithError(err).Error("failed to bind feedback data")

	}

	err = h.feedbackDbMethods.InsertFeedback(feedback, db)
	if err != nil {
		log.WithError(err).Error("failed to insert feedback into database")
		return c.JSON(http.StatusInternalServerError, "Failed to insert feedback")
	}

	return c.JSON(http.StatusCreated, "Feedback created")
}
