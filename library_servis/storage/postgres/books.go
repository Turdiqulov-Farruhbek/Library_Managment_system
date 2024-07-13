package postgres

import (
	"context"
	"fmt"
	book "library/genproto/books"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

type Books struct {
	Db *pgx.Conn
}

func NewBooks(db *pgx.Conn) *Books {
	return &Books{
		Db: db,
	}
}

func (b *Books) CreateBook(ctx context.Context, req *book.CreateBookRequest) (*book.CreateBookResponse, error) {
	// Generate a new UUID for the book
	id := uuid.NewString()

	// SQL query to insert a new book and return its details
	query := `
	INSERT INTO books (id, title, author_id, genre_id, summary) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, title, author_id, genre_id, summary, created_at, updated_at`

	// Prepare the book structure to hold the response data
	bk := &book.Book{}
	var createdAt, updatedAt time.Time

	// Execute the query and scan the returned values into the book structure
	err := b.Db.QueryRow(ctx, query, id, req.Title, req.AuthorId, req.GenreId, req.Summary).
		Scan(&bk.Id, &bk.Title, &bk.AuthorId, &bk.GenreId, &bk.Summary, &createdAt, &updatedAt)
	if err != nil {
		// Log and return a detailed error message in case of failure
		log.Error().Err(err).Msg("Error creating book")
		return nil, status.Error(codes.Internal, "Error creating book")
	}

	// Format timestamps to RFC3339 for consistency
	bk.CreatedAt = createdAt.Format(time.RFC3339)
	bk.UpdatedAt = updatedAt.Format(time.RFC3339)

	// Return the created book in the response
	return &book.CreateBookResponse{Book: bk}, nil
}


func (b *Books) UpdateBook(ctx context.Context, req *book.UpdateBookRequest) (*book.UpdateBookResponse, error) {
	query := `
	UPDATE books
	SET title = $1, author_id = $2, genre_id = $3, summary = $4, updated_at = NOW()
	WHERE id = $5 AND deleted_at = 0
	RETURNING id, title, author_id, genre_id, summary, created_at, updated_at`

	bk := book.Book{}
	var createdAt, updatedAt time.Time
	err := b.Db.QueryRow(ctx, query, req.Title, req.AuthorId, req.GenreId, req.Summary, req.Id).
		Scan(&bk.Id, &bk.Title, &bk.AuthorId, &bk.GenreId, &bk.Summary, &createdAt, &updatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Warn().Msgf("No book found with ID %s or book is deleted", req.Id)
			return nil, status.Error(codes.NotFound, fmt.Sprintf("Book with ID %s not found or deleted", req.Id))
		}
		log.Error().Err(err).Msg("Error updating book")
		return nil, status.Error(codes.Internal, "Error updating book")
	}
	bk.CreatedAt = createdAt.Format(time.RFC3339)
	bk.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &book.UpdateBookResponse{Book: &bk}, nil
}

func (b *Books) DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (*book.DeleteBookResponse, error) {
	query := `UPDATE books SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0`
	result, err := b.Db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting book")
		return nil, status.Error(codes.Internal, "Error deleting book")
	}
	if result.RowsAffected() == 0 {
		log.Warn().Msgf("No book found with ID %s or book is already deleted", req.Id)
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Book with ID %s not found or already deleted", req.Id))
	}

	return &book.DeleteBookResponse{Message: "Book deleted successfully"}, nil
}

func (b *Books) GetBookById(ctx context.Context, req *book.GetBookByIdRequest) (*book.GetBookByIdResponse, error) {
	query := `SELECT id, title, author_id, genre_id, summary, created_at, updated_at FROM books WHERE id = $1 AND deleted_at = 0`
	bk := book.Book{}
	var createdAt, updatedAt time.Time
	err := b.Db.QueryRow(ctx, query, req.Id).
		Scan(&bk.Id, &bk.Title, &bk.AuthorId, &bk.GenreId, &bk.Summary, &createdAt, &updatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Warn().Msgf("No book found with ID %s", req.Id)
			return nil, status.Error(codes.NotFound, fmt.Sprintf("Book with ID %s not found", req.Id))
		}
		log.Error().Err(err).Msg("Error getting book by id")
		return nil, status.Error(codes.Internal, "Error getting book")
	}
	bk.CreatedAt = createdAt.Format(time.RFC3339)
	bk.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &book.GetBookByIdResponse{Book: &bk}, nil
}

func (b *Books) GetallBooks(ctx context.Context, req *book.GetBooksRequest) (*book.GetBooksResponse, error) {
	var args []interface{}
	count := 1
	query := `SELECT id, title, author_id, genre_id, summary, created_at, updated_at FROM books WHERE deleted_at = 0`

	filter := ""
	if req.Title != "" {
		filter += fmt.Sprintf(" AND title ILIKE $%d", count)
		args = append(args, "%"+req.Title+"%")
		count++
	}
	if req.AuthorId != "" {
		filter += fmt.Sprintf(" AND author_id = $%d", count)
		args = append(args, req.AuthorId)
		count++
	}
	if req.GenreId != "" {
		filter += fmt.Sprintf(" AND genre_id = $%d", count)
		args = append(args, req.GenreId)
		count++
	}
	query += filter

	rows, err := b.Db.Query(ctx, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching books from the database")
		return nil, status.Error(codes.Internal, "Error fetching books")
	}
	defer rows.Close()

	var books []*book.Book
	for rows.Next() {
		bk := book.Book{}
		var createdAt, updatedAt time.Time
		err := rows.Scan(&bk.Id, &bk.Title, &bk.AuthorId, &bk.GenreId, &bk.Summary, &createdAt, &updatedAt)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning book row")
			return nil, status.Error(codes.Internal, "Error scanning book row")
		}
		bk.CreatedAt = createdAt.Format(time.RFC3339)
		bk.UpdatedAt = updatedAt.Format(time.RFC3339)
		books = append(books, &bk)
	}

	return &book.GetBooksResponse{Books: books}, nil
}

