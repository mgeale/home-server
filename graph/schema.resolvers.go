package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mgeale/homeserver/graph/generated"
	"github.com/mgeale/homeserver/graph/model"
	"github.com/mgeale/homeserver/internal/db"
)

// CreateBalance is the resolver for the createBalance field.
func (r *mutationResolver) CreateBalance(ctx context.Context, input model.NewBalance) (int, error) {
	var b *db.Balance

	b.Name = input.Name
	b.Balance = input.Balance
	b.BalanceAUD = input.Balanceaud
	b.PricebookID = input.Pricebookid
	b.ProductID = input.Productid

	id, err := r.app.CreateBalance(ctx, b)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// CreateTransaction is the resolver for the createTransaction field.
func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.NewTransaction) (int, error) {
	var t *db.Transaction

	t.Name = input.Name
	t.Amount = input.Amount
	t.Date = input.Date
	t.Type = input.Type

	id, err := r.app.CreateTransaction(ctx, t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *mutationResolver) UpdateBalance(ctx context.Context, input model.UpdateBalance) (*model.Balance, error) {
	panic(fmt.Errorf("not implemented: UpdateBalance - updateBalance"))
}

// UpdateTransaction is the resolver for the updateTransaction field.
func (r *mutationResolver) UpdateTransaction(ctx context.Context, input model.UpdateTransaction) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented: UpdateTransaction - updateTransaction"))
}

// DeleteBalance is the resolver for the deleteBalance field.
func (r *mutationResolver) DeleteBalance(ctx context.Context, id int) (int, error) {
	_, err := r.app.DeleteBalance(ctx, id)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

// DeleteTransaction is the resolver for the deleteTransaction field.
func (r *mutationResolver) DeleteTransaction(ctx context.Context, id int) (int, error) {
	_, err := r.app.DeleteTransaction(ctx, id)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

// Balances is the resolver for the balances field.
func (r *queryResolver) Balances(ctx context.Context) ([]*model.Balance, error) {
	balances, err := r.app.GetLatestBalances()
	if err != nil {
		return nil, err
	}

	result := make([]*model.Balance, len(balances))
	for i, b := range balances {
		result[i] = &model.Balance{
			ID:          b.ID,
			Name:        b.Name,
			Balance:     b.Balance,
			Balanceaud:  b.BalanceAUD,
			Pricebookid: b.PricebookID,
			Productid:   b.ProductID,
			Created:     b.Created.String(),
		}
	}
	return result, nil
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	transactions, err := r.app.GetLatestTransactions()
	if err != nil {
		return nil, err
	}

	result := make([]*model.Transaction, len(transactions))
	for i, t := range transactions {
		result[i] = &model.Transaction{
			ID:      t.ID,
			Name:    t.Name,
			Amount:  t.Amount,
			Date:    t.Date,
			Type:    t.Type,
			Created: t.Created.String(),
		}
	}

	return result, nil
}

// BalanceByID is the resolver for the balanceById field.
func (r *queryResolver) BalanceByID(ctx context.Context, id int) (*model.Balance, error) {
	b, err := r.app.Models.Balances.Get(id)
	if err != nil {
		return nil, err
	}

	return &model.Balance{
		ID:          b.ID,
		Name:        b.Name,
		Balance:     b.Balance,
		Balanceaud:  b.BalanceAUD,
		Pricebookid: b.PricebookID,
		Productid:   b.ProductID,
		Created:     b.Created.String(),
	}, nil
}

// TransactionByID is the resolver for the transactionById field.
func (r *queryResolver) TransactionByID(ctx context.Context, id int) (*model.Transaction, error) {
	t, err := r.app.Models.Transactions.Get(id)
	if err != nil {
		return nil, err
	}

	return &model.Transaction{
		ID:      t.ID,
		Name:    t.Name,
		Amount:  t.Amount,
		Date:    t.Date,
		Type:    t.Type,
		Created: t.Created.String(),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
