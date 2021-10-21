package shopping

import "time"

type Service interface {
	// Mengambil shopping cart aktif untuk berbelanja
	GetShoppingCartByUserId(id string) (*ShoppCart, error)

	// Membuat keranjang belanjaan baru, ketika keranjang belanjaan ada yang aktif akan dikembalikan error
	NewShoppingCart(id string, description string) (*ShoppCart, error)
}

type Repository interface {
	// Mengambil shopping cart aktif untuk berbelanja
	GetShoppingCartByUserId(id string) (*ShoppCart, error)

	// Membuat keranjang belanjaan baru, ketika keranjang belanjaan ada yang aktif akan dikembalikan error
	NewShoppingCart(id string, userid string, description string, createdAt time.Time) (*ShoppCart, error)
}
