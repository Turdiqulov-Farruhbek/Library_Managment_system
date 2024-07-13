package handlers

import (
	"getway/genproto/authors"
	"getway/genproto/books"
	"getway/genproto/borrowers"
	"getway/genproto/genres"
	"getway/genproto/users"

	"google.golang.org/grpc"
)

type HandlerStruct struct {
	Author authors.AuthorsServiceClient
	Books  books.BooksServiceClient
	Borrowers borrowers.BorrowersServiceClient
	Genres  genres.GenresServiceClient
	Users  users.UsersServiceClient
}


func NewHandler(libraryConn *grpc.ClientConn ) *HandlerStruct {
	return &HandlerStruct{
        Author: authors.NewAuthorsServiceClient(libraryConn),
        Books:  books.NewBooksServiceClient(libraryConn),
        Borrowers: borrowers.NewBorrowersServiceClient(libraryConn),
        Genres:  genres.NewGenresServiceClient(libraryConn),
        Users:  users.NewUsersServiceClient(libraryConn),
    }
}