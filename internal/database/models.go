// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package tutorial

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID          pgtype.UUID
	CreatedTime pgtype.Timestamp
	UpdateTime  pgtype.Timestamp
	Name        string
}
