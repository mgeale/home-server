package types

import (
	"testing"

	"github.com/mgeale/homeserver/graph/model"
	"github.com/mgeale/homeserver/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuery(t *testing.T) {
	t.Run("create query", func(t *testing.T) {
		val := "2"
		field := model.BalanceFieldID
		fieldBalance := model.BalanceFieldBalance
		subfilters := make([]*model.BalanceFilter, 2)
		subfilters[0] = &model.BalanceFilter{
			Field: &field,
			Kind:  model.FilterKindEqual,
			Value: &val,
		}
		subfilters[1] = &model.BalanceFilter{
			Field: &fieldBalance,
			Kind:  model.FilterKindNotEqual,
			Value: &val,
		}

		where := &model.BalanceFilter{
			Subfilters: subfilters,
			Kind:       model.FilterKindOr,
		}
		orderBy := model.BalanceSort{}
		limit := 100

		result := CreateBalanceQuery(where, orderBy, &limit)
		resultFilters := make([]*db.Filter, 2)
		resultFilters[0] = &db.Filter{
			Field: db.Field("id"),
			Kind:  db.FilterKind("EQUAL"),
			Value: &val,
		}
		resultFilters[1] = &db.Filter{
			Field: db.Field("balance"),
			Kind:  db.FilterKind("NOT_EQUAL"),
			Value: &val,
		}
		expected := &db.Query{
			Filters: &db.Filter{
				Subfilters: resultFilters,
				Kind:       db.FilterKind("OR_"),
			},
			Limit: 100,
		}
		if len(result.Filters.Subfilters) != 2 {
			t.Errorf("want %d; got %d", 2, len(result.Filters.Subfilters))
		}
		assert.Equal(t, expected, result, "Should return query struct.")
	})
}
