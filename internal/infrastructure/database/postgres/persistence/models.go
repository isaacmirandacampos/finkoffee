// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package persistence

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type Expense struct {
	ID          int32           `db:"id" json:"id"`
	Value       decimal.Decimal `db:"value" json:"value"`
	Description string          `db:"description" json:"description"`
	UserID      int32           `db:"user_id" json:"user_id"`
	PaymentAt   time.Time       `db:"payment_at" json:"payment_at"`
	PaidAt      sql.NullTime    `db:"paid_at" json:"paid_at"`
	Note        sql.NullString  `db:"note" json:"note"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime    `db:"deleted_at" json:"deleted_at"`
}

type User struct {
	ID        int32          `db:"id" json:"id"`
	FullName  string         `db:"full_name" json:"full_name"`
	Email     string         `db:"email" json:"email"`
	Password  sql.NullString `db:"password" json:"password"`
	CreatedAt time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime   `db:"deleted_at" json:"deleted_at"`
}
