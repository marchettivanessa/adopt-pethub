package api

// import (
// 	"adopt-pethub/backend/database"
// 	"adopt-pethub/backend/handler"
// 	"adopt-pethub/backend/repository"
// 	"github.com/labstack/echo"
// )

// func RegisterFeedbackRoutes(e *echo.Echo, db *database.Database) {
// 	// Instancia o repositório e o handler de feedback
// 	feedbackRepo := repository.FeedbackRepository{}
// 	feedbackHandler := handler.NewFeedbackHandler(feedbackRepo)

// 	// Define as rotas do domínio feedback
// 	e.POST("/feedback", func(c echo.Context) error {
// 		return feedbackHandler.CreateFeedback(c, db)
// 	})
// 	e.GET("/feedback/:id", func(c echo.Context) error {
// 		return feedbackHandler.GetFeedbackById(c, db)
// 	})
// }