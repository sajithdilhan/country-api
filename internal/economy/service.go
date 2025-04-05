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

func (s *Service) GetEconomyById(id int) (*Economy, error) {
	return s.repo.GetEconomyById(id)
}

func (s *Service) GetEconomyByCountryId(country_id int) (*Economy, error) {
	return s.repo.GetEconomyByCountryId(country_id)
}
