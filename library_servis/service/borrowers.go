package service

import (
	"context"
	"library/genproto/borrowers"
	"library/storage"

	"github.com/rs/zerolog/log"
)

// BorrowersService struct handles borrower-related operations by interacting with the storage layer.
type BorrowersService struct {
	stg storage.StorageI
	borrowers.UnimplementedBorrowersServiceServer
}

// NewBorrowersService initializes and returns a new BorrowersService instance.
func NewBorrowersService(stg storage.StorageI) *BorrowersService {
	return &BorrowersService{stg: stg}
}

// CreateService creates a new borrower record.
func (s *BorrowersService) CreateBorrower(ctx context.Context, req *borrowers.CreateBorrowerRequest) (*borrowers.CreateBorrowerResponse, error) {
	log.Info().Msg("BorrowersService: CreateBorrower called")

	resp, err := s.stg.Borrowers().CreateBorrower(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BorrowersService: Error creating borrower")
		return nil, err
	}

	return resp, nil
}

// UpdateService updates an existing borrower record.
func (s *BorrowersService) UpdateBorrower(ctx context.Context, req *borrowers.UpdateBorrowerRequest) (*borrowers.UpdateBorrowerResponse, error) {
	log.Info().Msg("BorrowersService: UpdateBorrower called")

	resp, err := s.stg.Borrowers().UpdateBorrower(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BorrowersService: Error updating borrower")
		return nil, err
	}

	return resp, nil
}

// DeleteService deletes a borrower record by marking it as deleted.
func (s *BorrowersService) DeleteBorrower(ctx context.Context, req *borrowers.DeleteBorrowerRequest) (*borrowers.DeleteBorrowerResponse, error) {
	log.Info().Msg("BorrowersService: DeleteBorrower called")

	resp, err := s.stg.Borrowers().DeleteBorrower(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BorrowersService: Error deleting borrower")
		return nil, err
	}

	return resp, nil
}

// GetBorrowerService retrieves a borrower by its ID.
func (s *BorrowersService) GetBorrowerById(ctx context.Context, req *borrowers.GetBorrowerByIdRequest) (*borrowers.GetBorrowerByIdResponse, error) {
	log.Info().Msg("BorrowersService: GetBorrowerById called")

	resp, err := s.stg.Borrowers().GetBorrowerById(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BorrowersService: Error getting borrower by ID")
		return nil, err
	}

	return resp, nil
}

// GetAllBorrowersService retrieves all borrowers or those matching the filter criteria.
func (s *BorrowersService) GetAllBorrowers(ctx context.Context, req *borrowers.GetAllBorrowersRequest) (*borrowers.GetAllBorrowersResponse, error) {
	log.Info().Msg("BorrowersService: GetAllBorrowers called")

	resp, err := s.stg.Borrowers().GetAllBorrowers(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BorrowersService: Error getting all borrowers")
		return nil, err
	}

	return resp, nil
}

// GetBorrowingHistoryService retrieves borrowing history for a user.
func (s *BorrowersService) GetBorrowingHistory(ctx context.Context, req *borrowers.GetBorrowingHistoryRequest) (*borrowers.GetBorrowingHistoryResponse, error) {
	log.Info().Msg("BorrowersService: GetBorrowingHistory called")

	resp, err := s.stg.Borrowers().GetBorrowingHistory(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BorrowersService: Error getting borrowing history")
		return nil, err
	}

	return resp, nil
}

// GetBorrowedBooksByUserService retrieves books currently borrowed by a user.
func (s *BorrowersService) GetBorrowedBooksByUser(ctx context.Context, req *borrowers.GetBorrowedBooksByUserRequest) (*borrowers.GetBorrowedBooksByUserResponse, error) {
	log.Info().Msg("BorrowersService: GetBorrowedBooksByUser called")

	resp, err := s.stg.Borrowers().GetBorrowedBooksByUser(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("BorrowersService: Error getting borrowing User")
	}

	return resp, nil
}
