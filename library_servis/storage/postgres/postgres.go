package postgres

import (
	"context"
	"fmt"
	"library/config"
	"library/storage"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type StorageStruct struct {
	DB          *pgx.Conn
	Author_S    storage.AuthorI
	Books_S     storage.BooksI
	Borrowers_S storage.BorrowersI
	Genre_S     storage.GenreI
	Users_S     storage.UsersI
}

func DBConn() (*StorageStruct, error) {
	var (
		db  *pgx.Conn
		err error
	)
	cfg := config.Load()
	dbCon := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	db, err = pgx.Connect(context.Background(), dbCon)
	if err != nil {
		slog.Warn("Unable to connect to database:", err)
		return nil, err
	}
	err = db.Ping(context.Background())
	if err != nil {
		slog.Warn("Unable to ping database:", err)
		return nil, err
	}

	return &StorageStruct{
		DB: db,
	}, nil
}

func (s *StorageStruct) Author() storage.AuthorI {
	if s.Author_S == nil {
		s.Author_S = NewAuthor(s.DB)
	}

	return s.Author_S
}

func (s *StorageStruct) Books() storage.BooksI {
	if s.Books_S == nil {
        s.Books_S = NewBooks(s.DB)
    }

    return s.Books_S
}

func (s *StorageStruct) Borrowers() storage.BorrowersI {
	if s.Borrowers_S == nil {
        s.Borrowers_S = NewBorrowers(s.DB)
    }

    return s.Borrowers_S
}

func (s *StorageStruct) Genre() storage.GenreI {
	if s.Genre_S == nil {
        s.Genre_S = NewGenre(s.DB)
    }

    return s.Genre_S
}

func (s *StorageStruct) Users() storage.UsersI {
	if s.Users_S == nil {
        s.Users_S = NewUsers(s.DB)
    }

    return s.Users_S
}
