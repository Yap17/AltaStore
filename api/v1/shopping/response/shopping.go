package response

import "time"

type ShoppData struct {
	ID          string    `json:"id"`
	IsCheckOut  bool      `json:"ischeckout"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}
