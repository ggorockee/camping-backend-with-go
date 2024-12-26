package test

import (
	"bytes"
	userdto "camping-backend-with-go/internal/application/dto/user"
	"camping-backend-with-go/internal/domain/entity"
	userservice "camping-backend-with-go/internal/domain/service/user"
	userroute "camping-backend-with-go/pkg/api/route/user"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	mock.Mock
	users map[string]*entity.User // 간단한 인메모리 저장소
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*entity.User),
	}
}

func (r *MockUserRepository) HashPassword(password string, context ...*fiber.Ctx) (string, error) {
	args := r.Called(password)
	return args.String(0), args.Error(1)
}

func (r *MockUserRepository) GetUserByEmail(email string, context ...*fiber.Ctx) (*entity.User, error) {
	args := r.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (r *MockUserRepository) CheckPasswordHash(password, hash string, context ...*fiber.Ctx) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (r *MockUserRepository) ChangePassword(input *userdto.ChangePasswordReq, context ...*fiber.Ctx) error {
	args := r.Called(input)
	return args.Error(0)
}

func (r *MockUserRepository) ValidToken(t *jwt.Token, id string, context ...*fiber.Ctx) bool {

	args := r.Called(t, id)
	return args.Bool(0)
}

func (r *MockUserRepository) GetUserById(id string, context ...*fiber.Ctx) (*entity.User, error) {
	args := r.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (r *MockUserRepository) GetValueFromToken(key string, context ...*fiber.Ctx) string {
	args := r.Called(key)
	return args.String(0)
}

func (r *MockUserRepository) CreateUser(key string, context ...*fiber.Ctx) string {
	args := r.Called(key)
	return args.String(0)
}

func TestCreateUser(t *testing.T) {
	app := fiber.New()
	mockRepo := new(MockUserRepository)
	userUseCase := userservice.NewUserService(mockRepo)

	v1 := app.Group("/api/v1")
	userroute.UserRouter(v1, userUseCase)

	// ID: "1", Name: "John Doe", Email: "john@example.com"
	user := &entity.User{
		Common: entity.Common{
			Id:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:    "john@example.com",
		Password: "test123!@#",
		Username: "john",
		Role:     "client",
	}
	mockRepo.On("CreateUser", mock.AnythingOfType("*entity.User")).Return(nil)

	// 요청 본문 생성
	reqBody := map[string]string{
		"email":            user.Email,
		"password":         user.Password,
		"password_confirm": user.Password,
		"username":         user.Username,
	}
	body, _ := json.Marshal(reqBody)
	// 요청 생성
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// 요청 실행
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// 응답 본문 파싱
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.NoError(t, err)

	// 응답 검증
	assert.Equal(t, user.Email, responseBody["email"])
	assert.Equal(t, user.Username, responseBody["username"])
	assert.NotEmpty(t, responseBody["id"])

	// 모의 객체 기대값 검증
	mockRepo.AssertExpectations(t)
}
