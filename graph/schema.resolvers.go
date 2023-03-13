package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mgeale/homeserver/graph/generated"
	"github.com/mgeale/homeserver/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateBalance is the resolver for the createBalance field.
func (r *mutationResolver) InsertBalances(ctx context.Context, input []*model.InsertBalance) ([]string, error) {
	ids, err := r.app.CreateBalance(ctx, input)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	return ids, nil
}

// CreateTransaction is the resolver for the createTransaction field.
func (r *mutationResolver) InsertTransactions(ctx context.Context, input []*model.InsertTransaction) ([]string, error) {
	ids, err := r.app.CreateTransaction(ctx, input)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	return ids, nil
}

// UpdateBalance is the resolver for the updateBalance field.
func (r *mutationResolver) UpdateBalances(ctx context.Context, input []*model.UpdateBalance) (string, error) {
	err := r.app.UpdateBalance(ctx, input)
	if err != nil {
		return "Error", gqlerror.Errorf(err.Error())
	}
	return "OK", nil
}

// UpdateTransaction is the resolver for the updateTransaction field.
func (r *mutationResolver) UpdateTransactions(ctx context.Context, input []*model.UpdateTransaction) (string, error) {
	err := r.app.UpdateTransaction(ctx, input)
	if err != nil {
		return "Error", gqlerror.Errorf(err.Error())
	}
	return "OK", nil
}

// DeleteBalance is the resolver for the deleteBalance field.
func (r *mutationResolver) DeleteBalances(ctx context.Context, ids []string) (string, error) {
	err := r.app.DeleteBalance(ctx, ids)
	if err != nil {
		return "Error", gqlerror.Errorf(err.Error())
	}
	return "OK", nil
}

// DeleteTransaction is the resolver for the deleteTransaction field.
func (r *mutationResolver) DeleteTransactions(ctx context.Context, ids []string) (string, error) {
	err := r.app.DeleteTransaction(ctx, ids)
	if err != nil {
		return "Error", gqlerror.Errorf(err.Error())
	}
	return "OK", nil
}

// Balances is the resolver for the balances field.
func (r *queryResolver) Balances(ctx context.Context, where *model.BalanceFilter, orderBy model.BalanceSort, limit *int) ([]*model.Balance, error) {
	balances, err := r.app.GetBalances(ctx, where, orderBy, limit)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return balances, nil
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context, where *model.TransactionFilter, orderBy model.TransactionSort, limit *int) ([]*model.Transaction, error) {
	transactions, err := r.app.GetTransactions(ctx, where, orderBy, limit)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return transactions, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
