package billingaddress

import "time"

// BillingAddress stores information about account making the payment
type BillingAddress struct {
	CheckOutID  string    `gorm:"checkout_id;type:uuid;primaryKey"`
	FullName    string    `gorm:"full_name"`
	Email       string    `gorm:"email"`
	PhoneNumber string    `gorm:"phone_number"`
	CreatedAt   time.Time `gorm:"created_at"`
	CreatedBy   string    `gorm:"created_by;type:varchar(50)"`
	UpdatedAt   time.Time `gorm:"updated_at"`
	UpdatedBy   string    `gorm:"updated_by;type:varchar(50)"`
	DeletedAt   time.Time `gorm:"deleted_at"`
	DeletedBy   string    `gorm:"deleted_by;type:varchar(50)"`
}
