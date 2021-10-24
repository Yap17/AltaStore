package checkout

type service struct {
	repository         Repository
	repoShoppingDetail RepoShoppingDetail
}

func NewService(repository Repository, repoShoppingDetail RepoShoppingDetail) Service {
	return &service{repository, repoShoppingDetail}
}

func (s *service) NewCheckoutShoppingCart(checkout *Checkout) error {
	return s.repository.NewCheckoutShoppingCart(checkout.toCheckout())
}

func (s *service) GetAllCheckout() (*[]Checkout, error) {
	return s.repository.GetAllCheckout()
}

func (s *service) GetCheckoutById(id string) (*CheckItemDetails, error) {
	dtCheckout, err := s.repository.GetCheckoutById(id)

	if err != nil {
		return nil, err
	}

	items, err := s.repoShoppingDetail.GetShopCartDetailById(dtCheckout.ShoppingCardId)
	if err != nil {
		return nil, err
	}
	details := toDetailItemInCart(items)

	return getCheckItemsDetails(dtCheckout, details), nil
}
