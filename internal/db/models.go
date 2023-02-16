package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ido50/sqlz"
)

var (
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrRecordNotFound     = errors.New("models: record not found")
	ErrEditConflict       = errors.New("models: edit conflict")
)

type SortDirection string

const (
	Ascending  SortDirection = "ASC"
	Descending SortDirection = "DESC"
)

type FilterKind string

const (
	Equal          FilterKind = "EQUAL"
	NotEqual       FilterKind = "NOT_EQUAL"
	Less           FilterKind = "LESS"
	Greater        FilterKind = "GREATER"
	LessOrEqual    FilterKind = "LESS_OR_EQUAL"
	GreaterOrEqual FilterKind = "GREATER_OR_EQUAL"
	Contains       FilterKind = "CONTAINS"
	And            FilterKind = "AND_"
	Or             FilterKind = "OR_"
	Not            FilterKind = "NOT_"
)

type Field string

type Query struct {
	Filters *Filter
	Sort    Sort
	Limit   uint
	Offset  uint
}

type Sort struct {
	Field     Field
	Direction SortDirection
}

type Filter struct {
	Field      Field
	Kind       FilterKind
	Value      interface{}
	Subfilters []*Filter
}

type Models struct {
	Balances interface {
		Delete(id int) error
		GetById(id int) (*Balance, error)
		Get(query *Query) ([]*Balance, error)
		Insert(name string, balance float64, balanceaud float64, pricebookid int, productid int) (int, error)
		Latest() ([]*Balance, error)
		Update(id int, name string, balance float64, balanceaud float64, pricebookid int, productid int) error
	}
	Transactions interface {
		Delete(id int) error
		GetById(id int) (*Transaction, error)
		Get(query *Query) ([]*Transaction, error)
		Insert(name string, amount float64, date string, transactionType string) (int, error)
		Latest() ([]*Transaction, error)
		Update(id int, name string, amount float64, date string, transactionType string) error
	}
	Users interface {
		GetByEmail(email string) (*User, error)
	}
	Permissions interface {
		AddForUser(userID int64, codes string) error
		GetAllForUser(userID int64) (Permissions, error)
	}
}

func NewModels(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		Balances: &BalanceModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Transactions: &TransactionModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Users: &UserModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Permissions: &PermissionModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
	}
}

func addOptsToSelectOrdersQuery(stmt *sqlz.SelectStmt, opts *Query) (*sqlz.SelectStmt, error) {
	if opts == nil {
		return stmt, nil
	}

	ordering := orderingFromOrderSortOpts(opts.Sort)
	if len(ordering) != 0 {
		stmt.OrderBy(ordering...)
	}
	if opts.Limit != 0 {
		stmt.Limit(int64(opts.Limit))
	}
	if opts.Offset != 0 {
		stmt.Offset(int64(opts.Offset))
	}

	whereConditions := whereConditionsFromOrderFilterOpts(opts.Filters)
	if whereConditions != nil {
		stmt.Where(whereConditions)
	}

	return stmt, nil
}

func orderingFromOrderSortOpts(sortOpt Sort) []sqlz.SQLStmt {
	ordering := []sqlz.SQLStmt{}
	if sortOpt.Direction == Ascending {
		ordering = append(ordering, sqlz.Asc(string(sortOpt.Field)))
	} else {
		ordering = append(ordering, sqlz.Desc(string(sortOpt.Field)))
	}
	return ordering
}

func whereConditionsFromOrderFilterOpts(filterOpt *Filter) sqlz.WhereCondition {
	switch filterOpt.Kind {
	case Equal:
		return sqlz.Eq(string(filterOpt.Field), filterOpt.Value)
	case NotEqual:
		return sqlz.Not(sqlz.Eq(string(filterOpt.Field), filterOpt.Value))
	case Less:
		return sqlz.Lt(string(filterOpt.Field), filterOpt.Value)
	case Greater:
		return sqlz.Gt(string(filterOpt.Field), filterOpt.Value)
	case LessOrEqual:
		return sqlz.Lte(string(filterOpt.Field), filterOpt.Value)
	case GreaterOrEqual:
		return sqlz.Gte(string(filterOpt.Field), filterOpt.Value)
	case Contains:
		return sqlz.Like(string(filterOpt.Field), fmt.Sprintf("%%%s%%", filterOpt.Value))
	case And:
		subfilters := make([]sqlz.WhereCondition, len(filterOpt.Subfilters))
		for i, opt := range filterOpt.Subfilters {
			subfilters[i] = whereConditionsFromOrderFilterOpts(opt)
		}
		return sqlz.And(subfilters...)
	case Or:
		subfilters := make([]sqlz.WhereCondition, len(filterOpt.Subfilters))
		for i, opt := range filterOpt.Subfilters {
			subfilters[i] = whereConditionsFromOrderFilterOpts(opt)
		}
		return sqlz.Or(subfilters...)
	// case Not:
	// 	return sqlz.Not()
	default:
		return nil
	}
}

func checkOrderQuery(query *Query) error {
	if query == nil {
		return nil
	}
	if query.Offset != 0 && query.Limit == 0 {
		return errors.New("can't use Offset without Limit")
	}
	return nil

}
