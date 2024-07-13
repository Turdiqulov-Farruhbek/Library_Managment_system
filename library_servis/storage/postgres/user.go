package postgres

import (
	"context"
	"library/genproto/users"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"github.com/google/uuid"
	"time"
)

type Users struct {
	Db *pgx.Conn
}

func NewUsers(db *pgx.Conn) *Users {
	return &Users{
		Db: db,
	}
}

// CreateUser creates a new user record in the database.
func (u *Users) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	id := uuid.NewString()
	query := `
	INSERT INTO users (id, username, email, password) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, username, email, created_at, updated_at`

	var user users.User
	err := u.Db.QueryRow(ctx, query, id, req.Username, req.Email, req.Password).
		Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Error().Err(err).Msg("Error creating user")
		return nil, err
	}

	return &users.CreateUserResponse{User: &user}, nil
}

// UpdateUser updates an existing user record in the database.
func (u *Users) UpdateUser(ctx context.Context, req *users.UpdateUserRequest) (*users.UpdateUserResponse, error) {
	query := `
	UPDATE users
	SET username = $1, email = $2, password = $3, updated_at = CURRENT_TIMESTAMP
	WHERE id = $4 AND deleted_at = 0
	RETURNING id, username, email, created_at, updated_at`

	var user users.User
	err := u.Db.QueryRow(ctx, query, req.Username, req.Email, req.Password, req.Id).
		Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Error().Err(err).Msg("Error updating user")
		return nil, err
	}

	return &users.UpdateUserResponse{User: &user}, nil
}

// DeleteUser marks a user record as deleted in the database.
func (u *Users) DeleteUser(ctx context.Context, req *users.DeleteUserRequest) (*users.DeleteUserResponse, error) {
	query := `UPDATE users SET deleted_at = $1 WHERE id = $2 AND deleted_at = 0`
	_, err := u.Db.Exec(ctx, query, time.Now().Unix(), req.Id)

	if err != nil {
		log.Error().Err(err).Msg("Error deleting user")
		return nil, err
	}

	return &users.DeleteUserResponse{Message: "User deleted successfully"}, nil
}

// GetUserById fetches a user record by its ID.
func (u *Users) GetUserById(ctx context.Context, req *users.GetUserByIdRequest) (*users.GetUserByIdResponse, error) {
	query := `SELECT id, username, email, created_at, updated_at FROM users WHERE id = $1 AND deleted_at = 0`
	var user users.User
	err := u.Db.QueryRow(ctx, query, req.Id).
		Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Error().Err(err).Msg("Error getting user by ID")
		return nil, err
	}

	return &users.GetUserByIdResponse{User: &user}, nil
}

// GetAllUsers fetches all user records from the database.
func (u *Users) GetAllUsers(ctx context.Context, req *users.GetAllUsersRequest) (*users.GetAllUsersResponse, error) {
	query := `SELECT id, username, email, created_at, updated_at FROM users WHERE deleted_at = 0`
	rows, err := u.Db.Query(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching all users")
		return nil, err
	}
	defer rows.Close()

	var usersList []*users.User
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning user row")
			return nil, err
		}
		usersList = append(usersList, &user)
	}

	return &users.GetAllUsersResponse{Users: usersList}, nil
}
