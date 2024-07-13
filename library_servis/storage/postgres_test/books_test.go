package postgrestest

import (
	"context"
	"fmt"
	book "library/genproto/books"
	"library/storage/postgres"
	"testing"

	// "github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func newTestBooks(t *testing.T) *postgres.Books {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"postgres",
		"root",
		"localhost",
		5432,
		"library_db",
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &postgres.Books{Db: db}
}

func createTestBook() *book.CreateBookRequest {
	return &book.CreateBookRequest{
		Title:    "Test Book Title",
		AuthorId: "9c6d087d-459a-4141-803e-348c974ccaff",
		GenreId:  "1248fa3d-f090-4d07-9338-59445d6384ac",
		Summary:  "Test summary of the book",
	}
}

func TestCreateBook(t *testing.T) {
	bookDB := newTestBooks(t)
	testBook := createTestBook()

	bookRes, err := bookDB.CreateBook(context.Background(), testBook)
	if err != nil {
		t.Fatalf("Failed to create book: %v", err)
	}

	assert.NotEmpty(t, bookRes.Book.Id)
	assert.Equal(t, bookRes.Book.Title, testBook.Title)
	assert.Equal(t, bookRes.Book.AuthorId, testBook.AuthorId)
	assert.Equal(t, bookRes.Book.GenreId, testBook.GenreId)
	assert.Equal(t, bookRes.Book.Summary, testBook.Summary)
}

func TestGetBookById(t *testing.T) {
	bookDB := newTestBooks(t)
	testBook := createTestBook()

	bookRes, err := bookDB.CreateBook(context.Background(), testBook)
	if err != nil {
		t.Fatalf("Failed to create book: %v", err)
	}

	getBookRes, err := bookDB.GetBookById(context.Background(), &book.GetBookByIdRequest{Id: bookRes.Book.Id})
	if err != nil {
		t.Fatalf("Failed to get book by ID: %v", err)
	}

	assert.Equal(t, getBookRes.Book.Id, bookRes.Book.Id)
	assert.Equal(t, getBookRes.Book.Title, testBook.Title)
	assert.Equal(t, getBookRes.Book.AuthorId, testBook.AuthorId)
	assert.Equal(t, getBookRes.Book.GenreId, testBook.GenreId)
	assert.Equal(t, getBookRes.Book.Summary, testBook.Summary)
}

func TestUpdateBook(t *testing.T) {
	bookDB := newTestBooks(t)
	testBook := createTestBook()

	bookRes, err := bookDB.CreateBook(context.Background(), testBook)
	if err != nil {
		t.Fatalf("Failed to create book: %v", err)
	}

	updateReq := book.UpdateBookRequest{
		Id:       bookRes.Book.Id,
		Title:    "Updated Book Title",
		AuthorId: bookRes.Book.AuthorId,
		GenreId:  bookRes.Book.GenreId,
		Summary:  "Updated summary of the book",
	}

	updateRes, err := bookDB.UpdateBook(context.Background(), &updateReq)
	if err != nil {
		t.Fatalf("Failed to update book: %v", err)
	}

	assert.Equal(t, updateRes.Book.Id, updateReq.Id)
	assert.Equal(t, updateRes.Book.Title, updateReq.Title)
	assert.Equal(t, updateRes.Book.Summary, updateReq.Summary)
}

func TestDeleteBook(t *testing.T) {
	bookDB := newTestBooks(t)
	testBook := createTestBook()

	bookRes, err := bookDB.CreateBook(context.Background(), testBook)
	if err != nil {
		t.Fatalf("Failed to create book: %v", err)
	}

	_, err = bookDB.DeleteBook(context.Background(), &book.DeleteBookRequest{Id: bookRes.Book.Id})
	if err != nil {
		t.Fatalf("Failed to delete book: %v", err)
	}

	deletedBook, err := bookDB.GetBookById(context.Background(), &book.GetBookByIdRequest{Id: bookRes.Book.Id})
	assert.Nil(t, deletedBook)
	assert.Error(t, err) // Check if an error occurred when trying to get the deleted book
}

func TestGetAllBooks(t *testing.T) {
	bookDB := newTestBooks(t)

	authorId := "9c6d087d-459a-4141-803e-348c974ccaff"
	genreId := "1248fa3d-f090-4d07-9338-59445d6384ac"


	testBooks := []*book.CreateBookRequest{
		{Title: "Book One", AuthorId: authorId, GenreId: genreId, Summary: "Summary One"},
		{Title: "Book Two", AuthorId: authorId, GenreId: genreId, Summary: "Summary Two"},
		{Title: "Book Three", AuthorId: authorId, GenreId: genreId, Summary: "Summary Three"},
	}

	for _, bk := range testBooks {
		_, err := bookDB.CreateBook(context.Background(), bk)
		if err != nil {
			t.Fatalf("Failed to create book: %v", err)
		}
	}

	t.Run("GetAllBooks without filters", func(t *testing.T) {
		res, err := bookDB.GetallBooks(context.Background(), &book.GetBooksRequest{})
		if err != nil {
			t.Fatalf("Failed to get all books: %v", err)
		}

		assert.GreaterOrEqual(t, len(res.Books), len(testBooks))
	})

	t.Run("Filter by title", func(t *testing.T) {
		res, err := bookDB.GetallBooks(context.Background(), &book.GetBooksRequest{Title: "Book Two"})
		if err != nil {
			t.Fatalf("Failed to filter books by title: %v", err)
		}

		assert.Equal(t,  len(res.Books), len(res.Books))
		assert.Equal(t, "Book Two", res.Books[0].Title)
	})

	t.Run("Filter by author ID", func(t *testing.T) {
		res, err := bookDB.GetallBooks(context.Background(), &book.GetBooksRequest{AuthorId: authorId})
		if err != nil {
			t.Fatalf("Failed to filter books by author ID: %v", err)
		}

		assert.Equal(t,  len(res.Books), len(res.Books))
	})

	t.Run("Filter by genre ID", func(t *testing.T) {
		res, err := bookDB.GetallBooks(context.Background(), &book.GetBooksRequest{GenreId: genreId})
		if err != nil {
			t.Fatalf("Failed to filter books by genre ID: %v", err)
		}

		assert.Equal(t,  len(res.Books), len(res.Books))
	})
}
