package entity

import (
	"time"
)

type Chikuchiku struct {
	ID        uint
	Date      time.Time
	Message   string
	URL       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
