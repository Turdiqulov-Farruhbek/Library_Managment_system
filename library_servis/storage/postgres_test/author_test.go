package postgrestest

import (
	"context"
	"fmt"
	"library/genproto/authors"
	"library/storage/postgres"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

// Helper function to initialize Authors storage
func newTestAuthors(t *testing.T) *postgres.Authors {
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
	return &postgres.Authors{Db: db}
}

// Helper function to create a test author
func createTestAuthor() *authors.CreateAuthorRequest {
	req := &authors.CreateAuthorRequest{
		Name:      "Test Author",
		Biography: "This is a test biography.",
	}
	return req
}

// Test creating an author
func TestCreateAuthor(t *testing.T) {
	authorDB := newTestAuthors(t)
	testAuthor := createTestAuthor()

	authorRes, err := authorDB.CreateAuthor(context.Background(), testAuthor)
	if err != nil {
		t.Fatalf("Failed to create author: %v", err)
	}

	assert.NotEmpty(t, authorRes.Author.Id)
	assert.Equal(t, testAuthor.Name, authorRes.Author.Name)
	assert.Equal(t, testAuthor.Biography, authorRes.Author.Biography)
}

// Test retrieving an author by ID
func TestGetAuthorById(t *testing.T) {
	authorDB := newTestAuthors(t)
	testAuthor := createTestAuthor()

	authorRes, err := authorDB.CreateAuthor(context.Background(), testAuthor)
	if err != nil {
		t.Fatalf("Failed to create author: %v", err)
	}

	getAuthorRes, err := authorDB.GetAuthorById(context.Background(), &authors.GetAuthorByIdRequest{Id: authorRes.Author.Id})
	if err != nil {
		t.Fatalf("Failed to get author by ID: %v", err)
	}

	assert.NotEmpty(t, getAuthorRes.Author.Id)
	assert.Equal(t, testAuthor.Name, getAuthorRes.Author.Name)
	assert.Equal(t, testAuthor.Biography, getAuthorRes.Author.Biography)
}

// Test updating an author
func TestUpdateAuthor(t *testing.T) {
	authorDB := newTestAuthors(t)
	testAuthor := createTestAuthor()

	authorRes, err := authorDB.CreateAuthor(context.Background(), testAuthor)
	if err != nil {
		t.Fatalf("Failed to create author: %v", err)
	}

	updateReq := &authors.UpdateAuthorRequest{
		Id:        authorRes.Author.Id,
		Name:      "Updated Test Author",
		Biography: "This is an updated test biography.",
	}

	updateRes, err := authorDB.UpdateAuthor(context.Background(), updateReq)
	if err != nil {
		t.Fatalf("Failed to update author: %v", err)
	}

	assert.Equal(t, updateReq.Name, updateRes.Author.Name)
	assert.Equal(t, updateReq.Biography, updateRes.Author.Biography)
}

// Test deleting an author
func TestDeleteAuthor(t *testing.T) {
	authorDB := newTestAuthors(t)
	testAuthor := createTestAuthor()

	authorRes, err := authorDB.CreateAuthor(context.Background(), testAuthor)
	if err != nil {
		t.Fatalf("Failed to create author: %v", err)
	}

	_, err = authorDB.DeleteAuthor(context.Background(), &authors.DeleteAuthorRequest{Id: authorRes.Author.Id})
	if err != nil {
		t.Fatalf("Failed to delete author: %v", err)
	}

	deletedAuthor, err := authorDB.GetAuthorById(context.Background(), &authors.GetAuthorByIdRequest{Id: authorRes.Author.Id})
	assert.Nil(t, deletedAuthor)
	assert.Error(t, err) // Check if an error occurred when trying to get the deleted author
}

// Test retrieving all authors
func TestGetAllAuthors(t *testing.T) {
	authorDB := newTestAuthors(t)

	testAuthors := []*authors.CreateAuthorRequest{
		{Name: "Author 1", Biography: "Biography 1"},
		{Name: "Author 2", Biography: "Biography 2"},
		{Name: "Author 3", Biography: "Biography 3"},
	}

	for _, auth := range testAuthors {
		_, err := authorDB.CreateAuthor(context.Background(), auth)
		if err != nil {
			t.Fatalf("Failed to create author: %v", err)
		}
	}

	t.Run("GetAllAuthors without filters", func(t *testing.T) {
		res, err := authorDB.GetAllAuthors(context.Background(), &authors.GetAuthorsRequest{})
		if err != nil {
			t.Fatalf("Failed to get all authors: %v", err)
		}

		assert.LessOrEqual(t, len(testAuthors), len(res.Authors))
	})

	t.Run("Filter by name", func(t *testing.T) {
		res, err := authorDB.GetAllAuthors(context.Background(), &authors.GetAuthorsRequest{Name: "Author 1"})
		if err != nil {
			t.Fatalf("Failed to filter by name: %v", err)
		}

		assert.Equal(t, len(res.Authors), len(res.Authors))
		assert.Equal(t, "Author 1", res.Authors[0].Name)
	})

	t.Run("Filter by biography", func(t *testing.T) {
		res, err := authorDB.GetAllAuthors(context.Background(), &authors.GetAuthorsRequest{Biography: "Biography 2"})
		if err != nil {
			t.Fatalf("Failed to filter by biography: %v", err)
		}

		assert.Equal(t, len(res.Authors), len(res.Authors))
		assert.Equal(t, "Biography 2", res.Authors[0].Biography)
	})
}
