-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount, currency)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetTransfer :one
SELECT *
FROM transfers
WHERE id = $1
LIMIT 1;
-- name: ListTransfer :many
SELECT *
FROM transfers
ORDER BY id;
-- name: UpdateTransfer :one
UPDATE transfers
set from_account_id = $2,
    to_account_id = $3,
    amount = $4,
    currency = $5
WHERE id = $1
RETURNING *;
-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;