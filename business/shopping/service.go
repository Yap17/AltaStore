package shopping

import (
	"time"

	"github.com/google/uuid"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetShoppingCartByUserId(id string) (*ShoppCart, error) {
	return s.repository.GetShoppingCartByUserId(id)
}

func (s *service) NewShoppingCart(id string, description string) (*ShoppCart, error) {
	return s.repository.NewShoppingCart(uuid.NewString(), id, description, time.Now())
}
