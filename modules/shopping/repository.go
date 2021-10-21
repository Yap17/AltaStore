package shopping

import (
	"AltaStore/business"
	"AltaStore/business/shopping"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type ShoppingCart struct {
	ID          string    `gorm:"id;type:uuid;primaryKey"`
	IsCheckOut  bool      `gorm:"is_check_out;type:boolean;default:false"`
	Description string    `gorm:"description;type:varchar(100)"`
	CreatedBy   string    `gorm:"created_by;type:uuid"`
	CreatedAt   time.Time `gorm:"created_at;type:timestamp"`
	UpdatedAt   time.Time `gorm:"updated_at;type:timestamp"`
	DeletedAt   time.Time `gorm:"deleted_at;type:timestamp"`
}

type ShoppingCartDetail struct {
	ID             string    `gorm:"id;type:uuid;primaryKey"`
	ShoppingCartId string    `gorm:"shopping_cart_id;type:varchar(50)"`
	ProductId      string    `gorm:"product_id;type:varchar(50)"`
	Price          int       `gorm:"price;type:integer"`
	Qty            int       `gorm:"qty;type:integer"`
	CreatedAt      time.Time `gorm:"created_at;type:timestamp"`
	UpdatedAt      time.Time `gorm:"updated_at;type:timestamp"`
	DeletedAt      time.Time `gorm:"deleted_at;type:timestamp"`
}

func (s *ShoppingCart) toShoppCart() *shopping.ShoppCart {
	return &shopping.ShoppCart{
		ID:          s.ID,
		IsCheckOut:  s.IsCheckOut,
		Description: s.Description,
		CreatedBy:   s.CreatedBy,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func (s *ShoppingCart) newShoppingCart(id string, userid string, description string, createdAt time.Time) {
	s.ID = id
	s.IsCheckOut = false
	s.Description = description
	s.CreatedBy = userid
	s.CreatedAt = createdAt
	s.UpdatedAt = createdAt
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetShoppingCartByUserId(id string) (*shopping.ShoppCart, error) {
	var shopCart ShoppingCart

	err := r.DB.First(&shopCart, "is_check_out = false and created_by = ?", id).Error
	if err != nil {
		return nil, err
	}

	return shopCart.toShoppCart(), nil
}

func (r *Repository) NewShoppingCart(id string, userid string, description string, createdAt time.Time) (*shopping.ShoppCart, error) {
	var shopCart ShoppingCart

	err := r.DB.First(&shopCart, "is_check_out = false and created_by = ?", id).Error

	// Pengecekan jika masih terdapat keranjang aktif maka dikembalikan bad request
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, business.ErrHasBeenModified
	}

	shopCart.newShoppingCart(id, userid, description, createdAt)

	if err := r.DB.Create(&shopCart).Error; err != nil {
		return nil, err
	}

	return shopCart.toShoppCart(), nil
}
