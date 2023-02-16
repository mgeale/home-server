package types

import (
	"github.com/mgeale/homeserver/graph/model"
	"github.com/mgeale/homeserver/internal/db"
)

func CreateBalanceQuery(where *model.BalanceFilter, orderBy model.BalanceSort, limit *int) *db.Query {
	filter := &db.Filter{}
	if where != nil {
		filter = toFilter(where)
	}
	return &db.Query{
		Filters: filter,
		Sort: db.Sort{
			Field:     db.Field(orderBy.Field),
			Direction: db.SortDirection(orderBy.Direction),
		},
		Limit: uint(*limit),
	}
}

func CreateTransactionQuery(where *model.TransactionFilter, orderBy model.TransactionSort, limit *int) *db.Query {
	filter := &db.Filter{}
	if where != nil {
		filter = toTransactionFilter(where)
	}
	return &db.Query{
		Filters: filter,
		Sort: db.Sort{
			Field:     db.Field(orderBy.Field),
			Direction: db.SortDirection(orderBy.Direction),
		},
		Limit: uint(*limit),
	}
}

func toFilter(where *model.BalanceFilter) *db.Filter {
	filter := &db.Filter{}
	if len(where.Subfilters) > 0 {
		subfilters := make([]*db.Filter, len(where.Subfilters))
		for i, f := range where.Subfilters {
			subfilters[i] = toFilter(f)
		}
		filter.Subfilters = subfilters
		filter.Kind = db.FilterKind(where.Kind)
	} else {
		filter.Field = db.Field(*where.Field)
		filter.Kind = db.FilterKind(where.Kind)
		filter.Value = where.Value
	}
	return filter
}

func toTransactionFilter(where *model.TransactionFilter) *db.Filter {
	filter := &db.Filter{}
	if len(where.Subfilters) > 0 {
		subfilters := make([]*db.Filter, len(where.Subfilters))
		for i, f := range where.Subfilters {
			subfilters[i] = toTransactionFilter(f)
		}
		filter.Subfilters = subfilters
		filter.Kind = db.FilterKind(where.Kind)
	} else {
		filter.Field = db.Field(*where.Field)
		filter.Kind = db.FilterKind(where.Kind)
		filter.Value = where.Value
	}
	return filter
}

func ToBalanceModel(balances []*db.Balance) []*model.Balance {
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
	return result
}

func ToTransactionModel(transactions []*db.Transaction) []*model.Transaction {
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
	return result
}
