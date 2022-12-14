// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: transactions.sql

package db

import (
	"context"
	"database/sql"

	"github.com/jackc/pgtype"
)

const createTransaction = `-- name: CreateTransaction :exec
INSERT INTO transactions (account_uid,
                          amount,
                          transaction_type,
                          created_by)
VALUES ($1, $2, $3, $4)
`

type CreateTransactionParams struct {
	AccountUid      pgtype.UUID    `json:"account_uid"`
	Amount          pgtype.Numeric `json:"amount"`
	TransactionType pgtype.Text    `json:"transaction_type"`
	CreatedBy       pgtype.Text    `json:"created_by"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) error {
	_, err := q.db.Exec(ctx, createTransaction,
		arg.AccountUid,
		arg.Amount,
		arg.TransactionType,
		arg.CreatedBy,
	)
	return err
}

const listTransactions = `-- name: ListTransactions :many
SELECT t.id,
       t.uid,
       a.id,
       t.transaction_type,
       t.created_by,
       t.transaction_date
FROM transactions t
         LEFT JOIN accounts a on a.uid = t.uid
ORDER BY t.id desc
`

type ListTransactionsRow struct {
	ID              pgtype.Int4        `json:"id"`
	Uid             pgtype.UUID        `json:"uid"`
	ID_2            sql.NullInt32      `json:"id_2"`
	TransactionType pgtype.Text        `json:"transaction_type"`
	CreatedBy       pgtype.Text        `json:"created_by"`
	TransactionDate pgtype.Timestamptz `json:"transaction_date"`
}

func (q *Queries) ListTransactions(ctx context.Context) ([]ListTransactionsRow, error) {
	rows, err := q.db.Query(ctx, listTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListTransactionsRow
	for rows.Next() {
		var i ListTransactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.ID_2,
			&i.TransactionType,
			&i.CreatedBy,
			&i.TransactionDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
