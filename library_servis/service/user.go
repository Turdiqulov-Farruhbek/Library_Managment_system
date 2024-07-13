package service

import (
	"context"
	"library/genproto/users"
	"library/storage"

	"github.com/rs/zerolog/log"
)

// UsersService struct handles user-related operations by interacting with the storage layer.
type UsersService struct {
	stg storage.StorageI
	users.UnimplementedUsersServiceServer
}

// NewUsersService initializes and returns a new UsersService instance.
func NewUsersService(stg storage.StorageI) *UsersService {
	return &UsersService{stg: stg}
}

// CreateUser creates a new user.
func (s *UsersService) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	log.Info().Msg("UsersService: CreateUser called")

	resp, err := s.stg.Users().CreateUser(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("UsersService: Error creating user")
		return nil, err
	}

	return resp, nil
}

// UpdateUser updates an existing user.
func (s *UsersService) UpdateUser(ctx context.Context, req *users.UpdateUserRequest) (*users.UpdateUserResponse, error) {
	log.Info().Msg("UsersService: UpdateUser called")

	resp, err := s.stg.Users().UpdateUser(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("UsersService: Error updating user")
		return nil, err
	}

	return resp, nil
}

// DeleteUser marks a user as deleted.
func (s *UsersService) DeleteUser(ctx context.Context, req *users.DeleteUserRequest) (*users.DeleteUserResponse, error) {
	log.Info().Msg("UsersService: DeleteUser called")

	resp, err := s.stg.Users().DeleteUser(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("UsersService: Error deleting user")
		return nil, err
	}

	return resp, nil
}

// GetUserById retrieves a user by their ID.
func (s *UsersService) GetUserById(ctx context.Context, req *users.GetUserByIdRequest) (*users.GetUserByIdResponse, error) {
	log.Info().Msg("UsersService: GetUserById called")

	resp, err := s.stg.Users().GetUserById(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("UsersService: Error getting user by ID")
		return nil, err
	}

	return resp, nil
}

// GetAllUsers retrieves all users.
func (s *UsersService) GetAllUsers(ctx context.Context, req *users.GetAllUsersRequest) (*users.GetAllUsersResponse, error) {
	log.Info().Msg("UsersService: GetAllUsers called")

	resp, err := s.stg.Users().GetAllUsers(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("UsersService: Error getting all users")
		return nil, err
	}

	return resp, nil
}
