package resource

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/tumpal2796/sophee/transaction/model"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetType(t *testing.T) {
	assert := assert.New(t)

	type TestCase struct {
		taxCode        int8
		expectedResult string
	}

	testCases := make(map[string]TestCase)
	testCases["get food type"] = TestCase{
		taxCode:        1,
		expectedResult: "food",
	}

	testCases["get  tobacco type"] = TestCase{
		taxCode:        2,
		expectedResult: "tobacco",
	}

	testCases["get entertainment type"] = TestCase{
		taxCode:        3,
		expectedResult: "entertainment",
	}

	for key, val := range testCases {
		result := getType(val.taxCode)
		assert.Equal(val.expectedResult, result, "[TestGetType] Failed on Test %s", key)
	}

	t.Logf("[TestGetType] Success")
	return
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("[TestInsert] Failed")
		return
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rTransaction := New(sqlxDB)

	var param model.Transaction
	param.Name = "oreo"
	param.Price = 1000
	param.TaxCode = 1

	mock.ExpectExec("INSERT (.+)").WithArgs(param.Name, param.TaxCode, param.Price).WillReturnResult(sqlmock.NewResult(1, 1))
	err = rTransaction.Insert(ctx, param)
	assert.NoError(err, "[TestInsert] Failed")
	err = mock.ExpectationsWereMet()
	assert.NoError(err, "[TestInsert] Failed")

	t.Logf("[TestInsert] Success")
	return
}

func TestErrorInsert(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("[TestErrorInsert] Failed")
		return
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rTransaction := New(sqlxDB)

	var param model.Transaction
	param.Name = "oreo"
	param.Price = 1000
	param.TaxCode = 1

	mock.ExpectExec("INSERT (.+)").WithArgs(param.Name, param.TaxCode, param.Price).WillReturnError(errors.New("Failed to Insert"))
	err = rTransaction.Insert(ctx, param)
	assert.Error(err, "[TestErrorInsert] Failed")
	err = mock.ExpectationsWereMet()
	assert.NoError(err, "[TestErrorInsert] Failed")

	t.Logf("[TestErrorInsert] Success")
	return
}

func TestGetAllTransaction(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("[TestErrorInsert] Failed")
		return
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rTransaction := New(sqlxDB)

	expectedResult := []model.Transaction{
		model.Transaction{
			Name:    "oreo",
			TaxCode: 1,
			Type:    "food",
			Price:   10000,
		},
	}

	rows := sqlmock.NewRows([]string{
		"name",
		"tax_code",
		"price",
	}).AddRow(
		"oreo",
		1,
		10000,
	)

	mock.ExpectQuery("SELECT (.+)").WillReturnRows(rows)
	result, err := rTransaction.GetAllTransaction(ctx)
	assert.NoError(err, "[TestGetAllTransaction] Failed")
	assert.Equal(expectedResult, result, "[TestGetAllTransaction] Failed")
	err = mock.ExpectationsWereMet()
	assert.NoError(err, "[TestGetAllTransaction] Failed")

	t.Logf("[TestGetAllTransaction] Success")
	return
}

func TestErrorGetAllTransaction(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("[TestErrorInsert] Failed")
		return
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	rTransaction := New(sqlxDB)

	var expectedResult []model.Transaction

	mock.ExpectQuery("SELECT (.+)").WillReturnError(errors.New("Failed Get All Transaction"))
	result, err := rTransaction.GetAllTransaction(ctx)
	assert.Error(err, "[TestErrorGetAllTransaction] Failed")
	assert.Equal(expectedResult, result, "[TestErrorGetAllTransaction] Failed")
	err = mock.ExpectationsWereMet()
	assert.NoError(err, "[TestErrorGetAllTransaction] Failed")

	t.Logf("[TestErrorGetAllTransaction] Success")
	return
}
