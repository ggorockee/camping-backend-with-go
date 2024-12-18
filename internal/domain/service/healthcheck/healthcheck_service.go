package healthcheckservice

type HealthCheckService interface {
	GetHealthCheck() error
}

type healthCheckService struct{}

// GetHealthCheck implements HealthCheckService.
func (h *healthCheckService) GetHealthCheck() error {
	return nil
}

func NewHealthCheckService() HealthCheckService {
	return &healthCheckService{}
}
