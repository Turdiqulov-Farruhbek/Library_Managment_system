package handlers

import (
	"context"
	"getway/genproto/genres" 
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all genres
// @Description Retrieve a list of all genres
// @Tags Genres
// @Produce json
// @Security BearerAuth
// @Success 200 {object} genres.GetGenresResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/genres/ [get]
func (h *HandlerStruct) GetGenresHandler(c *gin.Context) {
	req := &genres.GetGenresRequest{}
	resp, err := h.Genres.GetGenres(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get genre by ID
// @Description Retrieve details of a genre by its ID
// @Tags Genres
// @Produce json
// @Security BearerAuth
// @Param id path string true "Genre Genre ID"
// @Success 200 {object} genres.GetGenreByIdResponse
// @Failure 400 {object} string "Invalid ID"
// @Failure 404 {object} string "Genre Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/genres/{id} [get]
func (h *HandlerStruct) GetGenreByIdHandler(c *gin.Context) {
	id := c.Param("id")
	req := &genres.GetGenreByIdRequest{Id: id}
	resp, err := h.Genres.GetGenreById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get books by genre ID
// @Description Retrieve a list of books belonging to a specific genre
// @Tags Genres
// @Produce json
// @Security BearerAuth
// @Param genre_id path string true "Genre ID"
// @Success 200 {object} genres.GetGenreByIdResponse
// @Failure 400 {object} string "Genre ID"
// @Failure 404 {object} string "Genre Not Found for Genre"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/genres/{genre_id}/books [get]
func (h *HandlerStruct) GetBookGenreHandler(c *gin.Context) {
	genreID := c.Param("genre_id")
	req := &genres.GetGenreByIdRequest{Id: genreID}
	resp, err := h.Genres.GetGenreById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Create a new genre
// @Description Add a new genre to the system
// @Tags Genres
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param genre body genres.CreateGenreRequest true "Genre Data"
// @Success 201 {object} genres.CreateGenreResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/genres/ [post]
func (h *HandlerStruct) CreateGenreHandler(c *gin.Context) {
	var req genres.CreateGenreRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Genres.CreateGenre(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Update genre by ID
// @Description Update details of an existing genre by its ID
// @Tags Genres
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Genre ID"
// @Param genre body genres.UpdateGenreRequest true "Genre Data"
// @Success 200 {object} genres.UpdateGenreResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 404 {object} string "Genre Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/genres/{id} [put]
func (h *HandlerStruct) UpdateGenreHandler(c *gin.Context) {
	id := c.Param("id")
	var req genres.UpdateGenreRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = id
	resp, err := h.Genres.UpdateGenre(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete genre by ID
// @Description Remove a genre from the system by its ID
// @Tags Genres
// @Produce json
// @Security BearerAuth
// @Param id path string true "Genre ID"
// @Success 200 {object} string "Borrower deleted successfully"
// @Failure 400 {object} string "Invalid ID"
// @Failure 404 {object} string "Genre Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/genres/{id} [delete]
func (h *HandlerStruct) DeleteGenreHandler(c *gin.Context) {
	id := c.Param("id")
	req := &genres.DeleteGenreRequest{Id: id}
	_, err := h.Genres.DeleteGenre(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted successfully"})
}
