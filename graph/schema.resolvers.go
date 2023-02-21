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
func (r *mutationResolver) CreateBalance(ctx context.Context, input model.NewBalance) (string, error) {
	id, err := r.app.CreateBalance(ctx, &input)
	if err != nil {
		return "0", gqlerror.Errorf(err.Error())
	}
	return id, nil
}

// CreateTransaction is the resolver for the createTransaction field.
func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.NewTransaction) (string, error) {
	id, err := r.app.CreateTransaction(ctx, &input)
	if err != nil {
		return "0", gqlerror.Errorf(err.Error())
	}
	return id, nil
}

// UpdateBalance is the resolver for the updateBalance field.
func (r *mutationResolver) UpdateBalance(ctx context.Context, input model.UpdateBalance) (string, error) {
	id, err := r.app.UpdateBalance(ctx, &input)
	if err != nil {
		return "0", gqlerror.Errorf(err.Error())
	}
	return id, nil
}

// UpdateTransaction is the resolver for the updateTransaction field.
func (r *mutationResolver) UpdateTransaction(ctx context.Context, input model.UpdateTransaction) (string, error) {
	id, err := r.app.UpdateTransaction(ctx, &input)
	if err != nil {
		return "0", gqlerror.Errorf(err.Error())
	}
	return id, nil
}

// DeleteBalance is the resolver for the deleteBalance field.
func (r *mutationResolver) DeleteBalance(ctx context.Context, id string) (string, error) {
	_, err := r.app.DeleteBalance(ctx, id)
	if err != nil {
		return "0", gqlerror.Errorf(err.Error())
	}
	return "1", nil
}

// DeleteTransaction is the resolver for the deleteTransaction field.
func (r *mutationResolver) DeleteTransaction(ctx context.Context, id string) (string, error) {
	_, err := r.app.DeleteTransaction(ctx, id)
	if err != nil {
		return "0", gqlerror.Errorf(err.Error())
	}
	return "1", nil
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
