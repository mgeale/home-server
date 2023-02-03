package graph

import (
	"database/sql"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mgeale/homeserver/graph/generated"
	"github.com/mgeale/homeserver/internal/app"
	"github.com/mgeale/homeserver/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateBalance(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec("^INSERT INTO balances (.+)").WillReturnResult(sqlmock.NewResult(5555, 1))

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver(database))))

	t.Run("Balances mutation", func(t *testing.T) {
		var resp map[string]int

		c.MustPost(`
			mutation CreateBalance {
				createBalance(input: {
					name: "BAL-6789", 
					balance: 250.33, 
					balanceaud: 5680.66, 
					productid: 222, 
					pricebookid: 444
				})
			}`, &resp)

		assert.Equal(t, 5555, resp["createBalance"], "Should create new Balance record and return id.")
	})
}

func TestUpdateBalance(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec("^UPDATE balances (.+)").WillReturnResult(sqlmock.NewResult(5555, 1))

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver(database))))

	t.Run("Balances mutation", func(t *testing.T) {
		var resp map[string]int

		c.MustPost(`
			mutation UpdateBalance {
				updateBalance(input: {
					id: 5555,
					name: "BAL-6789", 
					balance: 250.33, 
					balanceaud: 5680.66, 
					productid: 222, 
					pricebookid: 444
				})
			}`, &resp)

		assert.Equal(t, 5555, resp["updateBalance"], "Should update Balance record and return id.")
	})
}

func TestDeleteBalance(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec("^DELETE FROM balances (.+)").WillReturnResult(sqlmock.NewResult(5555, 1))

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver(database))))

	t.Run("Balances mutation", func(t *testing.T) {
		var resp map[string]int

		c.MustPost(`
			mutation DeleteBalance {
				deleteBalance(id: 5555)
			}`, &resp)

		assert.Equal(t, 1, resp["deleteBalance"], "Should delete Balance record and return id.")
	})
}

func TestGetBalances(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "balance", "balanceaud", "pricebookid", "productid", "created"}).
		AddRow(1, "BAL-0001", 200.00, 400.00, 2222, 3333, time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)).
		AddRow(2, "BAL-0002", 1000.10, 2000.20, 2222, 3333, time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC))

	mock.ExpectQuery("^SELECT (.+) FROM balances (.+)").WillReturnRows(rows)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver(database))))

	t.Run("Balances query", func(t *testing.T) {
		var resp struct {
			Balances []struct {
				ID          string
				Name        string
				Balance     float64
				Balanceaud  float64
				Pricebookid float64
				Productid   float64
				Created     string
			}
		}

		c.MustPost(`
			query GetBalances {
				balances(where: {
					field: id
					kind: EQUAL
					value: "2"
				}, orderBy: {
					field: created
					direction: DESC
				}) {
					id 
					name 
					balance 
					balanceaud 
					pricebookid 
					productid 
					created
				}
			}`, &resp)

		assert.True(t, len(resp.Balances) > 0, "Should return list of Balances greater than 0.")
	})
}

func TestGetBalanceByID(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "balance", "balanceaud", "pricebookid", "productid", "created"}).
		AddRow(1, "BAL-0001", 200.00, 400.00, 2222, 3333, time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC))

	mock.ExpectQuery("^SELECT (.+) FROM balances WHERE id = (.+)").WillReturnRows(rows)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver(database))))

	t.Run("Balances query", func(t *testing.T) {
		var resp map[string]struct {
			ID          string
			Name        string
			Balance     float64
			Balanceaud  float64
			Pricebookid float64
			Productid   float64
			Created     string
		}

		c.MustPost(`
			query BalanceByID {
				balanceById(id: "1") {
					id 
					name 
					balance 
					balanceaud 
					pricebookid 
					productid 
					created
				}
			}`, &resp)

		assert.True(t, resp["balanceById"].ID == "1", "Should return Balance.")
	})
}

func TestCreateTransaction(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec("^INSERT INTO transactions (.+)").WillReturnResult(sqlmock.NewResult(5555, 1))

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver(database))))

	t.Run("Transaction update mutation", func(t *testing.T) {
		var resp map[string]int

		c.MustPost(`
			mutation CreateTransaction {
				createTransaction(input: {
					name: "BAL-6789", 
					amount: 220.22
					date: "2018-12-23 17:25:22",
					type: "Repayment"
				})
			}`, &resp)

		assert.Equal(t, 5555, resp["createTransaction"], "Should create new Transaction record and return id.")
	})
}

func TestUpdateTransaction(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec("^UPDATE transactions (.+)").WillReturnResult(sqlmock.NewResult(5555, 1))

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver(database))))

	t.Run("Transaction update mutation", func(t *testing.T) {
		var resp map[string]int

		c.MustPost(`
			mutation UpdateTransaction  {
				updateTransaction(input: {
					id: 5555,
					name: "BAL-6789", 
					amount: 220.22
					date: "2018-12-23 17:25:22",
					type: "Repayment"
				})
			}`, &resp)

		assert.Equal(t, 5555, resp["updateTransaction"], "Should update Transaction record and return id.")
	})
}

func TestDeleteTransaction(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec("^DELETE FROM transactions (.+)").WillReturnResult(sqlmock.NewResult(5555, 1))

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver(database))))

	t.Run("Transactions delete mutation", func(t *testing.T) {
		var resp map[string]int

		c.MustPost(`
			mutation DeleteTransactions {
				deleteTransaction(id: 5555)
			}`, &resp)

		assert.Equal(t, 1, resp["deleteTransaction"], "Should delete Transaction record and return id.")
	})
}

func TestGetTransactions(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "amount", "date", "type", "created"}).
		AddRow(1, "TRANS-0001", 200.00, "2018-12-23 17:25:22", "Repayment", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)).
		AddRow(1, "TRANS-0002", 400.00, "2018-12-23 17:25:22", "Repayment", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC))

	mock.ExpectQuery("^SELECT (.+) FROM transactions (.+)").WillReturnRows(rows)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver(database))))

	t.Run("Transactions get all query", func(t *testing.T) {
		var resp struct {
			Transactions []struct {
				ID      string
				Name    string
				Amount  float64
				Date    string
				Type    string
				Created string
			}
		}

		c.MustPost(`
			query GetTransactions {
				transactions(where: {
					field: id
					kind: EQUAL
					value: "2"
				}, orderBy: {
					field: created
					direction: DESC
				}) {
					id 
					name 
					amount 
					date 
					type  
					created
				}
			}`, &resp)

		assert.True(t, len(resp.Transactions) > 0, "Should return list of Transactions greater than 0.")
	})
}

func TestGetTransactionByID(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "amount", "date", "type", "created"}).
		AddRow(1, "TRANS-0001", 200.00, "2018-12-23 17:25:22", "Repayment", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC))

	mock.ExpectQuery("^SELECT (.+) FROM transactions WHERE id = (.+)").WillReturnRows(rows)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver(database))))

	t.Run("Transaction get by id query", func(t *testing.T) {
		var resp map[string]struct {
			ID      string
			Name    string
			Amount  float64
			Date    string
			Type    string
			Created string
		}

		c.MustPost(`
			query TransactionByID {
				transactionById(id: "1") {
					id 
					name 
					amount 
					date 
					type  
					created
				}
			}`, &resp)

		assert.True(t, resp["transactionById"].ID == "1", "Should return Transaction.")
	})
}

func newResolver(database *sql.DB) generated.Config {
	app := &app.Application{
		Models: db.NewModels(database),
	}
	return generated.Config{
		Resolvers: &Resolver{app},
	}
}
