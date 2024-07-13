package postgres

import (
	"context"
	"fmt"
	"library/genproto/borrowers"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Borrowers struct {
	Db *pgx.Conn
}

func NewBorrowers(db *pgx.Conn) *Borrowers {
	return &Borrowers{
		Db: db,
	}
}

var (
	borrowDate, returnDate, createdAt, updatedAt time.Time
)

// CreateBorrower creates a new borrower record in the database.
func (b *Borrowers) CreateBorrower(ctx context.Context, req *borrowers.CreateBorrowerRequest) (*borrowers.CreateBorrowerResponse, error) {
	// Check if the user exists
	var exists bool
	err := b.Db.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`, req.UserId).Scan(&exists)
	if err != nil {
		log.Error().Err(err).Msg("Error checking user existence")
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("user with ID %s does not exist", req.UserId)
	}

	// Parse input dates
	borrowDate, err := time.Parse("2006-01-02", req.BorrowDate)
	if err != nil {
		log.Error().Err(err).Msg("Invalid borrow date format")
		return nil, fmt.Errorf("invalid borrow date format: %s", req.BorrowDate)
	}
	returnDate, err := time.Parse("2006-01-02", req.ReturnDate)
	if err != nil {
		log.Error().Err(err).Msg("Invalid return date format")
		return nil, fmt.Errorf("invalid return date format: %s", req.ReturnDate)
	}

	// Proceed with creating the borrower
	id := uuid.NewString()
	query := `
	INSERT INTO borrowers (id, user_id, book_id, borrow_date, return_date) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, user_id, book_id, borrow_date, return_date, created_at, updated_at`

	var borrower borrowers.Borrower
	var createdAt, updatedAt time.Time

	err = b.Db.QueryRow(ctx, query, id, req.UserId, req.BookId, borrowDate, returnDate).
		Scan(&borrower.Id, &borrower.UserId, &borrower.BookId, &borrowDate, &returnDate, &createdAt, &updatedAt)

	if err != nil {
		log.Error().Err(err).Msg("Error creating borrower")
		return nil, err
	}

	// Convert time.Time to string
	borrower.CreatedAt = createdAt.Format(time.RFC3339)
	borrower.UpdatedAt = updatedAt.Format(time.RFC3339)
	borrower.BorrowDate = borrowDate.Format("2006-01-02")
	borrower.ReturnDate = returnDate.Format("2006-01-02")

	return &borrowers.CreateBorrowerResponse{Borrower: &borrower}, nil
}

// UpdateBorrower updates an existing borrower record in the database.
func (b *Borrowers) UpdateBorrower(ctx context.Context, req *borrowers.UpdateBorrowerRequest) (*borrowers.UpdateBorrowerResponse, error) {
	// Validate UUIDs
	_, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid UserId format: %v", err)
	}

	_, err = uuid.Parse(req.BookId)
	if err != nil {
		return nil, fmt.Errorf("invalid BookId format: %v", err)
	}

	query := `
		UPDATE borrowers
		SET user_id = $1, book_id = $2, borrow_date = $3, return_date = $4, updated_at = CURRENT_TIMESTAMP
		WHERE id = $5 AND deleted_at = 0
		RETURNING id, user_id, book_id, borrow_date, return_date, created_at, updated_at
	`

	var (
		borrower   borrowers.Borrower
		returnDate pq.NullTime
	)

	err = b.Db.QueryRow(ctx, query, req.UserId, req.BookId, req.BorrowDate, req.ReturnDate, req.Id).
		Scan(&borrower.Id, &borrower.UserId, &borrower.BookId, &borrowDate, &returnDate, &createdAt, &updatedAt)

	if err != nil {
		log.Error().Err(err).Msg("Error updating borrower")
		return nil, err
	}

	borrower.BorrowDate = borrowDate.Format("2006-01-02")
	if returnDate.Valid {
		borrower.ReturnDate = returnDate.Time.Format("2006-01-02")
	}
	borrower.CreatedAt = createdAt.Format(time.RFC3339)
	borrower.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &borrowers.UpdateBorrowerResponse{Borrower: &borrower}, nil
}



// DeleteBorrower marks a borrower record as deleted in the database.
func (b *Borrowers) DeleteBorrower(ctx context.Context, req *borrowers.DeleteBorrowerRequest) (*borrowers.DeleteBorrowerResponse, error) {
	query := `UPDATE borrowers SET deleted_at = EXTRACT(EPOCH FROM NOW()) 
		WHERE id = $1 AND deleted_at = 0`
	_, err := b.Db.Exec(ctx, query, req.Id)

	if err != nil {
		log.Error().Err(err).Msg("Error deleting borrower")
		return nil, err
	}

	return &borrowers.DeleteBorrowerResponse{Message: "Borrower record deleted successfully"}, nil
}

// GetBorrowerById fetches a borrower record by its ID.
func (b *Borrowers) GetBorrowerById(ctx context.Context, req *borrowers.GetBorrowerByIdRequest) (*borrowers.GetBorrowerByIdResponse, error) {
	query := `
	SELECT id, user_id, book_id, borrow_date, return_date, created_at, updated_at
	FROM borrowers
	WHERE id = $1 AND deleted_at = 0`

	var (
		borrower   borrowers.Borrower
		returnDate pq.NullTime
	)

	err := b.Db.QueryRow(ctx, query, req.Id).
		Scan(&borrower.Id, &borrower.UserId, &borrower.BookId, &borrowDate, &returnDate, &createdAt, &updatedAt)

	if err != nil {
		log.Error().Err(err).Msg("Error getting borrower by ID")
		return nil, err
	}

	borrower.BorrowDate = borrowDate.Format("2006-01-02")
	if returnDate.Valid {
		borrower.ReturnDate = returnDate.Time.Format("2006-01-02")
	}
	borrower.CreatedAt = createdAt.Format(time.RFC3339)
	borrower.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &borrowers.GetBorrowerByIdResponse{Borrower: &borrower}, nil
}

// GetAllBorrowers fetches all borrower records from the database.
func (b *Borrowers) GetAllBorrowers(ctx context.Context, req *borrowers.GetAllBorrowersRequest) (*borrowers.GetAllBorrowersResponse, error) {
	var args []interface{}
	count := 1

	query := `SELECT id, user_id, book_id, borrow_date, return_date, created_at, updated_at 
		FROM borrowers 
		WHERE deleted_at = 0`

	if req.UserId != "" {
		query += fmt.Sprintf(" AND user_id = $%d", count)
		args = append(args, req.UserId)
		count++
	}

	if req.BookId != "" {
		query += fmt.Sprintf(" AND book_id = $%d", count)
		args = append(args, req.BookId)
		count++
	}

	if req.BorrowDate != "" {
		query += fmt.Sprintf(" AND borrow_date = $%d", count)
		args = append(args, req.BorrowDate)
		count++
	}

	if req.ReturnDate != "" {
		query += fmt.Sprintf(" AND return_date = $%d", count)
		args = append(args, req.ReturnDate)
		count++
	}

	rows, err := b.Db.Query(ctx, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching all borrowers")
		return nil, err
	}
	defer rows.Close()

	var borrowersList []*borrowers.Borrower
	for rows.Next() {
		var borrower borrowers.Borrower

		err := rows.Scan(&borrower.Id, &borrower.UserId, &borrower.BookId, &borrowDate, &returnDate, &createdAt, &updatedAt)
		if err != nil {
			log.Error().Err(err).Msg("Error fetching borrower")
			return nil, err
		}

		// Convert time.Time to string
		borrower.BorrowDate = borrowDate.Format("2006-01-02")
		borrower.ReturnDate = returnDate.Format("2006-01-02")
		borrower.CreatedAt = createdAt.Format(time.RFC3339)
		borrower.UpdatedAt = updatedAt.Format(time.RFC3339)

		borrowersList = append(borrowersList, &borrower)
	}
	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over borrowers")
		return nil, err
	}

	return &borrowers.GetAllBorrowersResponse{Borrowers: borrowersList}, nil
}

func (b *Borrowers) GetBorrowingHistory(ctx context.Context, req *borrowers.GetBorrowingHistoryRequest) (*borrowers.GetBorrowingHistoryResponse, error) {
	query := `
	SELECT id, user_id, book_id, borrow_date, return_date, created_at, updated_at
	FROM borrowers
	WHERE user_id = $1 AND deleted_at = 0`

	rows, err := b.Db.Query(ctx, query, req.UserId)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching borrowing history from the database")
		return nil, err
	}
	defer rows.Close()

	var currentlyBorrowed []*borrowers.Borrower
	var history []*borrowers.Borrower

	for rows.Next() {
		var borrower borrowers.Borrower

		err := rows.Scan(&borrower.Id, &borrower.UserId, &borrower.BookId, &borrowDate, &returnDate, &createdAt, &updatedAt)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning borrowing history row")
			return nil, err
		}

		borrower.BorrowDate = borrowDate.Format("2006-01-02")
		if !returnDate.IsZero() {
			borrower.ReturnDate = returnDate.Format("2006-01-02")
		}
		borrower.CreatedAt = createdAt.Format(time.RFC3339)
		borrower.UpdatedAt = updatedAt.Format(time.RFC3339)

		if returnDate.IsZero() {
			currentlyBorrowed = append(currentlyBorrowed, &borrower)
		} else {
			history = append(history, &borrower)
		}
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over borrowing history rows")
		return nil, err
	}

	return &borrowers.GetBorrowingHistoryResponse{
		CurrentlyBorrowed: currentlyBorrowed,
		History:           history,
	}, nil
}

func (b *Borrowers) GetBorrowedBooksByUser(ctx context.Context, req *borrowers.GetBorrowedBooksByUserRequest) (*borrowers.GetBorrowedBooksByUserResponse, error) {
	query := `
	SELECT b.id, b.title, b.author_id, b.genre_id, b.summary, b.created_at, b.updated_at
	FROM borrowers br
	INNER JOIN books b ON br.book_id = b.id
	WHERE br.user_id = $1 AND br.return_date = 0`

	rows, err := b.Db.Query(ctx, query, req.UserId)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching borrowed books from the database")
		return nil, err
	}
	defer rows.Close()

	var borrowedBooks []*borrowers.BorrowedBook
	for rows.Next() {
		var book borrowers.BorrowedBook
		err := rows.Scan(&book.Title, &book.AuthorId, &book.GenreId, &book.Summary, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning borrowed books row")
			return nil, err
		}

		book.CreatedAt = createdAt.Format(time.RFC3339)
		book.UpdatedAt = updatedAt.Format(time.RFC3339)
		borrowedBooks = append(borrowedBooks, &book)
	}

	return &borrowers.GetBorrowedBooksByUserResponse{Books: borrowedBooks}, nil
}
