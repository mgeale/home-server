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
)

type Field string

type Query struct {
	Filters []Filter
	Sort    Sort
	Limit   uint
	Offset  uint
}

type Sort struct {
	Field     Field
	Direction SortDirection
}

type Filter struct {
	Field Field
	Kind  FilterKind
	Value interface{}
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
	whereConditions, err := whereConditionsFromOrderFilterOpts(opts.Filters)
	if err != nil {
		return nil, err
	}
	if len(whereConditions) != 0 {
		stmt.Where(whereConditions...)
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

func whereConditionsFromOrderFilterOpts(filterOpts []Filter) ([]sqlz.WhereCondition, error) {
	whereConditions := make([]sqlz.WhereCondition, len(filterOpts))
	for i, filterOpt := range filterOpts {
		switch filterOpt.Kind {
		case Equal:
			whereConditions[i] = sqlz.Eq(string(filterOpt.Field), filterOpt.Value)
		case NotEqual:
			whereConditions[i] = sqlz.Not(sqlz.Eq(string(filterOpt.Field), filterOpt.Value))
		case Less:
			whereConditions[i] = sqlz.Lt(string(filterOpt.Field), filterOpt.Value)
		case Greater:
			whereConditions[i] = sqlz.Gt(string(filterOpt.Field), filterOpt.Value)
		case LessOrEqual:
			whereConditions[i] = sqlz.Lte(string(filterOpt.Field), filterOpt.Value)
		case GreaterOrEqual:
			whereConditions[i] = sqlz.Gte(string(filterOpt.Field), filterOpt.Value)
		case Contains:
			whereConditions[i] = sqlz.Like(string(filterOpt.Field), fmt.Sprintf("%%%s%%", filterOpt.Value))
		default:
			return nil, fmt.Errorf("db.FindOrder: unknown FilterOpt.Kind: %s", filterOpt.Kind)
		}
	}
	return whereConditions, nil
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
