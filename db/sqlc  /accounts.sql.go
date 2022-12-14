// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: accounts.sql

package db

import (
	"context"

	"github.com/jackc/pgtype"
)

const createAccount = `-- name: CreateAccount :exec
INSERT INTO accounts (customer_uid,
                      opening_date,
                      account_type,
                      amount,
                      status,
                      updated_by)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateAccountParams struct {
	CustomerUid pgtype.UUID    `json:"customer_uid"`
	OpeningDate pgtype.Date    `json:"opening_date"`
	AccountType pgtype.Text    `json:"account_type"`
	Amount      pgtype.Numeric `json:"amount"`
	Status      pgtype.Int2    `json:"status"`
	UpdatedBy   pgtype.Text    `json:"updated_by"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.db.Exec(ctx, createAccount,
		arg.CustomerUid,
		arg.OpeningDate,
		arg.AccountType,
		arg.Amount,
		arg.Status,
		arg.UpdatedBy,
	)
	return err
}

const listAccounts = `-- name: ListAccounts :many
SELECT a.id,
       a.uid,
       c.name,
       a.opening_date,
       a.account_type,
       a.amount,
       a.status,
       a.updated_by,
       a.updated_date
FROM accounts a
         LEFT JOIN customers c on a.uid = c.uid
ORDER BY a.id desc
`

type ListAccountsRow struct {
	ID          pgtype.Int4        `json:"id"`
	Uid         pgtype.UUID        `json:"uid"`
	Name        pgtype.Text        `json:"name"`
	OpeningDate pgtype.Date        `json:"opening_date"`
	AccountType pgtype.Text        `json:"account_type"`
	Amount      pgtype.Numeric     `json:"amount"`
	Status      pgtype.Int2        `json:"status"`
	UpdatedBy   pgtype.Text        `json:"updated_by"`
	UpdatedDate pgtype.Timestamptz `json:"updated_date"`
}

func (q *Queries) ListAccounts(ctx context.Context) ([]ListAccountsRow, error) {
	rows, err := q.db.Query(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAccountsRow
	for rows.Next() {
		var i ListAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.Name,
			&i.OpeningDate,
			&i.AccountType,
			&i.Amount,
			&i.Status,
			&i.UpdatedBy,
			&i.UpdatedDate,
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
