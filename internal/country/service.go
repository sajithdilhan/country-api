package country

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllCountries() ([]Country, error) {
	return s.repo.GetAllCountries()
}
