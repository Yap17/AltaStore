package checkout

import (
	"AltaStore/business/checkoutpayment"

	snap "github.com/midtrans/midtrans-go/snap"
)

type service struct {
	checkoutpaymentService checkoutpayment.Service
	repository             Repository
	repoShoppingDetail     RepoShoppingDetail
}

func NewService(
	checkoutpaymentService checkoutpayment.Service,
	repository Repository,
	repoShoppingDetail RepoShoppingDetail,

) Service {
	return &service{
		checkoutpaymentService,
		repository,
		repoShoppingDetail,
	}
}

func (s *service) NewCheckoutShoppingCart(checkout *Checkout) (*snap.Response, error) {
	var newCheckout = checkout.toCheckout()
	dets, err := s.repoShoppingDetail.GetShopCartDetailById(newCheckout.ShoppingCardId)
	if err != nil {
		return nil, err
	}
	err = s.repository.NewCheckoutShoppingCart(newCheckout)
	if err != nil {
		return nil, err
	}
	var sum int64 = 0
	for _, val := range *dets {
		sum += int64(val.Qty)
	}

	return s.checkoutpaymentService.GenerateSnapPayment(
		newCheckout.CreatedBy,
		newCheckout.ID,
		sum)
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
