package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Project_Restaurant/Auth-Service/models"
	t "github.com/Project_Restaurant/Auth-Service/token"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}


func (u *UserRepo) Register(user models.UserRegister) (*models.LoginRes, error) {
	var us models.LoginRes
	id := uuid.NewString()
	hashPassword, err := t.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}

	err = u.db.QueryRow("insert into users(id, username, password, email) values ($1, $2, $3, $4) returning id, username", id, user.Name, hashPassword, user.Email).
		Scan(&us.ID, &us.Name)
	if err != nil {
		return nil, err
	}
	return &us, nil
}

func (u *UserRepo) Login(user models.UserLogin) (*models.LoginRes, error) {
	res := models.LoginRes{}

	var hashedPassword string

	err := u.db.QueryRow("select id, username, password from users where username = $1 ", user.Name).Scan(&res.ID, &res.Name, &hashedPassword)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return nil, err
	}
	fmt.Println(res.ID, res.Name)
	return &res, nil
}

func (u *UserRepo) GetByUsername(username string) (*models.User, error) {
	user := models.User{}
	err := u.db.QueryRow("select id, username, email, created_at, updated_at, deleted_at from users where username = $1", username).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
