package categoryhandler

import (
	"bytes"
	"camping-backend-with-go/internal/application/dto"
	"camping-backend-with-go/internal/domain/entity"
	"camping-backend-with-go/internal/domain/presenter"
	"camping-backend-with-go/pkg/util"
	"camping-backend-with-go/test"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) GetCategoryList(context ...*fiber.Ctx) (*[]entity.Category, error) {
	c, _ := util.ContextParser(context...)
	args := m.Called(c)
	return args.Get(0).(*[]entity.Category), args.Error(1)
}
func (m *MockCategoryService) CreateCategory(input *dto.CreateCategoryReq, context ...*fiber.Ctx) (*entity.Category, error) {
	c, _ := util.ContextParser(context...)
	args := m.Called(input, c)
	return args.Get(0).(*entity.Category), args.Error(1)
}
func (m *MockCategoryService) UpdateCategory(input *dto.UpdateCategoryReq, id string, context ...*fiber.Ctx) (*entity.Category, error) {
	c, _ := util.ContextParser(context...)
	args := m.Called(input, id, c)
	return args.Get(0).(*entity.Category), args.Error(1)
}
func (m *MockCategoryService) DeleteCategory(id string, context ...*fiber.Ctx) error {
	c, _ := util.ContextParser(context...)
	args := m.Called(id, c)
	return args.Error(0)
}
func (m *MockCategoryService) GetCategoryById(id string, context ...*fiber.Ctx) (*entity.Category, error) {
	c, _ := util.ContextParser(context...)
	args := m.Called(id, c)
	return args.Get(0).(*entity.Category), args.Error(1)
}

type mockSetup interface {
	setup(*MockCategoryService)
}

type defaultMockSetup struct {
	id, name string
}

type updateMockSetup struct {
	id, name string
}

func (d *updateMockSetup) setup(m *MockCategoryService) {
	category := &entity.Category{
		Common: entity.Common{
			Id: d.id,
		},
		Name: d.name,
		// Spots:  []entity.Spot{},
	}
	m.On(
		"UpdateCategory",
		mock.AnythingOfType("*dto.UpdateCategoryReq"),
		d.id,
		mock.AnythingOfType("*fiber.Ctx"),
	).Run(func(args mock.Arguments) {
		req := args.Get(0).(*dto.UpdateCategoryReq)
		category.Name = *req.Name
	}).Return(category, nil).Once()
}

func (d *defaultMockSetup) setup(m *MockCategoryService) {
	m.On(
		"CreateCategory",
		mock.AnythingOfType("*dto.CreateCategoryReq"),
		mock.AnythingOfType("*fiber.Ctx"),
	).Return(&entity.Category{
		Common: entity.Common{
			Id: d.id,
		},
		Name: d.name,
		// Spots:  []entity.Spot{},
	}, nil).Once()
}

type getCategoryByIdMock struct {
	id, name string
}

type getAllCategoriesMock struct {
	categories []entity.Category
}

func (g *getCategoryByIdMock) setup(m *MockCategoryService) {
	m.On("GetCategoryById", g.id, mock.AnythingOfType("*fiber.Ctx")).
		Return(&entity.Category{
			Common: entity.Common{Id: g.id},
			Name:   g.name,
		}, nil).Once()

}

func (g *getAllCategoriesMock) setup(m *MockCategoryService) {
	m.On("GetCategoryList", mock.AnythingOfType("*fiber.Ctx")).
		Return(&g.categories, nil).Once()
}

func TestCreateCategory(t *testing.T) {

	// Test Data
	categoryId := "1"
	categoryName := "categoryName"

	tests := []struct {
		name               string
		setupMock          mockSetup
		expectedStatusCode int
		expectedError      bool
		expectedMsg        string
	}{
		{
			name:               "성공적인 카테고리 생성",
			setupMock:          &defaultMockSetup{id: categoryId, name: categoryName},
			expectedStatusCode: fiber.StatusOK,
			expectedError:      false,
			expectedMsg:        "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, _, cleanup := test.SetupTestApp()
			defer cleanup()

			mockService := new(MockCategoryService)
			tt.setupMock.setup(mockService)
			v1 := app.Group("/api/v1")
			v1.Post("/category", CreateCategory(mockService))

			reqBody := dto.CreateCategoryReq{
				Name: util.StrPointer(categoryName),
			}

			jsonBody, _ := json.Marshal(reqBody)

			req := httptest.NewRequest("POST", "/api/v1/category", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)

			var result presenter.JsonResponse
			err := json.NewDecoder(resp.Body).Decode(&result)
			assert.NoError(t, err)
			assert.False(t, result.Error)

			categoryData, ok := result.Data.(map[string]interface{})

			assert.True(t, ok, "Data should be convert")
			assert.Equal(t, categoryName, categoryData["name"])
			assert.Equal(t, categoryId, categoryData["id"])
			mockService.AssertExpectations(t)
		})
	}

}

