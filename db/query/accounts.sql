-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id;

-- name: CountAccounts :one
SELECT count(*) FROM accounts;

-- name: CreateAccount :exec
INSERT INTO accounts (name, website, lat, long, primary_poc, sales_rep_id) VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateAccount :exec
UPDATE accounts SET name = $2, website = $3, lat = $4, long = $5, primary_poc = $6, sales_rep_id = $7 WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;
