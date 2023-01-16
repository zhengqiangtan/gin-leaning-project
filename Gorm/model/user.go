package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type Company struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
