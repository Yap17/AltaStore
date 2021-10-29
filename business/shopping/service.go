package shopping

import (
	"AltaStore/api/v1/shopping/request"
	"time"

	"github.com/google/uuid"
)

type service struct {
	repository     Repository
	repoCartDetail RepositoryCartDetail
}

func NewService(repository Repository, repoCartDetail RepositoryCartDetail) Service {
	return &service{repository, repoCartDetail}
}

func (s *service) GetShoppingCartByUserId(userid string) (*ShoppCart, error) {
	return s.repository.GetShoppingCartByUserId(userid)
}

func (s *service) NewShoppingCart(userid string, description string) (*ShoppCart, error) {
	return s.repository.NewShoppingCart(uuid.NewString(), userid, description, time.Now())
}

func (s *service) GetShopCartDetailById(id string) (*ShopCartDetail, error) {
	shopCart, err := s.repository.GetShoppingCartById(id)
	if err != nil {
		return nil, err
	}

	items, err := s.repoCartDetail.GetShopCartDetailById(id)
	if err != nil {
		return nil, err
	}

	cnvItems := toDetailItemInCart(items)

	return getShopCartDetailFormat(shopCart, cnvItems), nil
}

func (s *service) NewItemInShopCart(cartId string, item *request.DetailItemInShopCart) error {
	return s.repoCartDetail.NewItemInShopCart(cartId, insertItemFormat(item))

}

func (s *service) ModifyItemInShopCart(cartId string, productid string, item *request.DetailItemInShopCart) error {
	return s.repoCartDetail.ModifyItemInShopCart(cartId, productid, updateItemFormat(productid, item))
}

func (s *service) DeleteItemInShopCart(cartId string, productid string) error {
	return s.repoCartDetail.DeleteItemInShopCart(cartId, productid)
}
