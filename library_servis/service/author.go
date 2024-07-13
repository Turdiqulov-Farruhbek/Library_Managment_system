package service

import (
	"context"
	"fmt"
	"library/genproto/authors"
	"library/storage"

	"github.com/rs/zerolog/log"
)

type AuthorService struct {
	stg storage.StorageI
	authors.UnimplementedAuthorsServiceServer
}

func NewAuthorService(stge storage.StorageI) *AuthorService {
	if stge == nil {
		log.Error().Msg("storage.StorageI is nil")
	}
	return &AuthorService{stg: stge}
}

func (s *AuthorService) CreateAuthor(ctx context.Context, req *authors.CreateAuthorRequest) (*authors.CreateAuthorResponse, error) {
	log.Info().Msg("AuthorService: CreateAuthor called")

	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if s.stg == nil {
		return nil, fmt.Errorf("storage is nil")
	}
	if s.stg.Author() == nil {
		return nil, fmt.Errorf("author storage is nil")
	}

	resp, err := s.stg.Author().CreateAuthor(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *AuthorService) UpdateAuthor(ctx context.Context, req *authors.UpdateAuthorRequest) (*authors.UpdateAuthorResponse, error) {
	log.Info().Msg("AuthorService: UpdateAuthor called")

	resp, err := s.stg.Author().UpdateAuthor(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *AuthorService) DeleteAuthor(ctx context.Context, req *authors.DeleteAuthorRequest) (*authors.DeleteAuthorResponse, error) {
	log.Info().Msg("AuthorService: DeleteAuthor called")

	resp, err := s.stg.Author().DeleteAuthor(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *AuthorService) GetAuthorById(ctx context.Context, req *authors.GetAuthorByIdRequest) (*authors.GetAuthorByIdResponse, error) {
	log.Info().Msg("AuthorService: GetAuthorById called")

	resp, err := s.stg.Author().GetAuthorById(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *AuthorService) GetAuthors(ctx context.Context, req *authors.GetAuthorsRequest) (*authors.GetAuthorsResponse, error) {
	log.Info().Msg("AuthorService: GetAllAuthors called")

	resp, err := s.stg.Author().GetAllAuthors(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
