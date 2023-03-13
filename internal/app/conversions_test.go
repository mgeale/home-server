package app

import (
	"testing"

	"github.com/mgeale/homeserver/graph/model"
	"github.com/mgeale/homeserver/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuery(t *testing.T) {
	t.Run("create query", func(t *testing.T) {
		val := "2"
		field := model.BalanceFieldExternalID
		fieldBalance := model.BalanceFieldBalance
		subfilters := []*model.BalanceFilter{
			&model.BalanceFilter{
				Field: &field,
				Kind:  model.FilterKindEquals,
				Value: &val,
			},
			&model.BalanceFilter{
				Field: &fieldBalance,
				Kind:  model.FilterKindNotEquals,
				Value: &val,
			},
		}

		where := &model.BalanceFilter{
			Subfilters: subfilters,
			Kind:       model.FilterKindOr,
		}
		orderBy := model.BalanceSort{}
		limit := 100

		result := createBalanceQuery(where, orderBy, &limit)
		resultFilters := []*db.Filter{
			&db.Filter{
				Field: db.Field("id"),
				Kind:  db.FilterKind("EQUALS"),
				Value: &val,
			},
			&db.Filter{
				Field: db.Field("balance"),
				Kind:  db.FilterKind("NOT_EQUALS"),
				Value: &val,
			},
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
