package handlers

import (
	"context"
	"fmt"
	"getway/genproto/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all books
// @Description Retrieve a list of all books
// @Tags Books
// @Produce json
// @Security BearerAuth
// @Param  title query string false "Book title"
// @Param  author_id query string false "Book author_id"
// @Param  genre_id query string false "Book genre_id"
// @Param  summary query string false "Book summary"
// @Success 200 {object} books.GetBooksResponse
// @Failure 500 {object} string "Internal server error"
// @Router /api/books/ [get]
func (h *HandlerStruct) GetBooksHandler(c *gin.Context) {
	title := c.Query("title")
	author_id := c.Query("author_id")
	genre_id := c.Query("genre_id")
	summary := c.Query("summary")
	req := &books.GetBooksRequest{Title: title, AuthorId: author_id, GenreId: genre_id, Summary: summary}
	resp, err := h.Books.GetBooks(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get book by ID
// @Description Retrieve details of a book by its ID
// @Tags Books
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Success 200 {object} books.GetBookByIdResponse
// @Failure 400 {object} string "Invalid Book ID"
// @Failure 404 {object} string "Book Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/books/{id} [get]
func (h *HandlerStruct) GetBookByIdHandler(c *gin.Context) {
	id := c.Param("id")
	req := &books.GetBookByIdRequest{Id: id}
	resp, err := h.Books.GetBookById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get books by author ID
// @Description Retrieve a list of books by a specific author
// @Tags Books
// @Produce json
// @Security BearerAuth
// @Param author_id path string true "Author ID"
// @Success 200 {object} books.GetBooksResponse
// @Failure 400 {object} string "Invalid Author ID"
// @Failure 404 {object} string "Books Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/books/authors/{author_id}/books [get]
func (h *HandlerStruct) GetBooksAuthorHandler(c *gin.Context) {
	authorID := c.Param("author_id")
	req := &books.GetBooksByAuthorIdRequest{Id: authorID}
	resp, err := h.Books.GetBooksByAuthorId(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get books by genre ID
// @Description Retrieve a list of books by a specific genre
// @Tags Books
// @Produce json
// @Security BearerAuth
// @Param genre_id path string true "Genre ID"
// @Success 200 {object} books.GetBooksResponse
// @Failure 400 {object} string "Invalid Genre ID"
// @Failure 404 {object} string "Books Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/books/genres/{genre_id}/books [get]
func (h *HandlerStruct) GetBooksGenreHandler(c *gin.Context) {
    genreID := c.Param("genre_id")
    req := &books.GetBooksByGenreIdRequest{Id: genreID}
	fmt.Println("handler:", req)
    resp, err := h.Books.GetBooksByGenreId(context.Background(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}


// @Summary Get overdue books
// @Description Retrieve a list of overdue books
// @Tags Books
// @Produce json
// @Security BearerAuth
// @Success 200 {object} books.GetBooksResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/books/overdue [get]
func (h *HandlerStruct) GetBooksOverdueHandler(c *gin.Context) {
	req := &books.GetOverdueBooksRequest{}
	resp, err := h.Books.GetOverdueBooks(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Create a new book
// @Description Add a new book to the library system
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param book body books.CreateBookRequest true "Book Data"
// @Success 201 {object} books.CreateBookResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/books/ [post]
func (h *HandlerStruct) CreateBookHandler(c *gin.Context) {
	var req books.CreateBookRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Books.CreateBook(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Update book by ID
// @Description Update details of an existing book by its ID
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Param book body books.UpdateBookRequest true "Book Data"
// @Success 200 {object} books.UpdateBookResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 404 {object} string "Book Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/books/{id} [put]
func (h *HandlerStruct) UpdateBookHandler(c *gin.Context) {
	id := c.Param("id")
	var req books.UpdateBookRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = id
	resp, err := h.Books.UpdateBook(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete book by ID
// @Description Remove a book from the library system by its ID
// @Tags Books
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Success 200 {object} string "Book deleted successfully"
// @Failure 400 {object} string "Invalid Book ID"
// @Failure 404 {object} string "Book Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/books/{id} [delete]
func (h *HandlerStruct) DeleteBookHandler(c *gin.Context) {
	id := c.Param("id")
	req := &books.DeleteBookRequest{Id: id}
	_, err := h.Books.DeleteBook(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
