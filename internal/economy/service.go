package economy

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllEconomies() ([]Economy, error) {
	return s.repo.GetAllEconomies()
}
