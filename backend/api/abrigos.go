package api

// import (
// 	"adopt-pethub/backend/database"
// 	"adopt-pethub/backend/middleware"

// 	"github.com/labstack/echo/v4"
// )

// func RegisterAbrigoRoutes(e *echo.Echo, db *database.Database) {
// 	abrigoRepo := &repository.AbrigoRepository{}
// 	abrigoHandler := handler.NewAbrigoHandler(abrigoRepo)

// 	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			c.Set("db", db)
// 			return next(c)
// 		}
// 	})
// 	e.GET("/abrigo", abrigoHandler.GetAbrigos, middleware.AuthMiddleware)
// 	e.POST("/abrigo", abrigoHandler.CreateAbrigo, middleware.AuthMiddleware)
// }
