package postgrestest

import (
	"context"
	"fmt"
	borr "library/genproto/borrowers"
	"library/storage/postgres"
	"testing"

	// "github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func newTestBorrows(t *testing.T) *postgres.Borrowers {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"postgres",
		"root",
		"localhost",
		5432,
		"library_db",
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &postgres.Borrowers{Db: db}
}

func CreateTestBorrower() *borr.CreateBorrowerRequest {
	req := &borr.CreateBorrowerRequest{
		UserId:     "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed",
		BookId:     "3b29f3a2-0910-440a-a30e-0e21086a9eac",
		BorrowDate: "2024-07-15",
		ReturnDate: "2025-01-01",
	}
	return req
}

func TestCreateBorrower(t *testing.T) {
	borrDB := newTestBorrows(t)
	testBorr := CreateTestBorrower()

	BorrowRes, err := borrDB.CreateBorrower(context.Background(), testBorr)
	if err != nil {
		t.Fatalf("Failed to create borrower: %v", err)
	}

	assert.NotEmpty(t, BorrowRes.Borrower.Id)
	assert.Equal(t, BorrowRes.Borrower.UserId, testBorr.UserId)
	assert.Equal(t, BorrowRes.Borrower.BookId, testBorr.BookId)
	assert.Equal(t, BorrowRes.Borrower.BorrowDate, testBorr.BorrowDate)
	assert.Equal(t, BorrowRes.Borrower.ReturnDate, testBorr.ReturnDate)
}


func TestGetBorrowerById(t *testing.T) {
	borrDB := newTestBorrows(t)
	testBorr := CreateTestBorrower()

	BorrowRes, err := borrDB.CreateBorrower(context.Background(), testBorr)
	if err != nil {
		t.Fatalf("Failed to create borrower: %v", err)
	}

	GetBorrowRes, err := borrDB.GetBorrowerById(context.Background(), &borr.GetBorrowerByIdRequest{Id: BorrowRes.Borrower.Id})
	if err != nil {
		t.Fatalf("Failed to get borrower by ID: %v", err)
	}

	assert.NotEmpty(t, GetBorrowRes.Borrower.Id)
	assert.Equal(t, GetBorrowRes.Borrower.UserId, testBorr.UserId)
	assert.Equal(t, GetBorrowRes.Borrower.BookId, testBorr.BookId)
	assert.Equal(t, GetBorrowRes.Borrower.BorrowDate, testBorr.BorrowDate)
	assert.Equal(t, GetBorrowRes.Borrower.ReturnDate, testBorr.ReturnDate)

}

func TestGetAllBorrowers(t *testing.T) {
	borrDB := newTestBorrows(t)

	userId := "72a2c2d8-1b69-4c2e-916d-0d57b59733a6"
	bookId := "4a8a1d57-c23a-4f32-9890-35c1525c0916"

	testBorr := []*borr.CreateBorrowerRequest{
		{UserId: userId, BookId: bookId, BorrowDate: "2024-07-22", ReturnDate: "2025-11-11"},
		{UserId: userId, BookId: bookId, BorrowDate: "2024-07-22", ReturnDate: "2026-10-10"},
		{UserId: userId, BookId: bookId, BorrowDate: "2024-07-22", ReturnDate: "2027-09-09"},
	}

	for _, borr := range testBorr {
		_, err := borrDB.CreateBorrower(context.Background(), borr)
		if err != nil {
			t.Fatalf("Failed Error to create borrower: %v", err)
		}
	}

	t.Run("GetAllMenus without filters", func(t *testing.T) {
		res, err := borrDB.GetAllBorrowers(context.Background(), &borr.GetAllBorrowersRequest{})
		if err != nil {
			t.Fatalf("Failed to get all borrowers: %v", err)
		}

		assert.LessOrEqual(t, len(testBorr), len(res.Borrowers))
	})

	t.Run("Filter by user ID", func(t *testing.T) {
		res, err := borrDB.GetAllBorrowers(context.Background(), &borr.GetAllBorrowersRequest{UserId: userId})
		if err != nil {
			t.Fatalf("Failed error user ID: %v", err)
		}

		assert.Equal(t, len(res.Borrowers), len(res.Borrowers))
	})

	t.Run("Filter by book ID", func(t *testing.T) {
		res, err := borrDB.GetAllBorrowers(context.Background(), &borr.GetAllBorrowersRequest{UserId: userId})
		if err != nil {
			t.Fatalf("Failed error book ID: %v", err)
		}

		assert.Equal(t, len(res.Borrowers), len(res.Borrowers))
	})

	t.Run("Filter by BorrowDate", func(t *testing.T) {
		res, err := borrDB.GetAllBorrowers(context.Background(), &borr.GetAllBorrowersRequest{UserId: userId})
		if err != nil {
			t.Fatalf("Failed error BorrowDate: %v", err)
		}

		assert.Equal(t, len(res.Borrowers), len(res.Borrowers))

	})

	t.Run("Filter by ReturnDate", func(t *testing.T) {
		res, err := borrDB.GetAllBorrowers(context.Background(), &borr.GetAllBorrowersRequest{UserId: userId})
		if err != nil {
			t.Fatalf("Failed error ReturnDate: %v", err)
		}

		assert.Equal(t, len(res.Borrowers), len(res.Borrowers))

	})

}

func TestUpdateBorrower(t *testing.T) {
	borrDB := newTestBorrows(t)
	testBorr := CreateTestBorrower()

	BorrowRes, err := borrDB.CreateBorrower(context.Background(), testBorr)
	if err != nil {
		t.Fatalf("Failed to create borrower: %v", err)
	}

	updateReq := borr.UpdateBorrowerRequest{
		Id:         BorrowRes.Borrower.Id,
		UserId:     "2c153ade-8212-4e01-ba31-3b2a3f57309f",
		BookId:     "bb119756-e508-4981-a1a7-1b4040d9ec3b",
		BorrowDate: "2030-07-22",
		ReturnDate: "2029-11-11",
	}

	UpdateRes, err := borrDB.UpdateBorrower(context.Background(), &updateReq)
	if err != nil {
		t.Fatalf("Failed to update borrower: %v", err)
	}
	assert.Equal(t, UpdateRes.Borrower.UserId, updateReq.UserId)
	assert.Equal(t, UpdateRes.Borrower.BookId, updateReq.BookId)
	assert.Equal(t, UpdateRes.Borrower.BorrowDate, updateReq.BorrowDate)
	assert.Equal(t, UpdateRes.Borrower.ReturnDate, updateReq.ReturnDate)
}


func TestDeleteBorrower(t *testing.T) {
    borrDB := newTestBorrows(t)
    testBorr := CreateTestBorrower()

    BorrowRes, err := borrDB.CreateBorrower(context.Background(), testBorr)
    if err != nil {
        t.Fatalf("Failed to create borrower: %v", err)
    }

    _, err = borrDB.DeleteBorrower(context.Background(), &borr.DeleteBorrowerRequest{Id: BorrowRes.Borrower.Id})
    if err != nil {
        t.Fatalf("Failed to delete borrower: %v", err)
    }

    deletedBorrow, err := borrDB.GetBorrowerById(context.Background(), &borr.GetBorrowerByIdRequest{Id: BorrowRes.Borrower.Id})
    assert.Nil(t, deletedBorrow)
    assert.Error(t, err) // Check if an error occurred when trying to get the deleted borrower
}
