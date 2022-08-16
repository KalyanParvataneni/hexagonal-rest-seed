-- name: ListTransactions :many
SELECT t.id,
       t.uid,
       a.id,
       t.transaction_type,
       t.created_by,
       t.transaction_date
FROM transactions t
         LEFT JOIN accounts a on a.uid = t.uid
ORDER BY t.id desc;

-- name: CreateTransaction :exec
INSERT INTO transactions (account_uid,
                          amount,
                          transaction_type,
                          created_by)
VALUES ($1, $2, $3, $4);