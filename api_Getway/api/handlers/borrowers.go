package handlers

import (
	"context"
	"getway/genproto/borrowers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all borrowers
// @Description Retrieve a list of all borrowers
// @Tags Borrowers
// @Produce json
// @Security BearerAuth
// @Param query query borrowers.GetAllBorrowersRequest true "Book Querys"
// @Success 200 {object} borrowers.GetAllBorrowersResponse "Menu item found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/borrowers/ [get]
func (h *HandlerStruct) GetBorrowersHandler(c *gin.Context) {
	userId := c.Param("userId")
	bookId := c.Param("bookId")
	borrowDate := c.Param("borrowDate")
	returnDate := c.Param("returnDate")
	req := &borrowers.GetAllBorrowersRequest{UserId: userId, BookId: bookId,BorrowDate: borrowDate, ReturnDate: returnDate}
	resp, err := h.Borrowers.GetAllBorrowers(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get borrower by ID
// @Description Retrieve details of a borrower by their ID
// @Tags Borrowers
// @Produce json
// @Security BearerAuth
// @Param id path string true "Borrower ID"
// @Success 200 {object} borrowers.GetBorrowerByIdResponse
// @Failure 400 {object} string "Borrower ID"
// @Failure 404 {object} string "Borrower Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/borrowers/{id} [get]
func (h *HandlerStruct) GetBorrowerByIdHandler(c *gin.Context) {
	id := c.Param("id")
	req := &borrowers.GetBorrowerByIdRequest{Id: id}
	resp, err := h.Borrowers.GetBorrowerById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get borrowed books by user ID
// @Description Retrieve a list of books borrowed by a specific user
// @Tags Borrowers
// @Produce json
// @Security BearerAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} borrowers.GetAllBorrowersResponse
// @Failure 400 {object} string "User ID"
// @Failure 404 {object} string "Borrower Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/borrowers/users/{user_id}/borrowed_books [get]
func (h *HandlerStruct) GetBorrowersBookHandler(c *gin.Context) {
	userID := c.Param("user_id")
	req := &borrowers.GetAllBorrowersRequest{UserId: userID}
	resp, err := h.Borrowers.GetAllBorrowers(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get borrowing history by user ID
// @Description Retrieve borrowing history of a specific user
// @Tags Borrowers
// @Produce json
// @Security BearerAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} borrowers.GetBorrowingHistoryResponse
// @Failure 400 {object} string "Invalid User ID"
// @Failure 404 {object} string "Borrowing History Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/borrowers/users/{user_id}/borrowing_history [get]
func (h *HandlerStruct) GetBorrowingHistoryHandler(c *gin.Context) {
	userID := c.Param("user_id")
	req := &borrowers.GetBorrowingHistoryRequest{UserId: userID}
	resp, err := h.Borrowers.GetBorrowingHistory(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Create a new borrower
// @Description Add a new borrower to the library system
// @Tags Borrowers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param borrower body borrowers.CreateBorrowerRequest true "Borrower Data"
// @Success 201 {object} borrowers.CreateBorrowerResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/borrowers/ [post]
func (h *HandlerStruct) CreateBorrowerHandler(c *gin.Context) {
	var req borrowers.CreateBorrowerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Borrowers.CreateBorrower(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Update borrower by ID
// @Description Update details of an existing borrower by their ID
// @Tags Borrowers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Borrower ID"
// @Param borrower body borrowers.UpdateBorrowerRequest true "Borrower Data"
// @Success 200 {object} borrowers.UpdateBorrowerResponse
// @Failure 400 {object} string "Borrower ID"
// @Failure 404 {object} string "Borrower Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/borrowers/{id} [put]
func (h *HandlerStruct) UpdateBorrowerHandler(c *gin.Context) {
	id := c.Param("id")
	var req borrowers.UpdateBorrowerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = id
	resp, err := h.Borrowers.UpdateBorrower(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete borrower by ID
// @Description Remove a borrower from the library system by their ID
// @Tags Borrowers
// @Produce json
// @Security BearerAuth
// @Param id path string true "Borrower ID"
// @Success 200 {object} string "Borrower deleted successfully"
// @Failure 400 {object} string "Borrower ID"
// @Failure 404 {object} string "Borrower Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/borrowers/{id} [delete]
func (h *HandlerStruct) DeleteBorrowerHandler(c *gin.Context) {
	id := c.Param("id")
	req := &borrowers.DeleteBorrowerRequest{Id: id}
	_, err := h.Borrowers.DeleteBorrower(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Borrower deleted successfully"})
}
