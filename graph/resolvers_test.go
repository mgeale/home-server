package graph

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mgeale/homeserver/graph/generated"
	"github.com/mgeale/homeserver/internal/app"
	"github.com/mgeale/homeserver/internal/db"
	"github.com/mgeale/homeserver/internal/db/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBalance(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver())))

	t.Run("Balances insert mutation", func(t *testing.T) {
		var resp map[string][]string

		c.MustPost(`
			mutation InsertBalances {
				insertBalances(input: [{
					name: "BAL-6789", 
					balance: 250.33, 
					balanceaud: 5680.66, 
					productid: "222", 
					pricebookid: "444"
				}])
			}`, &resp)

		assert.True(t, len(resp["insertBalances"]) > 0, "Should insert Balance records and return ids.")
	})
}

func TestUpdateBalance(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver())))

	t.Run("Balances update mutation", func(t *testing.T) {
		var resp map[string]string

		c.MustPost(`
			mutation UpdateBalances {
				updateBalances(input: [{
					ExternalId: "5555",
					name: "BAL-6789", 
					balance: 250.33, 
					balanceaud: 5680.66, 
					productid: "222", 
					pricebookid: "444"
				}])
			}`, &resp)

		assert.Equal(t, "OK", resp["updateBalances"], "Should update Balance record.")
	})
}

func TestDeleteBalance(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver())))

	t.Run("Balances delete mutation", func(t *testing.T) {
		var resp map[string]string

		c.MustPost(`
			mutation DeleteBalances {
				deleteBalances(ids: ["5555"])
			}`, &resp)

		assert.Equal(t, "OK", resp["deleteBalances"], "Should delete Balance records.")
	})
}

func TestGetBalances(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver())))

	t.Run("Balances get query", func(t *testing.T) {
		var resp struct {
			Balances []struct {
				ExternalId  string
				Name        string
				Balance     float64
				Balanceaud  float64
				Pricebookid string
				Productid   string
				Created     string
			}
		}

		c.MustPost(`
			query GetBalances {
				balances(where: {
					field: ExternalId
					kind: EQUALS
					value: "7a59f5c1-b0b9-11ed-a356-0242ac110002"
				}, orderBy: {
					field: created
					direction: DESC
				}) {
					ExternalId 
					name 
					balance 
					balanceaud 
					pricebookid 
					productid 
					created
				}
			}`, &resp)

		assert.True(t, len(resp.Balances) > 0, "Should return list of Balances.")
	})
}

func TestCreateTransaction(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver())))

	t.Run("Transaction insert mutation", func(t *testing.T) {
		var resp map[string][]string

		c.MustPost(`
			mutation InsertTransactions {
				insertTransactions(input: {
					name: "BAL-6789", 
					amount: 220.22
					date: "2018-12-23 17:25:22",
					type: "Repayment"
				})
			}`, &resp)

		assert.True(t, len(resp["insertTransactions"]) > 0, "Should insert Transaction records and return ids.")
	})
}

func TestUpdateTransaction(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver())))

	t.Run("Transaction update mutation", func(t *testing.T) {
		var resp map[string]string

		c.MustPost(`
			mutation UpdateTransactions {
				updateTransactions(input: [{
					ExternalId: "5555",
					DisplayUrl: "",
					name: "BAL-6789", 
					amount: 220.22
					date: "2018-12-23 17:25:22",
					type: "Repayment"
				}])
			}`, &resp)

		assert.Equal(t, "OK", resp["updateTransactions"], "Should update Transaction records.")
	})
}

func TestDeleteTransaction(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver())))

	t.Run("Transactions delete mutation", func(t *testing.T) {
		var resp map[string]string

		c.MustPost(`
			mutation DeleteTransactions {
				deleteTransactions(ids: "5555")
			}`, &resp)

		assert.Equal(t, "OK", resp["deleteTransactions"], "Should delete Transaction records.")
	})
}

func TestGetTransactions(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(newResolver())))

	t.Run("Transactions get query", func(t *testing.T) {
		var resp struct {
			Transactions []struct {
				ExternalId string
				Name       string
				Amount     float64
				Date       string
				Type       string
				Created    string
			}
		}

		c.MustPost(`
			query GetTransactions {
				transactions(where: {
					field: ExternalId
					kind: EQUALS
					value: "2"
				}, orderBy: {
					field: created
					direction: DESC
				}) {
					ExternalId 
					name 
					amount 
					date 
					type  
					created
				}
			}`, &resp)

		assert.True(t, len(resp.Transactions) > 0, "Should return list of Transactions.")
	})
}

func newResolver() generated.Config {
	app := &app.Application{
		Models: db.Models{
			Balances:     &mock.BalanceModel{},
			Transactions: &mock.TransactionModel{},
			Users:        &mock.UserModel{},
		},
	}
	return generated.Config{
		Resolvers: &Resolver{app},
	}
}
