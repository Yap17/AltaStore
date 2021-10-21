package shopping

import (
	"AltaStore/api/v1/shopping/request"
	"time"

	"github.com/google/uuid"
)

type ShoppCart struct {
	ID          string
	IsCheckOut  bool
	Description string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ItemInCart struct {
	ID          string
	ProductId   string
	ProductName string
	Price       int
	Qty         int
	UpdatedAt   time.Time
}

type InsertItemInCartSpec struct {
	ID        string
	ProductId string
	Price     int
	Qty       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateItemInCartSpec struct {
	ID        string
	ProductId string
	Price     int
	Qty       int
	UpdatedAt time.Time
}

type ShopCartDetail struct {
	ID          string
	Description string
	CreatedBy   string
	UpdatedAt   time.Time
	Details     []ItemInCart
}

func getShopCartDetailFormat(sum *ShoppCart, details *[]ItemInCart) *ShopCartDetail {
	var shopCartDetail ShopCartDetail
	var itemInCart *ItemInCart

	shopCartDetail.ID = sum.ID
	shopCartDetail.Description = sum.Description
	shopCartDetail.CreatedBy = sum.CreatedBy
	shopCartDetail.UpdatedAt = sum.UpdatedAt

	for _, val := range *details {
		itemInCart = &val
		shopCartDetail.Details = append(shopCartDetail.Details, *itemInCart)
	}

	if shopCartDetail.Details == nil {
		shopCartDetail.Details = []ItemInCart{}
	}

	return &shopCartDetail
}

func insertItemFormat(item *request.DetailItemInShopCart) *InsertItemInCartSpec {
	return &InsertItemInCartSpec{
		ID:        uuid.NewString(),
		ProductId: item.ProductId,
		Price:     item.Price,
		Qty:       item.Qty,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func updateItemFormat(item *request.DetailItemInShopCart) *UpdateItemInCartSpec {
	return &UpdateItemInCartSpec{
		ID:        item.ID,
		ProductId: item.ProductId,
		Price:     item.Price,
		Qty:       item.Qty,
		UpdatedAt: time.Now(),
	}
}
