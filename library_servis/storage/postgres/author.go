package postgres

import (
	"context"
	"fmt"
	author "library/genproto/authors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Authors struct {
	Db *pgx.Conn
}

func NewAuthor(db *pgx.Conn) *Authors {
	if db == nil {
		log.Fatal().Msg("database connection is nil")
	}
	return &Authors{Db: db}
}

func (a *Authors) CreateAuthor(ctx context.Context, req *author.CreateAuthorRequest) (*author.CreateAuthorResponse, error) {
	id := uuid.NewString()
	query := `INSERT INTO authors (id, name, biography, created_at, updated_at) 
	VALUES ($1, $2, $3, NOW(), NOW()) 
	RETURNING id, name, biography, created_at, updated_at`

	aut := author.Author{}
	var createdAt, updatedAt time.Time
	err := a.Db.QueryRow(ctx, query, id, req.Name, req.Biography).
		Scan(&aut.Id, &aut.Name, &aut.Biography, &createdAt, &updatedAt)
	if err != nil {
		log.Error().Err(err).Msg("Error Creating author")
		return nil, err
	}
	aut.CreatedAt = createdAt.Format(time.RFC3339)
	aut.UpdatedAt = updatedAt.Format(time.RFC3339)
	return &author.CreateAuthorResponse{Author: &aut}, nil
}

func (a *Authors) UpdateAuthor(ctx context.Context, req *author.UpdateAuthorRequest) (*author.UpdateAuthorResponse, error) {
	query := `
    UPDATE authors
    SET name = $1, biography = $2, updated_at = NOW()
    WHERE id = $3 AND deleted_at = 0
    RETURNING id, name, biography, created_at, updated_at`

	aut := author.Author{}
	var createdAt, updatedAt time.Time
	err := a.Db.QueryRow(ctx, query, req.Name, req.Biography, req.Id).
		Scan(
			&aut.Id,
			&aut.Name,
			&aut.Biography,
			&createdAt,
			&updatedAt,
		)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Warn().Msgf("No author found with ID %s or author is deleted", req.Id)
			return nil, status.Error(codes.NotFound, fmt.Sprintf("Author with ID %s not found or deleted", req.Id))
		}
		log.Error().Err(err).Msg("Error updating author item")
		return nil, status.Error(codes.Internal, "Error updating author")
	}

	aut.CreatedAt = createdAt.Format(time.RFC3339)
	aut.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &author.UpdateAuthorResponse{Author: &aut}, nil
}

func (a *Authors) DeleteAuthor(ctx context.Context, req *author.DeleteAuthorRequest) (*author.DeleteAuthorResponse, error) {
	query := `UPDATE authors SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0`
	result, err := a.Db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting author item")
		return nil, status.Error(codes.Internal, "Error deleting author")
	}
	affectedRows := result.RowsAffected() // Adjusted to include error handling
	if affectedRows == 0 {
		log.Warn().Msgf("No author found with ID %s or author is already deleted", req.Id)
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Author with ID %s not found or already deleted", req.Id))
	}

	return &author.DeleteAuthorResponse{Message: "Author item deleted successfully"}, nil
}


func (a *Authors) GetAuthorById(ctx context.Context, req *author.GetAuthorByIdRequest) (*author.GetAuthorByIdResponse, error) {
	query := `SELECT id, name, biography, created_at, updated_at FROM authors WHERE id = $1 AND deleted_at = 0`
	aut := author.Author{}
	var createdAt, updatedAt time.Time
	err := a.Db.QueryRow(ctx, query, req.Id).
		Scan(
			&aut.Id,
			&aut.Name,
			&aut.Biography,
			&createdAt,
			&updatedAt,
		)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Warn().Msgf("No author found with ID %s", req.Id)
			return nil, status.Error(codes.NotFound, fmt.Sprintf("Author with ID %s not found", req.Id))
		}
		log.Error().Err(err).Msg("Error getting author item")
		return nil, status.Error(codes.Internal, "Error retrieving author")
	}

	aut.CreatedAt = createdAt.Format(time.RFC3339)
	aut.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &author.GetAuthorByIdResponse{Author: &aut}, nil
}

func (a *Authors) GetAllAuthors(ctx context.Context, req *author.GetAuthorsRequest) (*author.GetAuthorsResponse, error) {
	var args []interface{}
	count := 1
	query := `SELECT id, name, biography, created_at, updated_at FROM authors WHERE deleted_at = 0`

	filter := ""

	if req.Name != "" {
		filter += fmt.Sprintf(" AND name ILIKE $%d", count)
		args = append(args, "%"+req.Name+"%")
		count++
	}

	if req.Biography != "" {
		filter += fmt.Sprintf(" AND biography ILIKE $%d", count)
		args = append(args, "%"+req.Biography+"%")
		count++
	}

	query += filter

	rows, err := a.Db.Query(ctx, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching authors from the database")
		return nil, err
	}
	defer rows.Close()

	var authors []*author.Author
	for rows.Next() {
		aut := author.Author{}
		var createdAt, updatedAt time.Time

		err := rows.Scan(
			&aut.Id,
			&aut.Name,
			&aut.Biography,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error fetching authors from the database")
			return nil, err
		}

		aut.CreatedAt = createdAt.Format(time.RFC3339) // Format as needed
		aut.UpdatedAt = updatedAt.Format(time.RFC3339) // Format as needed

		authors = append(authors, &aut)
	}

	return &author.GetAuthorsResponse{Authors: authors}, nil
}
