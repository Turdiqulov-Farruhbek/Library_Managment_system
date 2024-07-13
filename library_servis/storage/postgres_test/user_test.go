package postgrestest

// import (
// 	"context"
// 	"fmt"
// 	"library/genproto/users"
// 	"library/storage/postgres"
// 	"testing"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/stretchr/testify/assert"
// )

// // Function to establish a new connection to the test database
// func newTestUsers(t *testing.T) *postgres.Users {
// 	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
// 		"postgres",
// 		"root",
// 		"localhost",
// 		5432,
// 		"library_db",
// 	)

// 	db, err := pgx.Connect(context.Background(), connString)
// 	if err != nil {
// 		t.Fatalf("Failed to connect to database: %v", err)
// 	}
// 	return postgres.NewUsers(db)
// }

// // Helper function to create a test user
// func createTestUser() *users.CreateUserRequest {
// 	return &users.CreateUserRequest{
// 		Username: "testuser",
// 		Email:    "testuser@example.com",
// 		Password: "password123",
// 	}
// }

// func TestCreateUser(t *testing.T) {
// 	userDB := newTestUsers(t)
// 	testUser := createTestUser()

// 	userRes, err := userDB.CreateUser(context.Background(), testUser)
// 	if err != nil {
// 		t.Fatalf("Failed to create user: %v", err)
// 	}

// 	assert.NotEmpty(t, userRes.User.Id)
// 	assert.Equal(t, testUser.Username, userRes.User.Username)
// 	assert.Equal(t, testUser.Email, userRes.User.Email)
// }

// func TestGetUserById(t *testing.T) {
// 	userDB := newTestUsers(t)
// 	testUser := createTestUser()

// 	userRes, err := userDB.CreateUser(context.Background(), testUser)
// 	if err != nil {
// 		t.Fatalf("Failed to create user: %v", err)
// 	}

// 	getUserRes, err := userDB.GetUserById(context.Background(), &users.GetUserByIdRequest{Id: userRes.User.Id})
// 	if err != nil {
// 		t.Fatalf("Failed to get user by ID: %v", err)
// 	}

// 	assert.Equal(t, userRes.User.Id, getUserRes.User.Id)
// 	assert.Equal(t, userRes.User.Username, getUserRes.User.Username)
// 	assert.Equal(t, userRes.User.Email, getUserRes.User.Email)
// }

// func TestUpdateUser(t *testing.T) {
// 	userDB := newTestUsers(t)
// 	testUser := createTestUser()

// 	userRes, err := userDB.CreateUser(context.Background(), testUser)
// 	if err != nil {
// 		t.Fatalf("Failed to create user: %v", err)
// 	}

// 	updateReq := users.UpdateUserRequest{
// 		Id:       userRes.User.Id,
// 		Username: "updateduser",
// 		Email:    "updateduser@example.com",
// 		Password: "newpassword123",
// 	}

// 	updateRes, err := userDB.UpdateUser(context.Background(), &updateReq)
// 	if err != nil {
// 		t.Fatalf("Failed to update user: %v", err)
// 	}

// 	assert.Equal(t, updateReq.Id, updateRes.User.Id)
// 	assert.Equal(t, updateReq.Username, updateRes.User.Username)
// 	assert.Equal(t, updateReq.Email, updateRes.User.Email)
// }

// func TestDeleteUser(t *testing.T) {
// 	userDB := newTestUsers(t)
// 	testUser := createTestUser()

// 	userRes, err := userDB.CreateUser(context.Background(), testUser)
// 	if err != nil {
// 		t.Fatalf("Failed to create user: %v", err)
// 	}

// 	_, err = userDB.DeleteUser(context.Background(), &users.DeleteUserRequest{Id: userRes.User.Id})
// 	if err != nil {
// 		t.Fatalf("Failed to delete user: %v", err)
// 	}

// 	getUserRes, err := userDB.GetUserById(context.Background(), &users.GetUserByIdRequest{Id: userRes.User.Id})
// 	assert.Nil(t, getUserRes)
// 	assert.Error(t, err)
// }

// func TestGetAllUsers(t *testing.T) {
// 	userDB := newTestUsers(t)

// 	testUsers := []*users.CreateUserRequest{
// 		{Username: "userone", Email: "userone@example.com", Password: "password123"},
// 		{Username: "usertwo", Email: "usertwo@example.com", Password: "password123"},
// 		{Username: "userthree", Email: "userthree@example.com", Password: "password123"},
// 	}

// 	for _, usr := range testUsers {
// 		_, err := userDB.CreateUser(context.Background(), usr)
// 		if err != nil {
// 			t.Fatalf("Failed to create user: %v", err)
// 		}
// 	}

// 	res, err := userDB.GetAllUsers(context.Background(), &users.GetAllUsersRequest{})
// 	if err != nil {
// 		t.Fatalf("Failed to get all users: %v", err)
// 	}

// 	assert.GreaterOrEqual(t, len(res.Users), len(testUsers))
// }

// func TestGetAllUsersFiltered(t *testing.T) {
// 	userDB := newTestUsers(t)

// 	testUsers := []*users.CreateUserRequest{
// 		{Username: "userfour", Email: "userfour@example.com", Password: "password123"},
// 		{Username: "userfive", Email: "userfive@example.com", Password: "password123"},
// 	}

// 	for _, usr := range testUsers {
// 		_, err := userDB.CreateUser(context.Background(), usr)
// 		if err != nil {
// 			t.Fatalf("Failed to create user: %v", err)
// 		}
// 	}

// 	res, err := userDB.GetAllUsers(context.Background(), &users.GetAllUsersRequest{})
// 	if err != nil {
// 		t.Fatalf("Failed to get all users: %v", err)
// 	}

// 	assert.GreaterOrEqual(t, len(res.Users), len(testUsers))

// 	// Filter by username
// 	res, err = userDB.GetAllUsers(context.Background(), &users.GetAllUsersRequest{})
// 	if err != nil {
// 		t.Fatalf("Failed to get all users with filter: %v", err)
// 	}

// 	assert.Equal(t, 1, len(res.Users))
// 	assert.Equal(t, "userfive", res.Users[0].Username)
// }
