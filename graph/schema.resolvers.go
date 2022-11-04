package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mgeale/homeserver/graph/generated"
	"github.com/mgeale/homeserver/graph/model"
)

// Balances is the resolver for the balances field.
func (r *queryResolver) Balances(ctx context.Context) ([]*model.Balance, error) {
	var res []*model.Balance

	res = append(res, &model.Balance{
		ID:          "30000",
		Name:        "Name",
		Balance:     230.45,
		Balanceaud:  255.33,
		Pricebookid: "4565563",
		Productid:   "3465465",
		Created:     "today",
	})

	return res, nil
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented: Transactions - transactions"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
