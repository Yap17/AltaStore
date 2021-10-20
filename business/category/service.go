package category

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAllCategory() (*[]Category, error) {
	return s.repository.GetAllCategory()
}
