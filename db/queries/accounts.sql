-- name: ListAccounts :many
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
ORDER BY a.id desc;

-- name: CreateAccount :exec
INSERT INTO accounts (customer_uid,
                      opening_date,
                      account_type,
                      amount,
                      status,
                      updated_by)
VALUES ($1, $2, $3, $4, $5, $6);