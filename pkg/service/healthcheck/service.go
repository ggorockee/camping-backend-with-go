package healthcheck

type Service interface {
	GetHealthCheck() error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) GetHealthCheck() error {
	return nil
}
