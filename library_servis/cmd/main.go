package main

// ` eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg3OTQ4MDYsImlhdCI6MTcxODcwODQwNiwiaWQiOiJGYXJydWhiZWsiLCJ1c2VybmFtZSI6ImExNzQzMDRkLTQxOTYtNDFmMC05NmQ4LThlYWEyMjI0NTRlZSJ9.cbVpviAzbCJPLhBYFc8EZ3sCKnvHkpLWoF1YChFkWmw`

import (
	"library/genproto/authors"
	"library/genproto/books"
	"library/genproto/borrowers"
	"library/genproto/genres"
	"library/genproto/users"

	"library/service"
	"library/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.DBConn()
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}
	liss, err := net.Listen("tcp", ":50020")
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}

	server := grpc.NewServer()
	authors.RegisterAuthorsServiceServer(server, service.NewAuthorService(db))
	books.RegisterBooksServiceServer(server, service.NewBooksService(db))
	borrowers.RegisterBorrowersServiceServer(server, service.NewBorrowersService(db))
	genres.RegisterGenresServiceServer(server, service.NewGenresService(db))
	users.RegisterUsersServiceServer(server, service.NewUsersService(db))

	// log.Println("Server is running on port :50020")
	if err := server.Serve(liss); err != nil {
		log.Fatalf("error listening: %v", err)
	}

}
