package handler

import (
	"database/sql"
	"net/http"

	"github.com/Project_Restaurant/Auth-Service/models"
	"github.com/Project_Restaurant/Auth-Service/postgres"
	"github.com/Project_Restaurant/Auth-Service/token"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db   *sql.DB
	User *postgres.UserRepo
}

func NewHandler(db *sql.DB, user *postgres.UserRepo) *Handler {
	return &Handler{
		db:   db,
		User: user,
	}
}

// @Summary Register new user
// @Description Register a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserRegister true "User data"
// @Success 200 {object} models.LoginRes
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /register [post]
// @Security ApiKeyAuth
func (h *Handler) Register(c *gin.Context) {
	var user models.UserRegister
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.User.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	token, err := token.CreateToken(res.Name, res.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary Login user
// @Description Log in an existing user with the provided credentials
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserLogin true "User data"
// @Success 200 {object} models.LoginRes
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /login [post]
// @Security ApiKeyAuth
func (h *Handler) Login(c *gin.Context) {
	var user models.UserLogin
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := token.CreateToken(res.Name, res.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary Get user by username
// @Description Retrieve a user's details by their username
// @Tags users
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.User
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /users/{username} [get]
// @Security ApiKeyAuth
func (h *Handler) GetByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := h.User.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
