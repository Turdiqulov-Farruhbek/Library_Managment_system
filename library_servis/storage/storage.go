package storage

import (
	"context"
	author "library/genproto/authors"
	book "library/genproto/books"
	borrower "library/genproto/borrowers"
	genre "library/genproto/genres"
	user "library/genproto/users"
)

type StorageI interface {
	Author() AuthorI
	Books() BooksI
	Borrowers() BorrowersI
	Genre() GenreI
	Users() UsersI
}

type AuthorI interface {
	GetAuthorById(ctx context.Context, req *author.GetAuthorByIdRequest) (*author.GetAuthorByIdResponse, error)
	GetAllAuthors(ctx context.Context, req *author.GetAuthorsRequest) (*author.GetAuthorsResponse, error)
	CreateAuthor(ctx context.Context, req *author.CreateAuthorRequest) (*author.CreateAuthorResponse, error)
	UpdateAuthor(ctx context.Context, req *author.UpdateAuthorRequest) (*author.UpdateAuthorResponse, error)
	DeleteAuthor(ctx context.Context, req *author.DeleteAuthorRequest) (*author.DeleteAuthorResponse, error)
}

type BooksI interface {
	GetBookById(ctx context.Context, req *book.GetBookByIdRequest) (*book.GetBookByIdResponse, error)
	GetallBooks(ctx context.Context, req *book.GetBooksRequest) (*book.GetBooksResponse, error)
	CreateBook(ctx context.Context, req *book.CreateBookRequest) (*book.CreateBookResponse, error)
	UpdateBook(ctx context.Context, req *book.UpdateBookRequest) (*book.UpdateBookResponse, error)
	DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (*book.DeleteBookResponse, error)
	GetBooksByAuthorId(ctx context.Context, req *book.GetBooksByAuthorIdRequest) (*book.GetBooksResponse, error)
	GetBooksByGenreId(ctx context.Context, req *book.GetBooksByGenreIdRequest) (*book.GetBooksResponse, error)
	GetOverdueBooks(ctx context.Context, req *book.GetOverdueBooksRequest) (*book.GetBooksResponse, error)
}

type BorrowersI interface {
	GetBorrowerById(ctx context.Context, req *borrower.GetBorrowerByIdRequest) (*borrower.GetBorrowerByIdResponse, error)
	GetAllBorrowers(ctx context.Context, req *borrower.GetAllBorrowersRequest) (*borrower.GetAllBorrowersResponse, error)
	CreateBorrower(ctx context.Context, req *borrower.CreateBorrowerRequest) (*borrower.CreateBorrowerResponse, error)
	UpdateBorrower(ctx context.Context, req *borrower.UpdateBorrowerRequest) (*borrower.UpdateBorrowerResponse, error)
	DeleteBorrower(ctx context.Context, req *borrower.DeleteBorrowerRequest) (*borrower.DeleteBorrowerResponse, error)
	GetBorrowingHistory(ctx context.Context, req *borrower.GetBorrowingHistoryRequest) (*borrower.GetBorrowingHistoryResponse, error)
	GetBorrowedBooksByUser(ctx context.Context, req *borrower.GetBorrowedBooksByUserRequest) (*borrower.GetBorrowedBooksByUserResponse, error)
}

type GenreI interface {
	// GetByGenreBooks(ctx context.Context, req *genre.GetgenresByGenreIdRequest) (*genre.GetgenresResponse, error)
	GetGenreById(ctx context.Context, req *genre.GetGenreByIdRequest) (*genre.GetGenreByIdResponse, error)
	GetAllGenres(ctx context.Context, req *genre.GetGenresRequest) (*genre.GetGenresResponse, error)
	CreateGenre(ctx context.Context, req *genre.CreateGenreRequest) (*genre.CreateGenreResponse, error)
	UpdateGenre(ctx context.Context, req *genre.UpdateGenreRequest) (*genre.UpdateGenreResponse, error)
	DeleteGenre(ctx context.Context, req *genre.DeleteGenreRequest) (*genre.DeleteGenreResponse, error)
	
}

type UsersI interface {
	GetUserById(ctx context.Context, req *user.GetUserByIdRequest) (*user.GetUserByIdResponse, error)
	GetAllUsers(ctx context.Context, req *user.GetAllUsersRequest) (*user.GetAllUsersResponse, error)
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error)
	UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error)
}
