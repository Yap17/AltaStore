package shopping

import "time"

type ShoppCart struct {
	ID          string
	IsCheckOut  bool
	Description string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
