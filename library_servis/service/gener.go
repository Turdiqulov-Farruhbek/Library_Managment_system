package service

import (
	"context"
	"library/genproto/genres"
	"library/storage"

	"github.com/rs/zerolog/log"
)

// GenresService struct handles genre-related operations by interacting with the storage layer.
type GenresService struct {
	stg storage.StorageI
	genres.UnimplementedGenresServiceServer
}

// NewGenresService initializes and returns a new GenresService instance.
func NewGenresService(stg storage.StorageI) *GenresService {
	return &GenresService{stg: stg}
}

// CreateGenre creates a new genre.
func (s *GenresService) CreateGenre(ctx context.Context, req *genres.CreateGenreRequest) (*genres.CreateGenreResponse, error) {
	log.Info().Msg("GenresService: CreateGenre called")

	resp, err := s.stg.Genre().CreateGenre(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("GenresService: Error creating genre")
		return nil, err
	}

	return resp, nil
}

// UpdateGenre updates an existing genre.
func (s *GenresService) UpdateGenre(ctx context.Context, req *genres.UpdateGenreRequest) (*genres.UpdateGenreResponse, error) {
	log.Info().Msg("GenresService: UpdateGenre called")

	resp, err := s.stg.Genre().UpdateGenre(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("GenresService: Error updating genre")
		return nil, err
	}

	return resp, nil
}

// DeleteGenre marks a genre as deleted.
func (s *GenresService) DeleteGenre(ctx context.Context, req *genres.DeleteGenreRequest) (*genres.DeleteGenreResponse, error) {
	log.Info().Msg("GenresService: DeleteGenre called")

	resp, err := s.stg.Genre().DeleteGenre(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("GenresService: Error deleting genre")
		return nil, err
	}

	return resp, nil
}

// GetGenreById retrieves a genre by its ID.
func (s *GenresService) GetGenreById(ctx context.Context, req *genres.GetGenreByIdRequest) (*genres.GetGenreByIdResponse, error) {
	log.Info().Msg("GenresService: GetGenreById called")

	resp, err := s.stg.Genre().GetGenreById(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("GenresService: Error getting genre by ID")
		return nil, err
	}

	return resp, nil
}

// func (s *GenresService) GetByGenreIdBooks(ctx context.Context, req *genres.GetgenresByGenreIdRequest) (*genres.GetgenresResponse, error) {
// 	log.Info().Msg("Fetching genres by genre ID")

// 	resp, err := s.stg.Genre().GetByGenreBooks(ctx, req)
// 	if err != nil {
// 		log.Error().Err(err).Msg("Error fetching genres by genre ID")
// 		return nil, err
// 	}

// 	return resp, nil
// }


// GetAllGenres retrieves all genres.
func (s *GenresService) GetGenres(ctx context.Context, req *genres.GetGenresRequest) (*genres.GetGenresResponse, error) {
	log.Info().Msg("GenresService: GetAllGenres called")

	resp, err := s.stg.Genre().GetAllGenres(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("GenresService: Error getting all genres")
		return nil, err
	}

	return resp, nil
}
