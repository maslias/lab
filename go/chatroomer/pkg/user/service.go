package user

type Service struct {
	Store UserStore
}

func NewService(store UserStore) *Service {
	return &Service{Store: store}
}

func (s *Service) CreateUser() {
}
