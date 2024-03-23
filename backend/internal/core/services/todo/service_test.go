package todo

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"mime/multipart"
	"testing"
)

type mockTodoRepository struct {
	mock.Mock
}

func (m *mockTodoRepository) Create(ctx context.Context, t domain.Todo) (domain.Todo, error) {
	args := m.Called(ctx, t)
	return args.Get(0).(domain.Todo), args.Error(1)
}

func (m *mockTodoRepository) Get(ctx context.Context, id uint) (domain.Todo, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Todo), args.Error(1)
}

func (m *mockTodoRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockTodoRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Todo, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).([]domain.Todo), args.Error(1)
}

func (m *mockTodoRepository) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	args := m.Called(ctx, id, filters)
	return args.Error(0)
}

func (m *mockTodoRepository) Start(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockTodoRepository) Complete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockTodoRepository) Restart(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockTodoRepository) Activate(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockTodoRepository) Import(ctx context.Context, f multipart.File) error {
	args := m.Called(ctx, f)
	return args.Error(0)
}

type ServiceTestSuite struct {
	suite.Suite
	service        Service
	mockRepository *mockTodoRepository
	todo           domain.Todo
	user           domain.User
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.mockRepository = new(mockTodoRepository)
	suite.service = NewTodoService(suite.mockRepository, nil)
	suite.todo = domain.Todo{
		ID:   1,
		Name: "Test",
	}
	suite.user = domain.User{
		ID: 1,
	}
}

func (suite *ServiceTestSuite) TestCreate() {
	// Test case 1: successful creation
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Create", ctx, suite.todo).Return(suite.todo, nil).Once()

	// Act
	t, err := suite.service.Create(ctx, suite.todo)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), suite.todo, t, "Todo should be equal")

	// Test case 2: failed creation
	// Arrange
	suite.mockRepository.On("Create", ctx, suite.todo).Return(domain.Todo{}, assert.AnError).Once()

	// Act
	t, err = suite.service.Create(ctx, suite.todo)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), domain.Todo{}, t, "Todo should be empty")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestGet() {
	// Test case 1: successful retrieval
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Get", ctx, suite.todo.ID).Return(suite.todo, nil).Once()

	// Act
	t, err := suite.service.Get(ctx, suite.todo.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), suite.todo, t, "Todo should be equal")

	// Test case 2: failed retrieval
	// Arrange
	suite.mockRepository.On("Get", ctx, suite.todo.ID).Return(domain.Todo{}, assert.AnError).Once()

	// Act
	t, err = suite.service.Get(ctx, suite.todo.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), domain.Todo{}, t, "Todo should be empty")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestDelete() {
	// Test case 1: successful deletion
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Delete", ctx, suite.todo.ID).Return(nil).Once()

	// Act
	err := suite.service.Delete(ctx, suite.todo.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed deletion
	// Arrange
	suite.mockRepository.On("Delete", ctx, suite.todo.ID).Return(assert.AnError).Once()

	// Act
	err = suite.service.Delete(ctx, suite.todo.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestList() {
	// Test case 1: successful listing
	// Arrange
	ctx := context.Background()
	filters := &map[string]interface{}{}
	suite.mockRepository.On("List", ctx, filters).Return([]domain.Todo{suite.todo}, nil).Once()

	// Act
	todos, err := suite.service.List(ctx, filters)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), []domain.Todo{suite.todo}, todos, "Todos should be equal")

	// Test case 2: failed listing
	// Arrange
	suite.mockRepository.On("List", ctx, filters).Return([]domain.Todo{}, assert.AnError).Once()

	// Act
	todos, err = suite.service.List(ctx, filters)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), []domain.Todo{}, todos, "Todos should be empty")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestUpdate() {
	// Test case 1: successful update
	// Arrange
	ctx := context.Background()
	filters := &map[string]interface{}{}
	suite.mockRepository.On("Update", ctx, suite.todo.ID, filters).Return(nil).Once()

	// Act
	err := suite.service.Update(ctx, suite.todo.ID, filters)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed update
	// Arrange
	suite.mockRepository.On("Update", ctx, suite.todo.ID, filters).Return(assert.AnError).Once()

	// Act
	err = suite.service.Update(ctx, suite.todo.ID, filters)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestStart() {
	// Test case 1: successful start
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Start", ctx, suite.todo.ID).Return(nil).Once()

	// Act
	err := suite.service.Start(ctx, suite.todo.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed start
	// Arrange
	suite.mockRepository.On("Start", ctx, suite.todo.ID).Return(assert.AnError).Once()

	// Act
	err = suite.service.Start(ctx, suite.todo.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestComplete() {
	// Test case 1: successful completion
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Complete", ctx, suite.todo.ID).Return(nil).Once()

	// Act
	err := suite.service.Complete(ctx, suite.todo.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed completion
	// Arrange
	suite.mockRepository.On("Complete", ctx, suite.todo.ID).Return(assert.AnError).Once()

	// Act
	err = suite.service.Complete(ctx, suite.todo.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestRestart() {
	// Test case 1: successful restart
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Restart", ctx, suite.todo.ID).Return(nil).Once()

	// Act
	err := suite.service.Restart(ctx, suite.todo.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed restart
	// Arrange
	suite.mockRepository.On("Restart", ctx, suite.todo.ID).Return(assert.AnError).Once()

	// Act
	err = suite.service.Restart(ctx, suite.todo.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestActivate() {
	// Test case 1: successful activation
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Activate", ctx, suite.todo.ID).Return(nil).Once()

	// Act
	err := suite.service.Activate(ctx, suite.todo.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed activation
	// Arrange
	suite.mockRepository.On("Activate", ctx, suite.todo.ID).Return(assert.AnError).Once()

	// Act
	err = suite.service.Activate(ctx, suite.todo.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestImport() {
	// Test case 1: successful import
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Import", ctx, mock.Anything).Return(nil).Once()

	// Act
	err := suite.service.Import(ctx, nil)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed import
	// Arrange
	suite.mockRepository.On("Import", ctx, mock.Anything).Return(assert.AnError).Once()

	// Act
	err = suite.service.Import(ctx, nil)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
