package shoppingdetail

import (
	"AltaStore/modules/product"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type ShoppingCartDetail struct {
	ID             string          `gorm:"id;type:uuid;primaryKey"`
	ShoppingCartId string          `gorm:"shopping_cart_id;type:uuid;index:shopping_detail_uniq"`
	ProductId      string          `gorm:"product_id;type:uuid;index:shopping_detail_uniq"`
	Product        product.Product `gorm:"foreignKey:ProductId"`
	Price          int             `gorm:"price;type:integer"`
	Qty            int             `gorm:"qty;type:integer"`
	CreatedAt      time.Time       `gorm:"created_at;type:timestamp"`
	UpdatedAt      time.Time       `gorm:"updated_at;type:timestamp"`
	DeletedAt      time.Time       `gorm:"deleted_at;type:timestamp"`
}

type ShopCartDetailItemWithProductName struct {
	ShoppingCartDetail
	ProductName string `gorm:"name"`
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

func createItemInCart(cartId string, item *InsertItemInCartSpec) *ShoppingCartDetail {
	return &ShoppingCartDetail{
		ID:             item.ID,
		ShoppingCartId: cartId,
		ProductId:      item.ProductId,
		Price:          item.Price,
		Qty:            item.Qty,
		CreatedAt:      item.CreatedAt,
		UpdatedAt:      item.UpdatedAt,
	}
}

func modifyItemInCart(item *UpdateItemInCartSpec) *ShoppingCartDetail {
	return &ShoppingCartDetail{
		ProductId: item.ProductId,
		Price:     item.Price,
		Qty:       item.Qty,
		UpdatedAt: item.UpdatedAt,
	}
}

func (r *Repository) GetShopCartDetailById(id string) (*[]ShopCartDetailItemWithProductName, error) {
	var shopCartDetail []ShopCartDetailItemWithProductName

	err := r.DB.Raw("select t1.*, t2.name product_name from shopping_cart_details t1 inner join products t2 on t2.id = t1.product_id where t1.shopping_cart_id = ?", id).Scan(&shopCartDetail).Error
	if err != nil {
		return nil, err
	}

	return &shopCartDetail, nil
}

func (r *Repository) NewItemInShopCart(cartId string, item *InsertItemInCartSpec) error {
	err := r.DB.Create(createItemInCart(cartId, item)).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ModifyItemInShopCart(cartId string, item *UpdateItemInCartSpec) error {
	var itemInCart ShoppingCartDetail

	err := r.DB.Where("shopping_cart_id = ? and id = ?", cartId, item.ID).Find(&itemInCart).Error
	if err != nil {
		return err
	}

	r.DB.Model(&itemInCart).Updates(modifyItemInCart(item))

	return nil
}

func (r *Repository) DeleteItemInShopCart(cartId string, id string) error {
	var itemInCart ShoppingCartDetail

	err := r.DB.Where("shopping_cart_id = ? and id = ?", cartId, id).Find(&itemInCart).Error
	if err != nil {
		return err
	}

	r.DB.Delete(&itemInCart)
	return nil
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
