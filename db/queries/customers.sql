-- name: ListCustomers :many
SELECT uid,
       name,
       date_of_birth,
       city,
       zipcode,
       status,
       updated_by,
       updated_date
FROM customers
ORDER BY updated_date desc;

-- name: CreateCustomer :exec
INSERT INTO customers (name,
                       date_of_birth,
                       city,
                       zipcode,
                       status,
                       updated_by)
VALUES ($1,$2,$3,$4,$5,$6);