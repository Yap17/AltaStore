package category

import "time"

type CategorySpec struct {
	UserId string
	Code   string
	Name   string
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAllCategory() (*[]Category, error) {
	return s.repository.GetAllCategory()
}

func (s *service) FindCategoryById(id string) (*Category, error) {
	return s.repository.FindCategoryById(id)
}

func (s *service) InsertCategory(category *CategorySpec) error {
	dataCategory := NewProductCategory(
		category.Code, category.Name, category.UserId, time.Now(),
	)
	return s.repository.InsertCategory(dataCategory)
}

func (s *service) UpdateCategory(id string, category *CategorySpec) error {
	dataCategory := ModifyProductCategory(category.Code, category.Name, category.UserId, time.Now())

	return s.repository.UpdateCategory(id, dataCategory)
}

func (s *service) DeleteCategory(id string, userid string) error {
	return s.repository.DeleteCategory(id, userid)
}
