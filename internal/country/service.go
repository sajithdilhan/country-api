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

func (s *Service) GetCountryById(id int) (*Country, error) {
	return s.repo.GetCountryById(id)
}

func (s *Service) GetCountryByCode(code string) (*Country, error) {
	return s.repo.GetCountryByCode(code)
}
