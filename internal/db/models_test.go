package db

import (
	"reflect"
	"testing"

	"github.com/ido50/sqlz"
)

func TestWhereConditions(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name           string
		filterOpts     *Filter
		wantConditions sqlz.WhereCondition
		wantError      error
	}{
		{
			name: "Single",
			filterOpts: &Filter{
				Field: Field("id"),
				Kind:  FilterKind("EQUAL"),
				Value: "2",
			},
			wantConditions: sqlz.Eq("id", "2"),
			wantError:      nil,
		},
		{
			name: "Nested",
			filterOpts: &Filter{
				Subfilters: []*Filter{
					{
						Field: Field("id"),
						Kind:  FilterKind("EQUAL"),
						Value: "2",
					},
					{
						Field: Field("id"),
						Kind:  FilterKind("EQUAL"),
						Value: "3",
					},
				},
				Kind: FilterKind("OR_"),
			},
			wantConditions: sqlz.Or(sqlz.Eq("id", "2"), sqlz.Eq("id", "3")),
			wantError:      nil,
		},
		{
			name: "Double Nested",
			filterOpts: &Filter{
				Subfilters: []*Filter{
					{
						Field: Field("id"),
						Kind:  FilterKind("EQUAL"),
						Value: "2",
					},
					{
						Subfilters: []*Filter{
							{
								Field: Field("balance"),
								Kind:  FilterKind("GREATER"),
								Value: 100,
							},
							{
								Field: Field("balance"),
								Kind:  FilterKind("LESS"),
								Value: 200,
							},
						},
						Kind: FilterKind("AND_"),
					},
				},
				Kind: FilterKind("OR_"),
			},
			wantConditions: sqlz.Or(sqlz.Eq("id", "2"), sqlz.And(sqlz.Gt("balance", 100), sqlz.Lt("balance", 200))),
			wantError:      nil,
		},
		{
			name: "Triple Nested",
			filterOpts: &Filter{
				Subfilters: []*Filter{
					{
						Field: Field("id"),
						Kind:  FilterKind("EQUAL"),
						Value: "2",
					},
					{
						Subfilters: []*Filter{
							{
								Field: Field("balance"),
								Kind:  FilterKind("GREATER"),
								Value: 100,
							},
							{
								Subfilters: []*Filter{
									{
										Field: Field("balanceaud"),
										Kind:  FilterKind("GREATER_OR_EQUAL"),
										Value: 250,
									},
									{
										Field: Field("balanceaud"),
										Kind:  FilterKind("LESS_OR_EQUAL"),
										Value: 500,
									},
								},
								Kind: FilterKind("AND_"),
							},
						},
						Kind: FilterKind("AND_"),
					},
				},
				Kind: FilterKind("OR_"),
			},
			wantConditions: sqlz.Or(sqlz.Eq("id", "2"), sqlz.And(sqlz.Gt("balance", 100), sqlz.And(sqlz.Gte("balanceaud", 250), sqlz.Lte("balanceaud", 500)))),
			wantError:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			whereConditions := whereConditionsFromOrderFilterOpts(tt.filterOpts)

			if !reflect.DeepEqual(whereConditions, tt.wantConditions) {
				t.Errorf("want %v; got %v", tt.wantConditions, whereConditions)
			}
		})
	}
}