func TestGetCategoryById(t *testing.T) {
	// DB Data
	categoryId := "1"
	categoryName := "category_name"

	tests := []struct {
		name               string
		setupMock          mockSetup
		expectedStatusCode int
		expectedError      bool
		expectedMsg        string
	}{

		{
			name:               "성공적인 한개 조회",
			setupMock:          &getCategoryByIdMock{id: categoryId, name: categoryName},
			expectedStatusCode: fiber.StatusOK,
			expectedError:      false,
			expectedMsg:        "",
		},
	}

	for _, tt := range tests {
		t.Run("성공적인 카테고리 한개 조회", func(t *testing.T) {
			app, _, cleanup := test.SetupTestApp()
			defer cleanup()

			mockService := new(MockCategoryService)
			tt.setupMock.setup(mockService)

			v1 := app.Group("/api/v1")
			v1.Get("/category/:id", GetCategory(mockService))

			req := httptest.NewRequest("GET", "/api/v1/category/"+categoryId, nil)
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)
			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)
		})
	}
}

func TestGetCategories(t *testing.T) {
	// DB Data
	categoryIds := []string{"1", "2"}
	categoryNames := []string{"category_name1", "category_name2"}

	tests := []struct {
		name               string
		setupMock          mockSetup
		expectedStatusCode int
		expectedError      bool
		expectedMsg        string
	}{

		{
			name: "카테고리 여러개 조회",
			setupMock: &getAllCategoriesMock{
				categories: []entity.Category{
					{Common: entity.Common{Id: categoryIds[0]}, Name: categoryNames[0]},
					{Common: entity.Common{Id: categoryIds[1]}, Name: categoryNames[1]},
				},
			},
			expectedStatusCode: fiber.StatusOK,
			expectedError:      false,
			expectedMsg:        "",
		},
	}

	for _, tt := range tests {
		t.Run("성공적인 카테고리들 조회", func(t *testing.T) {
			app, _, cleanup := test.SetupTestApp()
			defer cleanup()

			mockService := new(MockCategoryService)
			tt.setupMock.setup(mockService)

			v1 := app.Group("/api/v1")
			v1.Get("/category", GetCategoryList(mockService))

			req := httptest.NewRequest("GET", "/api/v1/category", nil)
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)

			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)

			var result presenter.JsonResponse

			json.NewDecoder(resp.Body).Decode(&result)
			categoryData, ok := result.Data.([]interface{})

			assert.True(t, ok, "Data should be convert")
			assert.Len(t, categoryData, 2)
			mockService.AssertExpectations(t)

			// 각 카테고리 항목 확인
			for i, category := range categoryData {
				categoryMap, ok := category.(map[string]interface{})
				assert.True(t, ok, "Each category should be a map")
				assert.Equal(t, categoryIds[i], categoryMap["id"])
				assert.Equal(t, categoryNames[i], categoryMap["name"])
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	// Test Data
	categoryId := "1"
	categoryName := "categoryName"
	updatedCategoryName := "updatedCategoryName"

	tests := []struct {
		name               string
		setupMock          mockSetup
		expectedStatusCode int
		expectedError      bool
		expectedMsg        string
	}{
		{
			name:               "업데이트 한개 테스트",
			setupMock:          &updateMockSetup{id: categoryId, name: categoryName},
			expectedStatusCode: fiber.StatusOK,
			expectedError:      false,
			expectedMsg:        "",
		},
	}

	for _, tt := range tests {
		app, _, cleanup := test.SetupTestApp()
		defer cleanup()

		mockService := new(MockCategoryService)
		v1 := app.Group("/api/v1")
		v1.Put("/category/:id", UpdateCategory(mockService))

		reqBody := dto.UpdateCategoryReq{
			Name: util.StrPointer(updatedCategoryName),
		}

		jsonBody, err := json.Marshal(reqBody)
		assert.NoError(t, err)

		req := httptest.NewRequest("PUT", "/api/v1/category/"+categoryId, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)
		var result presenter.JsonResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)

		categoryData, ok := result.Data.(map[string]interface{})
		assert.True(t, ok, "result data should be a map")
		assert.Equal(t, updatedCategoryName, categoryData["name"])
		mockService.AssertExpectations(t)
	}
}
