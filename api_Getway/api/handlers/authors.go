package handlers

import (
	"context"
	"getway/genproto/authors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all authors
// @Description Retrieve a list of all authors
// @Tags Authors
// @Produce json
// @Security BearerAuth
// @Param  name query string false "Author Name"
// @Param  biography query string false "Author Biography"
// @Success 200 {object} authors.GetAuthorsResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/authors/ [get]
func (h *HandlerStruct) GetAuthorsHandler(c *gin.Context) {
	name := c.Query("name")
	biography := c.Query("biography")
	req := &authors.GetAuthorsRequest{Name: name, Biography: biography}
	resp, err := h.Author.GetAuthors(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get author by ID
// @Description Retrieve details of an author by their ID
// @Tags Authors
// @Produce json
// @Security BearerAuth
// @Param id path string true "Author ID"
// @Success 200 {object} authors.GetAuthorByIdResponse
// @Failure 400 {object} string "Invalid Author ID"
// @Failure 404 {object} string "Author Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/authors/{id} [get]
func (h *HandlerStruct) GetAuthorByIdHandler(c *gin.Context) {
	id := c.Param("id")
	req := &authors.GetAuthorByIdRequest{Id: id}
	resp, err := h.Author.GetAuthorById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Create a new author
// @Description Add a new author to the library system
// @Tags Authors
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param author body authors.CreateAuthorRequest true "Author Data"
// @Success 201 {object} authors.CreateAuthorResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/authors/ [post]
func (h *HandlerStruct) CreateAuthorHandler(c *gin.Context) {
	var req authors.CreateAuthorRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Author.CreateAuthor(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Update author by ID
// @Description Update details of an existing author by their ID
// @Tags Authors
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Author ID"
// @Param author body authors.UpdateAuthorRequest true "Author Data"
// @Success 200 {object} authors.UpdateAuthorResponse
// @Failure 400 {object} string "Invalid Author ID"
// @Failure 404 {object} string "Author Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/authors/{id} [put]
func (h *HandlerStruct) UpdateAuthorHandler(c *gin.Context) {
	id := c.Param("id")
	var req authors.UpdateAuthorRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = id
	resp, err := h.Author.UpdateAuthor(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete author by ID
// @Description Remove an author from the library system by their ID
// @Tags Authors
// @Produce json
// @Security BearerAuth
// @Param id path string true "Author ID"
// @Success 200 {object} string "Author deleted successfully"
// @Failure 400 {object} string "Invalid Author ID"
// @Failure 404 {object} string "Author Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/authors/{id} [delete]
func (h *HandlerStruct) DeleteAuthorHandler(c *gin.Context) {
	id := c.Param("id")
	req := &authors.DeleteAuthorRequest{Id: id}
	_, err := h.Author.DeleteAuthor(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
