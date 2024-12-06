package store

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"your-project-name/internal/db"
)

// Store provides all functions to execute SQL queries and transactions
type Store interface {
	db.Querier
	execTx(ctx context.Context, fn func(*db.Queries) error) error
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*db.Queries
	connPool *pgxpool.Pool
}

// NewStore creates a new store instance
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		Queries:  db.New(connPool),
		connPool: connPool,
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}

// Example of a transaction method:
/*
type CreateUserTxParams struct {
    CreateUserParams
    // Add any additional parameters needed
}

type CreateUserTxResult struct {
    User User
    // Add any additional result fields
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
    var result CreateUserTxResult

    err := store.execTx(ctx, func(q *Queries) error {
        var err error

        result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
        if err != nil {
            return err
        }

        // Add any additional operations within the transaction

        return nil
    })

    return result, err
}
*/
