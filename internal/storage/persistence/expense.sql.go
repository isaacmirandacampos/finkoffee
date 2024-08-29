// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: expense.sql

package persistence

import (
	"context"

	"github.com/shopspring/decimal"
)

const createExpense = `-- name: CreateExpense :one
INSERT INTO expenses (
  description, value
) VALUES (
  $1, $2
) RETURNING id, value, description, created_at, updated_at, deleted_at
`

type CreateExpenseParams struct {
	Description string          `db:"description" json:"description"`
	Value       decimal.Decimal `db:"value" json:"value"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, createExpense, arg.Description, arg.Value)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Value,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteExpense = `-- name: DeleteExpense :one
UPDATE expenses
SET deleted_at = now()
WHERE id = $1 and deleted_at is null RETURNING id, value, description, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteExpense(ctx context.Context, id int32) (Expense, error) {
	row := q.db.QueryRowContext(ctx, deleteExpense, id)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Value,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getExpense = `-- name: GetExpense :one
SELECT id, value, description, created_at, updated_at, deleted_at FROM expenses WHERE id = $1 and deleted_at is null
`

func (q *Queries) GetExpense(ctx context.Context, id int32) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getExpense, id)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Value,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getLastExpense = `-- name: GetLastExpense :one
SELECT id, value, description, created_at, updated_at, deleted_at FROM expenses where deleted_at is null ORDER BY id desc LIMIT 1
`

func (q *Queries) GetLastExpense(ctx context.Context) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getLastExpense)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Value,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listExpenses = `-- name: ListExpenses :many
SELECT id, value, description, created_at, updated_at, deleted_at FROM expenses where deleted_at is null ORDER BY id desc
`

func (q *Queries) ListExpenses(ctx context.Context) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, listExpenses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Expense{}
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Value,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExpense = `-- name: UpdateExpense :one
UPDATE expenses
SET description = $1, value = $2, updated_at = now()
WHERE id = $3 and deleted_at is null RETURNING id, value, description, created_at, updated_at, deleted_at
`

type UpdateExpenseParams struct {
	Description string          `db:"description" json:"description"`
	Value       decimal.Decimal `db:"value" json:"value"`
	ID          int32           `db:"id" json:"id"`
}

func (q *Queries) UpdateExpense(ctx context.Context, arg UpdateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, updateExpense, arg.Description, arg.Value, arg.ID)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Value,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