// GetBooksByAuthorId retrieves books by a specific author from the database.
func (b *Books) GetBooksByAuthorId(ctx context.Context, req *book.GetBooksByAuthorIdRequest) (*book.GetBooksResponse, error) {
	query := `SELECT id, title, author_id, genre_id, summary, created_at, updated_at FROM books WHERE author_id = $1 AND deleted_at = 0`
	rows, err := b.Db.Query(ctx, query, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching books by author ID")
		return nil, status.Error(codes.Internal, "Error fetching books by author ID")
	}
	defer rows.Close()

	var books []*book.Book
	for rows.Next() {
		bk := book.Book{}
		var createdAt, updatedAt time.Time
		err := rows.Scan(&bk.Id, &bk.Title, &bk.AuthorId, &bk.GenreId, &bk.Summary, &createdAt, &updatedAt)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning book row by author ID")
			return nil, status.Error(codes.Internal, "Error scanning book row by author ID")
		}
		bk.CreatedAt = createdAt.Format(time.RFC3339)
		bk.UpdatedAt = updatedAt.Format(time.RFC3339)
		books = append(books, &bk)
	}

	return &book.GetBooksResponse{Books: books}, nil
}

// GetBooksByGenreId retrieves books by a specific genre from the database.
func (b *Books) GetBooksByGenreId(ctx context.Context, req *book.GetBooksByGenreIdRequest) (*book.GetBooksResponse, error) {
	query := `
	SELECT id, title, author_id, genre_id, summary, created_at, updated_at
	FROM books
	WHERE genre_id = $1 AND deleted_at = 0`

	rows, err := b.Db.Query(ctx, query, req.Id) // req.GenreId olarak güncellendi
	if err != nil {
		log.Error().Err(err).Msg("Error fetching books by genre ID")
		return nil, status.Error(codes.Internal, "Error fetching books by genre ID")
	}
	defer rows.Close()

	var books []*book.Book
	for rows.Next() {
		bk := book.Book{}
		var createdAt, updatedAt time.Time
		err := rows.Scan(
			&bk.Id,
			&bk.Title,
			&bk.AuthorId,
			&bk.GenreId,
			&bk.Summary,
			&createdAt, // time.Time olarak düzeltildi
			&updatedAt, // time.Time olarak düzeltildi
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning book row by genre ID")
			return nil, status.Error(codes.Internal, "Error scanning book row by genre ID")
		}

		bk.CreatedAt = createdAt.Format(time.RFC3339)
		bk.UpdatedAt = updatedAt.Format(time.RFC3339)
		books = append(books, &bk)
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Row iteration error while fetching books by genre ID")
		return nil, status.Error(codes.Internal, "Error fetching books by genre ID")
	}

	return &book.GetBooksResponse{Books: books}, nil
}

// GetOverdueBooks retrieves books that are overdue from the database.

func (b *Books) GetOverdueBooks(ctx context.Context, req *book.GetOverdueBooksRequest) (*book.GetBooksResponse, error) {
	// SQL sorgusunu hazırlama
	query := `
		SELECT b.id, b.title, br.user_id, b.author_id, b.genre_id, b.summary, b.created_at, b.updated_at
		FROM books b
		INNER JOIN borrowers br ON b.id = br.book_id
		WHERE br.return_date > CURRENT_DATE AND br.deleted_at = 0 AND b.deleted_at = 0 AND br.user_id IS NOT NULL
	`
	// Sorguyu çalıştırma
	rows, err := b.Db.Query(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching overdue books")
		return nil, status.Error(codes.Internal, "Error fetching overdue books")
	}
	defer rows.Close()
	// Kitapları toplama
	var books []*book.Book
	for rows.Next() {
		bk := &book.Book{}
		var createdAt, updatedAt time.Time
		// Satırdan verileri okuma
		err := rows.Scan(
			&bk.Id,
			&bk.Title,
			&bk.UserId,
			&bk.AuthorId,
			&bk.GenreId,
			&bk.Summary,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning overdue book row")
			return nil, status.Error(codes.Internal, "Error scanning overdue book row")
		}
		// Zaman bilgilerini RFC3339 formatına dönüştürme
		bk.CreatedAt = createdAt.Format(time.RFC3339)
		bk.UpdatedAt = updatedAt.Format(time.RFC3339)
		books = append(books, bk)
	}
	// Satırların sonunda oluşan hata kontrolü
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Row iteration error while fetching overdue books")
		return nil, status.Error(codes.Internal, "Error fetching overdue books")
	}

	return &book.GetBooksResponse{Books: books}, nil
}
