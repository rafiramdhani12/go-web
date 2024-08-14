package entities

import "time"

type Product struct {
	ID          uint
	Name        string
	Category    Category
	Stock       int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
