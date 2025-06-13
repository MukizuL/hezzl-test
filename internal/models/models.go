package models

import "time"

type Goods struct {
	ID          int
	ProjectID   int
	Name        string
	Description string
	Priority    int
	Removed     bool
	CreatedAt   time.Time
}
