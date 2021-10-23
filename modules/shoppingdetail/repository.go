package shoppingdetail

import (
	"time"

	"AltaStore/business/shopping"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
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

type ShopCartDetailItemWithProductName struct {
	ShoppingCartDetail
	ProductName string `gorm:"name"`
}

func (s *ShopCartDetailItemWithProductName) toDetailItem() *shopping.ItemInCart {
	return &shopping.ItemInCart{
		ID:          s.ID,
		ProductId:   s.ProductId,
		ProductName: s.ProductName,
		Price:       s.Price,
		Qty:         s.Qty,
		UpdatedAt:   s.UpdatedAt,
	}
}

func createItemInCart(cartId string, item *shopping.InsertItemInCartSpec) *ShoppingCartDetail {
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

func modifyItemInCart(item *shopping.UpdateItemInCartSpec) *ShoppingCartDetail {
	return &ShoppingCartDetail{
		ProductId: item.ProductId,
		Price:     item.Price,
		Qty:       item.Qty,
		UpdatedAt: item.UpdatedAt,
	}
}

func (r *Repository) GetShopCartDetailById(id string) (*[]shopping.ItemInCart, error) {
	var shopCartDetail []ShopCartDetailItemWithProductName
	var itemInCart []shopping.ItemInCart

	// err := r.DB.Model(&ShoppingCartDetail{}).Select("shopping_cart_details.*, t1.name").Joins("inner join products t1 on trim(cast(t1.id as varchar)) = trim(product_id)").Scan(&shopCartDetail).Error
	// err := r.DB.Where("shopping_cart_id = ?", id).Order("created_at asc").Find(&shopCartDetail).Error
	err := r.DB.Raw("select t1.*, t2.name product_name from shopping_cart_details t1 inner join products t2 on cast(t2.id as varchar) = t1.product_id where t1.shopping_cart_id = ?", id).Scan(&shopCartDetail).Error
	if err != nil {
		return nil, err
	}

	for _, val := range shopCartDetail {
		itemInCart = append(itemInCart, *val.toDetailItem())
	}

	return &itemInCart, nil
}

func (r *Repository) NewItemInShopCart(cartId string, item *shopping.InsertItemInCartSpec) error {
	err := r.DB.Create(createItemInCart(cartId, item)).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ModifyItemInShopCart(cartId string, item *shopping.UpdateItemInCartSpec) error {
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
