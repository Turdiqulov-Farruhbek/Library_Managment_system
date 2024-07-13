package postgres

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Project_Restaurant/Auth-Service/models"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_Register(t *testing.T) {
	db,err := ConnectDb()
	if err!= nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
	userRepo := NewUserRepo(db)

	user := models.UserRegister{
		Name:     "testuser",
		Password: "password",
		Email:    "testuser@example.com",
	}
	
	
	// Call the Register method
	res, err := userRepo.Register(user)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, 1, res.ID)
	assert.Equal(t, "testuser", res.Name)
}

func TestUserRepo_Login(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userRepo := NewUserRepo(db)

	user := models.UserLogin{
		Name:     "testuser",
		Password: "password",
	}

	// Define the expected query and result
	mock.ExpectQuery("select id, username from users where username = \\$1 and password = \\$2").
		WithArgs(user.Name, user.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).AddRow(1, user.Name))

	// Call the Login method
	res, err := userRepo.Login(user)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, 1, res.ID)
	assert.Equal(t, "testuser", res.Name)
}

func TestUserRepo_GetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userRepo := NewUserRepo(db)

	username := "testuser"

	// Define the expected query and result
	mock.ExpectQuery("select \\* from users where username = \\$1").
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "password", "created_at", "updated_at", "deleted_at"}).
			AddRow(1, username, "testuser@example.com", "password", time.Now(), time.Now(), nil))

	// Call the GetByUsername method
	res, err := userRepo.GetByUsername(username)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, 1, res.ID)
	assert.Equal(t, "testuser", res.Name)
	assert.Equal(t, "testuser@example.com", res.Email)
}
