package postgres

import (
	"context"
	"fmt"
	genre "library/genproto/genres"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

type Genres struct {
	Db *pgx.Conn
}

func NewGenre(db *pgx.Conn) *Genres {
	return &Genres{
		Db: db,
	}
}

func (g *Genres) CreateGenre(ctx context.Context, req *genre.CreateGenreRequest) (*genre.CreateGenreResponse, error) {
	id := uuid.NewString()
	query := `
    INSERT INTO genres (id, name) 
    VALUES ($1, $2)
    RETURNING id, name, created_at, updated_at`

	var gn genre.Genre
	var createdAt, updatedAt time.Time
	err := g.Db.QueryRow(ctx, query, id, req.Name).
		Scan(
			&gn.Id,
			&gn.Name,
			&createdAt,
			&updatedAt,
		)
	if err != nil {
		log.Error().Err(err).Msg("Error creating genre")
		return nil, err
	}

	gn.CreatedAt = createdAt.Format(time.RFC3339)
	gn.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &genre.CreateGenreResponse{Genre: &gn}, nil
}

func (g *Genres) UpdateGenre(ctx context.Context, req *genre.UpdateGenreRequest) (*genre.UpdateGenreResponse, error) {
	query := `
        UPDATE genres
        SET name = $1, updated_at = CURRENT_TIMESTAMP
        WHERE id = $2 AND deleted_at = 0
        RETURNING id, name, created_at, updated_at
    `

	var gn genre.Genre
	var createdAt, updatedAt pq.NullTime

	err := g.Db.QueryRow(ctx, query, req.Name, req.Id).
		Scan(
			&gn.Id,
			&gn.Name,
			&createdAt,
			&updatedAt,
		)
	if err != nil {
		log.Error().Err(err).Msg("Error updating genre")
		return nil, err
	}

	if createdAt.Valid {
		gn.CreatedAt = createdAt.Time.Format(time.RFC3339)
	}
	if updatedAt.Valid {
		gn.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &genre.UpdateGenreResponse{Genre: &gn}, nil
}

func (g *Genres) DeleteGenre(ctx context.Context, req *genre.DeleteGenreRequest) (*genre.DeleteGenreResponse, error) {
	query := `UPDATE genres SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0`
	_, err := g.Db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting genre")
		return nil, err
	}
	return &genre.DeleteGenreResponse{Message: "Genre deleted successfully"}, nil
}

func (g *Genres) GetGenreById(ctx context.Context, req *genre.GetGenreByIdRequest) (*genre.GetGenreByIdResponse, error) {
	query := `SELECT id, name, created_at, updated_at FROM genres WHERE id = $1 AND deleted_at = 0`
	gn := genre.Genre{}
	var createdAt, updatedAt pq.NullTime // Use pq.NullTime to handle NULL values

	err := g.Db.QueryRow(ctx, query, req.Id).
		Scan(
			&gn.Id,
			&gn.Name,
			&createdAt,
			&updatedAt,
		)
	if err != nil {
		log.Error().Err(err).Msg("Error getting genre by id")
		return nil, err
	}

	if createdAt.Valid {
		gn.CreatedAt = createdAt.Time.Format(time.RFC3339)
	}
	if updatedAt.Valid {
		gn.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &genre.GetGenreByIdResponse{Genre: &gn}, nil
}

func (b *Books) GetByGenreBooks(ctx context.Context, req *genre.GetgenresByGenreIdRequest) (*genre.GetgenresResponse, error) {
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

	var books []*genre.Book
	for rows.Next() {
		bk := genre.Book{}
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

	return &genre.GetgenresResponse{Books: books}, nil
}

func (g *Genres) GetAllGenres(ctx context.Context, req *genre.GetGenresRequest) (*genre.GetGenresResponse, error) {
	var args []interface{}
	count := 1
	query := `SELECT id, name, created_at, updated_at FROM genres WHERE deleted_at = 0`

	if req.Name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", count)
		args = append(args, "%"+req.Name+"%")
		count++
	}

	rows, err := g.Db.Query(ctx, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching genres from the database")
		return nil, err
	}
	defer rows.Close()

	var genres []*genre.Genre
	for rows.Next() {
		var gn genre.Genre
		var createdAt, updatedAt time.Time
		err := rows.Scan(
			&gn.Id,
			&gn.Name,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning genre row")
			return nil, err
		}
		gn.CreatedAt = createdAt.Format(time.RFC3339)
		gn.UpdatedAt = updatedAt.Format(time.RFC3339)
		genres = append(genres, &gn)
	}

	return &genre.GetGenresResponse{Genres: genres}, nil
}
