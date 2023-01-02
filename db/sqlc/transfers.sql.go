// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: transfers.sql

package db

import (
	"context"
)

const createTransfers = `-- name: CreateTransfers :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransfersParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfers, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfers = `-- name: DeleteTransfers :exec
DELETE FROM transfers
WHERE id = $1
`

func (q *Queries) DeleteTransfers(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfers, id)
	return err
}

const getTranfers = `-- name: GetTranfers :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = $1
`

func (q *Queries) GetTranfers(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTranfers, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const updateTransfers = `-- name: UpdateTransfers :exec
UPDATE transfers
SET amount = $2
WHERE id = $1
`

type UpdateTransfersParams struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) UpdateTransfers(ctx context.Context, arg UpdateTransfersParams) error {
	_, err := q.db.ExecContext(ctx, updateTransfers, arg.ID, arg.Amount)
	return err
}
