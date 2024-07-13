package api

import (
	"database/sql"

	"github.com/Project_Restaurant/Auth-Service/api/handler.go"
	"github.com/Project_Restaurant/Auth-Service/postgres"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

)

func NewRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	userRepo := postgres.NewUserRepo(db)
	h := handler.NewHandler(db, userRepo)
	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	

	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
	router.GET("/users/:username", h.GetByUsername)

	return router
}
