package service

import (
	"context"
	"fmt"
	book "library/genproto/books"
	"library/storage"

	"github.com/rs/zerolog/log"
)

// BooksService struct handles book-related operations by interacting with the storage layer.
type BooksService struct {
	stg storage.StorageI
	book.UnimplementedBooksServiceServer
}

// NewBooksService initializes and returns a new BooksService instance.
func NewBooksService(stg storage.StorageI) *BooksService {
	return &BooksService{stg: stg}
}

// CreateService creates a new book record.
func (s *BooksService) CreateBook(ctx context.Context, req *book.CreateBookRequest) (*book.CreateBookResponse, error) {
	log.Info().Msg("BooksService: CreateBook called")

	resp, err := s.stg.Books().CreateBook(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BooksService: Error creating book")
		return nil, err
	}

	return resp, nil
}

// UpdateService updates an existing book record.
func (s *BooksService) UpdateBook(ctx context.Context, req *book.UpdateBookRequest) (*book.UpdateBookResponse, error) {
	log.Info().Msg("BooksService: UpdateBook called")

	resp, err := s.stg.Books().UpdateBook(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BooksService: Error updating book")
		return nil, err
	}

	return resp, nil
}

// DeleteService deletes a book record by marking it as deleted.
func (s *BooksService) DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (*book.DeleteBookResponse, error) {
	log.Info().Msg("BooksService: DeleteBook called")

	resp, err := s.stg.Books().DeleteBook(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BooksService: Error deleting book")
		return nil, err
	}

	return resp, nil
}

// GetBookService retrieves a book by its ID.
func (s *BooksService) GetBookById(ctx context.Context, req *book.GetBookByIdRequest) (*book.GetBookByIdResponse, error) {
	log.Info().Msg("BooksService: GetBookById called")

	resp, err := s.stg.Books().GetBookById(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BooksService: Error getting book by ID")
		return nil, err
	}

	return resp, nil
}

// GetAllBooksService retrieves all books or those matching the filter criteria.
func (s *BooksService) GetBooks(ctx context.Context, req *book.GetBooksRequest) (*book.GetBooksResponse, error) {
	log.Info().Msg("BooksService: GetAllBooks called")

	resp, err := s.stg.Books().GetallBooks(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BooksService: Error getting all books")
		return nil, err
	}

	return resp, nil
}

// GetBooksByAuthorId retrieves books by a specific author.
func (s *BooksService) GetBooksByAuthorId(ctx context.Context, req *book.GetBooksByAuthorIdRequest) (*book.GetBooksResponse, error) {
	log.Info().Msg("Fetching books by author ID")

	resp, err := s.stg.Books().GetBooksByAuthorId(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching books by author ID")
		return nil, err
	}

	return resp, nil
}

// GetBooksByGenreId retrieves books by a specific genre.
func (s *BooksService) GetBooksByGenreId(ctx context.Context, req *book.GetBooksByGenreIdRequest) (*book.GetBooksResponse, error) {
	log.Info().Msg("Fetching books by genre ID")

	resp, err := s.stg.Books().GetBooksByGenreId(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching books by genre ID")
		return nil, err
	}

	return resp, nil
}

// GetOverdueBooks retrieves books that are overdue.
func (s *BooksService) GetOverdueBooks(ctx context.Context, req *book.GetOverdueBooksRequest) (*book.GetBooksResponse, error) {
	log.Info().Msg("Fetching overdue books")
	fmt.Println(0)
	resp, err := s.stg.Books().GetOverdueBooks(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching overdue books")
		return nil, err
	}

	return resp, nil
}
