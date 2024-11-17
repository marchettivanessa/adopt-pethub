package api

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/handler"
	"adopt-pethub/backend/middleware"
	"adopt-pethub/backend/repository"

	"github.com/labstack/echo/v4"
)

func RegisterFeedbackRoutes(e *echo.Echo, db *database.Database) {
	feedbackRepo := &repository.FeedbackRepository{}
	feedbackHandler := handler.NewFeedbackHandler(feedbackRepo)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})
	e.GET("/feedback", feedbackHandler.GetFeedbacks, middleware.AuthMiddleware)
	e.POST("/feedback", feedbackHandler.CreateFeedback, middleware.AuthMiddleware)
}
