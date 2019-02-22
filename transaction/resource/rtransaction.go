package resource

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/tumpal2796/sophee/transaction/model"
)

const (
	InsertQuery = `
		INSERT INTO transaction(
			name,
			tax_code,
			price
		) VALUES (
			:name,
			:tax_code,
			:price
		)`

	SelectQuery = `
		SELECT
			name,
			tax_code,
			price	
		FROM
			transaction`
)

type TransactionInf interface {
	Insert(ctx context.Context, transaction model.Transaction) error
	GetAllTransaction(ctx context.Context) ([]model.Transaction, error)
}

type TransactionImpl struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) TransactionInf {
	return &TransactionImpl{
		DB: db,
	}
}

func (transImpl *TransactionImpl) Insert(ctx context.Context, transaction model.Transaction) error {
	_, err := transImpl.DB.NamedExecContext(ctx, InsertQuery, transaction)
	if err != nil {
		return err
	}

	return err
}

func (transImpl *TransactionImpl) GetAllTransaction(ctx context.Context) ([]model.Transaction, error) {
	var Result []model.Transaction
	err := transImpl.DB.Select(&Result, SelectQuery)
	if err != nil {
		return Result, err
	}

	for i := 0; i < len(Result); i++ {
		Result[i].Type = getType(Result[i].TaxCode)
	}

	return Result, err
}

func getType(taxCode int8) string {
	tipe := make(map[int8]string)
	tipe[1] = "food"
	tipe[2] = "tobacco"
	tipe[3] = "entertainment"

	return tipe[taxCode]
}
