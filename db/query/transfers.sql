-- name: CreateTransfers :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTranfers :one
SELECT * FROM transfers
WHERE id = $1;

-- name: UpdateTransfers :exec
UPDATE transfers
SET amount = $2
WHERE id = $1;

-- name: DeleteTransfers :exec
DELETE FROM transfers
WHERE id = $1;