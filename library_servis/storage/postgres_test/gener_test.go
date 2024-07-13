package postgrestest

import (
	"context"
	"fmt"
	genre "library/genproto/genres"
	"library/storage/postgres"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

// Function to establish a new connection to the test database
func newTestGenres(t *testing.T) *postgres.Genres {
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
	return &postgres.Genres{Db: db}
}

// Helper function to create a test genre
func createTestGenre() *genre.CreateGenreRequest {
	return &genre.CreateGenreRequest{
		Name: "Test Genre",
	}
}

func TestCreateGenre(t *testing.T) {
	genreDB := newTestGenres(t)
	testGenre := createTestGenre()

	genreRes, err := genreDB.CreateGenre(context.Background(), testGenre)
	if err != nil {
		t.Fatalf("Failed to create genre: %v", err)
	}

	assert.NotEmpty(t, genreRes.Genre.Id)
	assert.Equal(t, genreRes.Genre.Name, testGenre.Name)
}

func TestGetGenreById(t *testing.T) {
	genreDB := newTestGenres(t)
	testGenre := createTestGenre()

	genreRes, err := genreDB.CreateGenre(context.Background(), testGenre)
	if err != nil {
		t.Fatalf("Failed to create genre: %v", err)
	}

	getGenreRes, err := genreDB.GetGenreById(context.Background(), &genre.GetGenreByIdRequest{Id: genreRes.Genre.Id})
	if err != nil {
		t.Fatalf("Failed to get genre by ID: %v", err)
	}

	assert.Equal(t, getGenreRes.Genre.Id, genreRes.Genre.Id)
	assert.Equal(t, getGenreRes.Genre.Name, testGenre.Name)
}

func TestUpdateGenre(t *testing.T) {
	genreDB := newTestGenres(t)
	testGenre := createTestGenre()

	genreRes, err := genreDB.CreateGenre(context.Background(), testGenre)
	if err != nil {
		t.Fatalf("Failed to create genre: %v", err)
	}

	updateReq := genre.UpdateGenreRequest{
		Id:   genreRes.Genre.Id,
		Name: "Updated Genre Name",
	}

	updateRes, err := genreDB.UpdateGenre(context.Background(), &updateReq)
	if err != nil {
		t.Fatalf("Failed to update genre: %v", err)
	}

	assert.Equal(t, updateRes.Genre.Id, updateReq.Id)
	assert.Equal(t, updateRes.Genre.Name, updateReq.Name)
}

func TestDeleteGenre(t *testing.T) {
	genreDB := newTestGenres(t)
	testGenre := createTestGenre()

	genreRes, err := genreDB.CreateGenre(context.Background(), testGenre)
	if err != nil {
		t.Fatalf("Failed to create genre: %v", err)
	}

	_, err = genreDB.DeleteGenre(context.Background(), &genre.DeleteGenreRequest{Id: genreRes.Genre.Id})
	if err != nil {
		t.Fatalf("Failed to delete genre: %v", err)
	}

	deletedGenre, err := genreDB.GetGenreById(context.Background(), &genre.GetGenreByIdRequest{Id: genreRes.Genre.Id})
	assert.Nil(t, deletedGenre)
	assert.Error(t, err) // Check if an error occurred when trying to get the deleted genre
}

func TestGetAllGenres(t *testing.T) {
	genreDB := newTestGenres(t)

	testGenres := []*genre.CreateGenreRequest{
		{Name: "Genre One"},
		{Name: "Genre Two"},
		{Name: "Genre Three"},
	}

	for _, gn := range testGenres {
		_, err := genreDB.CreateGenre(context.Background(), gn)
		if err != nil {
			t.Fatalf("Failed to create genre: %v", err)
		}
	}

	t.Run("GetAllGenres without filters", func(t *testing.T) {
		res, err := genreDB.GetAllGenres(context.Background(), &genre.GetGenresRequest{})
		if err != nil {
			t.Fatalf("Failed to get all genres: %v", err)
		}

		assert.GreaterOrEqual(t, len(res.Genres), len(testGenres))
	})

	t.Run("Filter by name", func(t *testing.T) {
		res, err := genreDB.GetAllGenres(context.Background(), &genre.GetGenresRequest{Name: "Genre Two"})
		if err != nil {
			t.Fatalf("Failed to filter genres by name: %v", err)
		}

		assert.Equal(t, len(res.Genres), len(res.Genres))
		assert.Equal(t, "Genre Two", res.Genres[0].Name)
	})
}
