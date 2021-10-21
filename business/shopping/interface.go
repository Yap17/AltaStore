package shopping

import (
	"AltaStore/api/v1/shopping/request"
	"time"
)

type Service interface {
	// Mengambil shopping cart aktif untuk berbelanja
	GetShoppingCartByUserId(userid string) (*ShoppCart, error)

	// Membuat keranjang belanjaan baru, ketika keranjang belanjaan ada yang aktif akan dikembalikan error
	NewShoppingCart(userid string, description string) (*ShoppCart, error)

	// Mengambil detail item pada shopping cart
	GetShopCartDetailById(id string) (*ShopCartDetail, error)

	// Menambahkan item produk pada shopping cart
	NewItemInShopCart(cartId string, item *request.DetailItemInShopCart) error

	// Merubah item produk pada shopping cart
	ModifyItemInShopCart(cartId string, item *request.DetailItemInShopCart) error

	// Menghapus item produk pada shopping cart
	DeleteItemInShopCart(cartId string, id string) error
}

type Repository interface {
	// Mengambil shopping cart aktif untuk berbelanja berdasarkan userid
	GetShoppingCartByUserId(userid string) (*ShoppCart, error)

	// Mengambil shopping cart berdasarkan id
	GetShoppingCartById(id string) (*ShoppCart, error)

	// Membuat keranjang belanjaan baru, ketika keranjang belanjaan ada yang aktif akan dikembalikan error
	NewShoppingCart(id string, userid string, description string, createdAt time.Time) (*ShoppCart, error)
}

type RepositoryCartDetail interface {
	GetShopCartDetailById(id string) (*[]ItemInCart, error)

	// Menambahkan item produk pada shopping cart
	NewItemInShopCart(cartId string, item *InsertItemInCartSpec) error

	// Merubah item produk pada shopping cart
	ModifyItemInShopCart(cartId string, item *UpdateItemInCartSpec) error

	// Menghapus item produk pada shopping cart
	DeleteItemInShopCart(cartId string, id string) error
}
