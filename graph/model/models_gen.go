// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Balance struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Balance     float64 `json:"balance"`
	Balanceaud  float64 `json:"balanceaud"`
	Pricebookid int     `json:"pricebookid"`
	Productid   int     `json:"productid"`
	Created     string  `json:"created"`
}

type BalanceFilter struct {
	Field BalanceField `json:"field"`
	Kind  FilterKind   `json:"kind"`
	Value string       `json:"value"`
}

type BalanceSort struct {
	Field     BalanceField  `json:"field"`
	Direction SortDirection `json:"direction"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewBalance struct {
	Name        string  `json:"name"`
	Balance     float64 `json:"balance"`
	Balanceaud  float64 `json:"balanceaud"`
	Pricebookid int     `json:"pricebookid"`
	Productid   int     `json:"productid"`
}

type NewTransaction struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	Type   string  `json:"type"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Transaction struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`
	Date    string  `json:"date"`
	Type    string  `json:"type"`
	Created string  `json:"created"`
}

type TransactionFilter struct {
	Field TransactionField `json:"field"`
	Kind  FilterKind       `json:"kind"`
	Value string           `json:"value"`
}

type TransactionSort struct {
	Field     TransactionField `json:"field"`
	Direction SortDirection    `json:"direction"`
}

type UpdateBalance struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Balance     float64 `json:"balance"`
	Balanceaud  float64 `json:"balanceaud"`
	Pricebookid int     `json:"pricebookid"`
	Productid   int     `json:"productid"`
}

type UpdateTransaction struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	Type   string  `json:"type"`
}

type BalanceField string

const (
	BalanceFieldID          BalanceField = "id"
	BalanceFieldName        BalanceField = "name"
	BalanceFieldBalance     BalanceField = "balance"
	BalanceFieldBalanceaud  BalanceField = "balanceaud"
	BalanceFieldPricebookid BalanceField = "pricebookid"
	BalanceFieldProductid   BalanceField = "productid"
	BalanceFieldCreated     BalanceField = "created"
)

var AllBalanceField = []BalanceField{
	BalanceFieldID,
	BalanceFieldName,
	BalanceFieldBalance,
	BalanceFieldBalanceaud,
	BalanceFieldPricebookid,
	BalanceFieldProductid,
	BalanceFieldCreated,
}

func (e BalanceField) IsValid() bool {
	switch e {
	case BalanceFieldID, BalanceFieldName, BalanceFieldBalance, BalanceFieldBalanceaud, BalanceFieldPricebookid, BalanceFieldProductid, BalanceFieldCreated:
		return true
	}
	return false
}

func (e BalanceField) String() string {
	return string(e)
}

func (e *BalanceField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BalanceField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BalanceField", str)
	}
	return nil
}

func (e BalanceField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FilterKind string

const (
	FilterKindEqual          FilterKind = "EQUAL"
	FilterKindNotEqual       FilterKind = "NOT_EQUAL"
	FilterKindGreater        FilterKind = "GREATER"
	FilterKindGreaterOrEqual FilterKind = "GREATER_OR_EQUAL"
	FilterKindLess           FilterKind = "LESS"
	FilterKindLessOrEqual    FilterKind = "LESS_OR_EQUAL"
)

var AllFilterKind = []FilterKind{
	FilterKindEqual,
	FilterKindNotEqual,
	FilterKindGreater,
	FilterKindGreaterOrEqual,
	FilterKindLess,
	FilterKindLessOrEqual,
}

func (e FilterKind) IsValid() bool {
	switch e {
	case FilterKindEqual, FilterKindNotEqual, FilterKindGreater, FilterKindGreaterOrEqual, FilterKindLess, FilterKindLessOrEqual:
		return true
	}
	return false
}

func (e FilterKind) String() string {
	return string(e)
}

func (e *FilterKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterKind", str)
	}
	return nil
}

func (e FilterKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

var AllSortDirection = []SortDirection{
	SortDirectionAsc,
	SortDirectionDesc,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TransactionField string

const (
	TransactionFieldID          TransactionField = "id"
	TransactionFieldName        TransactionField = "name"
	TransactionFieldBalance     TransactionField = "balance"
	TransactionFieldBalanceaud  TransactionField = "balanceaud"
	TransactionFieldPricebookid TransactionField = "pricebookid"
	TransactionFieldProductid   TransactionField = "productid"
	TransactionFieldCreated     TransactionField = "created"
)

var AllTransactionField = []TransactionField{
	TransactionFieldID,
	TransactionFieldName,
	TransactionFieldBalance,
	TransactionFieldBalanceaud,
	TransactionFieldPricebookid,
	TransactionFieldProductid,
	TransactionFieldCreated,
}

func (e TransactionField) IsValid() bool {
	switch e {
	case TransactionFieldID, TransactionFieldName, TransactionFieldBalance, TransactionFieldBalanceaud, TransactionFieldPricebookid, TransactionFieldProductid, TransactionFieldCreated:
		return true
	}
	return false
}

func (e TransactionField) String() string {
	return string(e)
}

func (e *TransactionField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionField", str)
	}
	return nil
}

func (e TransactionField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
