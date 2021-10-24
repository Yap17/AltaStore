package checkout

import "AltaStore/modules/shoppingdetail"

type Service interface {
	NewCheckoutShoppingCart(checkout *Checkout) error
	GetAllCheckout() (*[]Checkout, error)
	GetCheckoutById(id string) (*CheckItemDetails, error)
}

type Repository interface {
	NewCheckoutShoppingCart(checkout *Checkout) error
	GetAllCheckout() (*[]Checkout, error)
	GetCheckoutById(id string) (*Checkout, error)
}

type RepoShoppingDetail interface {
	GetShopCartDetailById(id string) (*[]shoppingdetail.ShopCartDetailItemWithProductName, error)
}
