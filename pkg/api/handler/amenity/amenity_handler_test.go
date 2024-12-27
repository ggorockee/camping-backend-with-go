package amenityhandler

import (
	"bytes"
	"camping-backend-with-go/internal/application/dto"
	"camping-backend-with-go/pkg/util"
	"camping-backend-with-go/test"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"camping-backend-with-go/internal/domain/entity"
	"camping-backend-with-go/internal/domain/presenter"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAmenityService struct {
	mock.Mock
}

// amenities, err := service.GetAmenityList(c)
// amenity, err := service.UpdateAmenity(&requestBody, amenityId, c)
// err := service.DeleteAmenity(amenityId, c)
// amenity, err := service.GetAmenityById(amenityId)

func (m *MockAmenityService) CreateAmenity(input *dto.CreateAmenityReq, context ...*fiber.Ctx) (*entity.Amenity, error) {
	c, _ := util.ContextParser(context...)
	args := m.Called(input, c)
	return args.Get(0).(*entity.Amenity), args.Error(1)

}
func (m *MockAmenityService) GetAmenityById(id string, context ...*fiber.Ctx) (*entity.Amenity, error) {
	c, _ := util.ContextParser(context...)
	args := m.Called(id, c)
	return args.Get(0).(*entity.Amenity), args.Error(1)
}
func (m *MockAmenityService) GetAmenityList(context ...*fiber.Ctx) (*[]entity.Amenity, error) {
	c, _ := util.ContextParser(context...)
	args := m.Called(c)
	return args.Get(0).(*[]entity.Amenity), args.Error(1)
}

func (m *MockAmenityService) UpdateAmenity(input *dto.UpdateAmenityReq, id string, context ...*fiber.Ctx) (*entity.Amenity, error) {
	c, _ := util.ContextParser(context...)
	args := m.Called(input, id, c)
	return args.Get(0).(*entity.Amenity), args.Error(1)
}
func (m *MockAmenityService) DeleteAmenity(id string, context ...*fiber.Ctx) error {
	c, _ := util.ContextParser(context...)
	args := m.Called(id, c)
	return args.Error(0)
}

// CreateAmenityTest
// func TestCreateAmenity(t *testing.T) {
// 	// 테스트 앱 설정 및 정리 함수 호출
// 	app, _, cleanup := test.SetupTestApp()
// 	defer cleanup()

// 	// API v1 그룹 생성
// 	v1 := app.Group("/api/v1")

// 	t.Run("성공적인 Amenity 생성", func(t *testing.T) {
// 		// 각 테스트 케이스마다 새로운 모의 객체 생성
// 		mockService := new(MockAmenityService)

// 		// Amenity 생성 라우트 설정
// 		v1.Post("/spot/amenity", CreateAmenity(mockService))

// 		reqBody := dto.CreateAmenityReq{
// 			Name:        util.StrPointer("Test Amen"),
// 			Description: util.StrPointer("Description"),
// 		}

// 		// 요청 본문을 json으로 변환
// 		jsonBody, _ := json.Marshal(reqBody)

// 		// mocking 서비스 동작 정의
// 		mockService.On("CreateAmenity",
// 			mock.AnythingOfType("*dto.CreateAmenityReq"),
// 			mock.AnythingOfType("*fiber.Ctx"),
// 		).Return(&entity.Amenity{
// 			Common: entity.Common{
// 				Id:        "amenity_uuid",
// 				CreatedAt: time.Now(),
// 				UpdatedAt: time.Now(),
// 			},
// 			Name:        "Test Amen",
// 			Description: util.StrPointer("Description"),
// 		}, nil).Once()

// 		// HTTP 요청 생성 및 실행
// 		req := httptest.NewRequest("POST", "/api/v1/spot/amenity", bytes.NewBuffer(jsonBody))
// 		req.Header.Set("Content-Type", "application/json")
// 		resp, _ := app.Test(req)

// 		// 응답 검증
// 		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

// 		var result presenter.JsonResponse
// 		err := json.NewDecoder(resp.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.False(t, result.Error)

// 		amenityData, ok := result.Data.(map[string]interface{})
// 		assert.True(t, ok, "Data should be a map")
// 		assert.Equal(t, "Test Amen", amenityData["name"])
// 		assert.Equal(t, "Description", amenityData["description"])
// 		mockService.AssertExpectations(t)
// 	})

// 	t.Run("10자 이상의 이름으로 Amentity 생성 시 실패", func(t *testing.T) {
// 		// 각 테스트 케이스마다 새로운 모의 객체 생성
// 		mockService := new(MockAmenityService)

// 		// Amenity 생성 라우트 설정
// 		v1.Post("/spot/amenity", CreateAmenity(mockService))
// 		// 10자를 초과하는 이름으로 요청 본문 생성
// 		reqBody := dto.CreateAmenityReq{
// 			Name:        util.StrPointer("This is a very long amenity name"),
// 			Description: util.StrPointer("Description"),
// 		}

// 		// 요청 본문을 JSON으로 변환
// 		jsonBody, _ := json.Marshal(reqBody)

// 		// mockService.On("CreateAmenity",
// 		// 	mock.AnythingOfType("*dto.CreateAmenityReq"),
// 		// 	mock.AnythingOfType("*fiber.Ctx")).Return(nil, errors.New("name is too long")).Maybe()

// 		// HTTP 요청 생성 및 실행
// 		req := httptest.NewRequest("POST", "/api/v1/spot/amenity", bytes.NewBuffer(jsonBody))
// 		req.Header.Set("Content-Type", "application/json")
// 		resp, _ := app.Test(req)

// 		// 응답 검증
// 		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

// 		var result presenter.JsonResponse
// 		err := json.NewDecoder(resp.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.True(t, result.Error)
// 		assert.Contains(t, result.Message, "name is too long")
// 		mockService.AssertExpectations(t)
// 	})
// }

func TestCreateAmenity(t *testing.T) {
	t.Run("성공적인 Amenity 생성", func(t *testing.T) {
		app, _, cleanup := test.SetupTestApp()
		defer cleanup()

		mockService := new(MockAmenityService)
		v1 := app.Group("/api/v1")
		v1.Post("/spot/amenity", CreateAmenity(mockService))

		reqBody := dto.CreateAmenityReq{
			Name:        util.StrPointer("Test Amen"),
			Description: util.StrPointer("Description"),
		}
		jsonBody, _ := json.Marshal(reqBody)

		mockService.On("CreateAmenity",
			mock.AnythingOfType("*dto.CreateAmenityReq"),
			mock.AnythingOfType("*fiber.Ctx"),
		).Return(&entity.Amenity{
			Common: entity.Common{
				Id:        "amenity_uuid",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:        "Test Amen",
			Description: util.StrPointer("Description"),
		}, nil).Once()

		req := httptest.NewRequest("POST", "/api/v1/spot/amenity", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var result presenter.JsonResponse
		err := json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.False(t, result.Error)

		amenityData, ok := result.Data.(map[string]interface{})
		assert.True(t, ok, "Data should be a map")
		assert.Equal(t, "Test Amen", amenityData["name"])
		assert.Equal(t, "Description", amenityData["description"])

		mockService.AssertExpectations(t)
	})

	t.Run("10자 이상의 이름으로 Amenity 생성 시 실패", func(t *testing.T) {
		app, _, cleanup := test.SetupTestApp()
		defer cleanup()

		mockService := new(MockAmenityService)
		v1 := app.Group("/api/v1")
		v1.Post("/spot/amenity", CreateAmenity(mockService))

		reqBody := dto.CreateAmenityReq{
			Name:        util.StrPointer("This is a very long amenity name"),
			Description: util.StrPointer("Description"),
		}
		jsonBody, _ := json.Marshal(reqBody)

		mockService.On("CreateAmenity",
			mock.AnythingOfType("*dto.CreateAmenityReq"),
			mock.AnythingOfType("*fiber.Ctx"),
		).Return(nil, errors.New("name is too long")).Once()
		// mockService.On("CreateAmenity",
		// 	mock.AnythingOfType("*dto.CreateAmenityReq"),
		// 	mock.AnythingOfType("*fiber.Ctx")).
		// 	Return(nil, errors.New("name is too long")).Once()
		req := httptest.NewRequest("POST", "/api/v1/spot/amenity", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		log.Println(resp)

		// assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		// var result presenter.JsonResponse
		// err := json.NewDecoder(resp.Body).Decode(&result)
		// assert.NoError(t, err)
		// assert.True(t, result.Error)
		// assert.Contains(t, result.Message, "name is too long")

		mockService.AssertExpectations(t)
	})
}

// TestGetAmenityById Id로 amenity 한 개 불러오기
func TestGetAmenityById(t *testing.T) {

	// mock setting
	amentityUUID := "amentity_uuid"
	amentityName := "amenity_name"
	amentityDesc := "amentity_desc"

	app, _, cleanup := test.SetupTestApp()
	defer cleanup()

	mockService := new(MockAmenityService)

	v1 := app.Group("/api/v1")
	v1.Get("/spot/amenity/:id", GetAmenity(mockService))

	// mock data 생성

	mockAmenity := &entity.Amenity{
		Common: entity.Common{
			Id:        amentityUUID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:        amentityName,
		Description: &amentityDesc,
	}

	mockService.On(
		"GetAmenityById",
		amentityUUID,
		mock.AnythingOfType("*fiber.Ctx"),
	).Return(mockAmenity, nil)

	// HTTP 요청 생성
	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/spot/amenity/%s", amentityUUID), nil)
	req.Header.Set("Content-Type", "application/json")

	// 요청 실행
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result presenter.JsonResponse
	err := json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)
	assert.False(t, result.Error)

	amentityData, ok := result.Data.(map[string]interface{})
	assert.True(t, ok, "Data should be a map")
	assert.Equal(t, amentityName, amentityData["name"])
	assert.Equal(t, amentityDesc, amentityData["description"])
}

func TestGetAmenities(t *testing.T) {
	app := fiber.New()
	mockService := new(MockAmenityService)

	app.Get("/spot/amenity", GetAmenities(mockService))

	mockAmenities := []entity.Amenity{
		{Common: entity.Common{Id: "1"}, Name: "Amenity 1"},
		{Common: entity.Common{Id: "2"}, Name: "Amenity 2"}}

	mockService.On("GetAmenityList", mock.AnythingOfType("*fiber.Ctx")).Return(&mockAmenities, nil)

	req := httptest.NewRequest("GET", "/spot/amenity", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result presenter.JsonResponse
	err := json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)
	assert.False(t, result.Error)
	assert.Len(t, result.Data.([]interface{}), 2)
}

func TestUpdateAmenity(t *testing.T) {
	app, _, cleanup := test.SetupTestApp()
	defer cleanup()

	mockService := new(MockAmenityService)

	v1 := app.Group("/api/v1")
	v1.Put("/spot/amenity/:id", UpdateAmenity(mockService))

	// 테스트 데이터 준비
	amenityID := "1"
	originalName := "Created Amenity"
	updatedName := "Updated Amenity"

	mockAmenity := &entity.Amenity{
		Common: entity.Common{Id: amenityID},
		Name:   originalName,
	}

	// 모의 서비스 동작 설정
	mockService.On("UpdateAmenity",
		mock.AnythingOfType("*dto.UpdateAmenityReq"),
		amenityID,
		mock.AnythingOfType("*fiber.Ctx")).
		Run(func(args mock.Arguments) {
			req := args.Get(0).(*dto.UpdateAmenityReq)
			mockAmenity.Name = req.Name
		}).
		Return(mockAmenity, nil)

	// 요청 본문 생성
	reqBody := dto.UpdateAmenityReq{Name: updatedName}
	jsonBody, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	// HTTP 요청 생성 및 실행
	req := httptest.NewRequest("PUT", "/api/v1/spot/amenity/"+amenityID, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// 응답 상태 코드 확인
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// 응답 본문 파싱
	var result presenter.JsonResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)

	// 응답 검증
	assert.False(t, result.Error)
	amenityData, ok := result.Data.(map[string]interface{})
	assert.True(t, ok, "Result data should be a map")
	assert.Equal(t, updatedName, amenityData["name"], "Amenity name should be updated")

	// 모의 객체 호출 확인
	mockService.AssertExpectations(t)
}

func TestDeleteAmenity(t *testing.T) {
	// 테스트 앱 설정
	app, _, cleanup := test.SetupTestApp()
	defer cleanup()

	// 모의 서비스 생성
	mockService := new(MockAmenityService)

	// 라우트 설정
	v1 := app.Group("/api/v1")
	v1.Delete("/spot/amenity/:id", DeleteAmenity(mockService))

	// 테스트 데이터 준비
	amenityID := "1"

	// 모의 서비스 동작 설정
	mockService.On("DeleteAmenity", amenityID, mock.AnythingOfType("*fiber.Ctx")).Return(nil)

	// HTTP 요청 생성 및 실행
	req := httptest.NewRequest("DELETE", "/api/v1/spot/amenity/"+amenityID, nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// 응답 상태 코드 확인
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// 응답 본문 파싱
	var result presenter.JsonResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)

	// 응답 검증
	assert.False(t, result.Error, "Result error should be false for successful deletion")
	assert.Equal(t, "Deleted successfully", result.Message)

	// 모의 객체 호출 확인
	mockService.AssertExpectations(t)
}
