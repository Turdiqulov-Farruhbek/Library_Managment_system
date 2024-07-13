package api

import (
	"getway/api/handlers"
	// "getway/api/middlerware"

	_ "getway/docs"
	_ "getway/genproto/authors"
	_ "getway/genproto/books"
	_ "getway/genproto/borrowers"
	_ "getway/genproto/genres"
	_ "getway/genproto/users"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Library Management System API
// @version 1.0
// @description API for managing Library Management System resources
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(l_conn *grpc.ClientConn) *gin.Engine {
	handler := handlers.NewHandler(l_conn)

	router := gin.Default()
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.Use(middlerware.Authorizations)
	router.Use(cors.Default())
	// router.Use(corsMiddleware())
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"}, // Adjust for your specific origins
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	authors := router.Group("api/authors")
	{
		authors.GET("/", handler.GetAuthorsHandler)
		authors.GET("/:id", handler.GetAuthorByIdHandler)
		authors.POST("/", handler.CreateAuthorHandler)
		authors.PUT("/:id", handler.UpdateAuthorHandler)
		authors.DELETE("/:id", handler.DeleteAuthorHandler)
	}

	books := router.Group("api/books")
	{
		books.GET("/", handler.GetBooksHandler)
		books.GET("/:id", handler.GetBookByIdHandler)
		books.GET("/authors/:author_id/books", handler.GetBooksAuthorHandler)
		books.GET("/genres/:genre_id/books", handler.GetBooksGenreHandler)
		books.GET("/overdue", handler.GetBooksOverdueHandler)
		books.POST("/", handler.CreateBookHandler)
		books.PUT("/:id", handler.UpdateBookHandler)
		books.DELETE("/:id", handler.DeleteBookHandler)
	}

	borrowers := router.Group("api/borrowers")
	{
		borrowers.GET("/", handler.GetBorrowersHandler)
		borrowers.GET("/:id", handler.GetBorrowerByIdHandler)
		borrowers.GET("/users/:user_id/borrowed_books", handler.GetBorrowersBookHandler)
		borrowers.GET("/users/:user_id/borrowing_history", handler.GetBorrowingHistoryHandler)
		borrowers.POST("/", handler.CreateBorrowerHandler)
		borrowers.PUT("/:id", handler.UpdateBorrowerHandler)
		borrowers.DELETE("/:id", handler.DeleteBorrowerHandler)
	}

	genres := router.Group("api/genres")
	{
		genres.GET("/", handler.GetGenresHandler)
		genres.GET("/:id", handler.GetGenreByIdHandler)
		genres.GET("/genres/:genre_id/books", handler.GetBookGenreHandler)
		genres.POST("/", handler.CreateGenreHandler)
		genres.PUT("/:id", handler.UpdateGenreHandler)
		genres.DELETE("/:id", handler.DeleteGenreHandler)
	}

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
