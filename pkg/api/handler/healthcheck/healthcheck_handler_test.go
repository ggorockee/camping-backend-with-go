package healthcheckhandler

import (
	"camping-backend-with-go/internal/domain/presenter"
	"encoding/json"
	"errors"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHealthCheckService struct {
	mock.Mock
}

func (m *MockHealthCheckService) GetHealthCheck() error {
	args := m.Called()
	return args.Error(0)
}

func TestGetHealthCheck(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(*MockHealthCheckService)
		expectedStatus int
		expectedError  bool
		expectedMsg    string
	}{
		{
			name: "성공 여부",
			setupMock: func(m *MockHealthCheckService) {
				m.On("GetHealthCheck").Return(nil)
			},
			expectedStatus: fiber.StatusOK,
			expectedError:  false,
			expectedMsg:    "welcome",
		},
		{
			name: "실패",
			setupMock: func(m *MockHealthCheckService) {
				m.On("GetHealthCheck").Return(errors.New("service unavailable"))
			},
			expectedStatus: fiber.StatusInternalServerError,
			expectedError:  true,
			expectedMsg:    "service unavailable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock service 설정
			mockService := new(MockHealthCheckService)
			tt.setupMock(mockService)

			// fiber 앱 및 핸들러 설정
			app := fiber.New()
			app.Get("/api/v1/healthcheck", GetHealthCheck(mockService))

			// 요청 생성
			req := httptest.NewRequest(http.MethodGet, "/api/v1/healthcheck", nil)

			// 테스트 실행
			resp, _ := app.Test(req)

			// 상태 코드 검증
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// log.Println("resp>>>>>>>>>>", )

			// 응답 본문 파싱
			var result presenter.JsonResponse
			json.NewDecoder(resp.Body).Decode(&result)

			// assert.NoError(t, err)

			// 응답 내용 검증

			assert.Equal(t, tt.expectedError, result.Error)
			assert.Equal(t, tt.expectedMsg, result.Message)

			// assert.Equal(t, tt.expectedMsg, result.Message)
		})
	}
}
