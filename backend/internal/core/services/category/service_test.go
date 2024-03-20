package category

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type mockedCategoryRepository struct {
	mock.Mock
}

func (m *mockedCategoryRepository) Create(ctx context.Context, c domain.Category) (domain.Category, error) {
	args := m.Called(ctx, c)
	return args.Get(0).(domain.Category), args.Error(1)
}

func (m *mockedCategoryRepository) Get(ctx context.Context, id uint) (domain.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Category), args.Error(1)
}

func (m *mockedCategoryRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockedCategoryRepository) List(ctx context.Context, ids *[]uint, filters *map[string]interface{}) ([]domain.Category, error) {
	args := m.Called(ctx, ids, filters)
	return args.Get(0).([]domain.Category), args.Error(1)
}

func (m *mockedCategoryRepository) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	args := m.Called(ctx, id, filters)
	return args.Error(0)
}

func (m *mockedCategoryRepository) Share(ctx context.Context, id uint, u domain.User) error {
	args := m.Called(ctx, id, u)
	return args.Error(0)
}

func (m *mockedCategoryRepository) Unshare(ctx context.Context, id uint, u domain.User) error {
	args := m.Called(ctx, id, u)
	return args.Error(0)
}

func (m *mockedCategoryRepository) GetSummary(ctx context.Context, id uint) ([]domain.CategorySummary, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]domain.CategorySummary), args.Error(1)
}

func (m *mockedCategoryRepository) ListCategoryUsers(ctx context.Context, id uint) ([]domain.CategoryUser, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]domain.CategoryUser), args.Error(1)
}

type ServiceTestSuite struct {
	suite.Suite
	service        Service
	mockRepository *mockedCategoryRepository
	category       domain.Category
	user           domain.User
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.mockRepository = new(mockedCategoryRepository)
	suite.service = NewCategoryService(suite.mockRepository, nil)
	suite.category = domain.Category{
		ID:   1,
		Name: "test",
	}
	suite.user = domain.User{
		ID: 1,
	}
}

func (suite *ServiceTestSuite) TestCreate() {
	// Test case 1: successful creation
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Create", ctx, suite.category).Return(suite.category, nil).Once()

	// Act
	c, err := suite.service.Create(ctx, suite.category)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), suite.category, c, "Category should be equal")

	// Test case 2: failed creation
	// Arrange
	suite.mockRepository.On("Create", ctx, suite.category).Return(domain.Category{}, assert.AnError).Once()

	// Act
	c, err = suite.service.Create(ctx, suite.category)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), domain.Category{}, c, "Category should be empty")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestGet() {
	// Test case 1: successful retrieval
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Get", ctx, suite.category.ID).Return(suite.category, nil).Once()

	// Act
	c, err := suite.service.Get(ctx, suite.category.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), suite.category, c, "Category should be equal")

	// Test case 2: failed retrieval
	// Arrange
	suite.mockRepository.On("Get", ctx, suite.category.ID).Return(domain.Category{}, assert.AnError).Once()

	// Act
	c, err = suite.service.Get(ctx, suite.category.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), domain.Category{}, c, "Category should be empty")

	// Check if the expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestDelete() {
	// Test case 1: successful deletion
	// Arrange
	ctx := context.Background()
	id := uint(1)
	suite.mockRepository.On("Delete", ctx, id).Return(nil).Once()

	// Act
	err := suite.service.Delete(ctx, id)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed deletion
	// Arrange
	suite.mockRepository.On("Delete", ctx, id).Return(assert.AnError).Once()

	// Act
	err = suite.service.Delete(ctx, id)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if the expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestList() {
	// Test case 1: successful listing
	// Arrange
	ctx := context.Background()
	categories := []domain.Category{suite.category}
	ids := []uint{suite.category.ID}
	suite.mockRepository.On("List", ctx, &ids, (*map[string]interface{})(nil)).Return(categories, nil).Once()

	// Act
	c, err := suite.service.List(ctx, &[]uint{1}, nil)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), categories, c, "Categories should be equal")

	// Test case 2: failed listing
	// Arrange
	suite.mockRepository.On("List", ctx, (*[]uint)(nil), (*map[string]interface{})(nil)).Return([]domain.Category{}, assert.AnError).Once()

	// Act
	c, err = suite.service.List(ctx, nil, nil)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), []domain.Category{}, c, "Categories should be empty")

	// Check if the expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestUpdate() {
	// Test case 1: successful update
	// Arrange
	ctx := context.Background()
	id := uint(1)
	suite.mockRepository.On("Update", ctx, id, (*map[string]interface{})(nil)).Return(nil).Once()

	// Act
	err := suite.service.Update(ctx, id, nil)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed update
	// Arrange
	suite.mockRepository.On("Update", ctx, id, (*map[string]interface{})(nil)).Return(assert.AnError).Once()

	// Act
	err = suite.service.Update(ctx, id, nil)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if the expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestShare() {
	// Test case 1: successful sharing
	// Arrange
	ctx := context.Background()
	id := uint(1)
	suite.mockRepository.On("Share", ctx, id, suite.user).Return(nil).Once()

	// Act
	err := suite.service.Share(ctx, id, suite.user)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed sharing
	// Arrange
	suite.mockRepository.On("Share", ctx, id, suite.user).Return(assert.AnError).Once()

	// Act
	err = suite.service.Share(ctx, id, suite.user)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if the expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestUnshare() {
	// Test case 1: successful unsharing
	// Arrange
	ctx := context.Background()
	id := uint(1)
	suite.mockRepository.On("Unshare", ctx, id, suite.user).Return(nil).Once()

	// Act
	err := suite.service.Unshare(ctx, id, suite.user)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed unsharing
	// Arrange
	suite.mockRepository.On("Unshare", ctx, id, suite.user).Return(assert.AnError).Once()

	// Act
	err = suite.service.Unshare(ctx, id, suite.user)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if the expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestGetSummary() {
	// Test case 1: successful retrieval
	// Arrange
	ctx := context.Background()
	id := uint(1)
	summary := []domain.CategorySummary{
		{
			ID: 1,
		},
	}
	suite.mockRepository.On("GetSummary", ctx, id).Return(summary, nil).Once()

	// Act
	s, err := suite.service.GetSummary(ctx, id)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), summary, s, "Summary should be equal")

	// Test case 2: failed retrieval
	// Arrange
	suite.mockRepository.On("GetSummary", ctx, id).Return([]domain.CategorySummary{}, assert.AnError).Once()

	// Act
	s, err = suite.service.GetSummary(ctx, id)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), []domain.CategorySummary{}, s, "Summary should be empty")

	// Check if the expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestListCategoryUsers() {
	// Test case 1: successful listing
	// Arrange
	ctx := context.Background()
	id := uint(1)
	users := []domain.CategoryUser{
		{
			UserID: 1,
		},
	}
	suite.mockRepository.On("ListCategoryUsers", ctx, id).Return(users, nil).Once()

	// Act
	u, err := suite.service.ListCategoryUsers(ctx, id)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), users, u, "Users should be equal")

	// Test case 2: failed listing
	// Arrange
	suite.mockRepository.On("ListCategoryUsers", ctx, id).Return([]domain.CategoryUser{}, assert.AnError).Once()

	// Act
	u, err = suite.service.ListCategoryUsers(ctx, id)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), []domain.CategoryUser{}, u, "Users should be empty")

	// Check if the expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
