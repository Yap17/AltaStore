package request

import (
	"AltaStore/business/checkout"
)

type NewCheckoutShoppingCart struct {
	UserId         string `json:"userid" validate:"required"`
	ShoppingCartId string `json:"shoppingcartid" validate:"required"`
	Description    string `json:"description"`
}

func (n *NewCheckoutShoppingCart) ToBusinessCheckout() *checkout.Checkout {
	var checkout checkout.Checkout

	checkout.CreatedBy = n.UserId
	checkout.ShoppingCardId = n.ShoppingCartId
	checkout.Description = n.Description

	return &checkout
}
