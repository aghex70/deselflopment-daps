package topic

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type mockTopicRepository struct {
	mock.Mock
}

func (m *mockTopicRepository) Create(ctx context.Context, t domain.Topic) (domain.Topic, error) {
	args := m.Called(ctx, t)
	return args.Get(0).(domain.Topic), args.Error(1)
}

func (m *mockTopicRepository) Get(ctx context.Context, id uint) (domain.Topic, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Topic), args.Error(1)
}

func (m *mockTopicRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockTopicRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Topic, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).([]domain.Topic), args.Error(1)
}

func (m *mockTopicRepository) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	args := m.Called(ctx, id, filters)
	return args.Error(0)
}

type ServiceTestSuite struct {
	suite.Suite
	service        Service
	mockRepository *mockTopicRepository
	topic          domain.Topic
	user           domain.User
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.mockRepository = new(mockTopicRepository)
	suite.service = NewTopicService(suite.mockRepository, nil)
	suite.topic = domain.Topic{
		Name: "test",
	}
	suite.user = domain.User{
		Email: "email",
	}
}

func (suite *ServiceTestSuite) TestCreate() {
	// Test case 1: successful creation
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Create", ctx, suite.topic).Return(suite.topic, nil).Once()

	// Act
	t, err := suite.service.Create(ctx, suite.topic)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), suite.topic, t, "Topic should be equal")

	// Test case 2: failed creation
	// Arrange
	suite.mockRepository.On("Create", ctx, suite.topic).Return(domain.Topic{}, assert.AnError).Once()

	// Act
	t, err = suite.service.Create(ctx, suite.topic)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), domain.Topic{}, t, "Topic should be empty")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestGet() {
	// Test case 1: successful retrieval
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Get", ctx, suite.topic.ID).Return(suite.topic, nil).Once()

	// Act
	t, err := suite.service.Get(ctx, suite.topic.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), suite.topic, t, "Topic should be equal")

	// Test case 2: failed retrieval
	// Arrange
	suite.mockRepository.On("Get", ctx, suite.topic.ID).Return(domain.Topic{}, assert.AnError).Once()

	// Act
	t, err = suite.service.Get(ctx, suite.topic.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), domain.Topic{}, t, "Topic should be empty")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestDelete() {
	// Test case 1: successful deletion
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Delete", ctx, suite.topic.ID).Return(nil).Once()

	// Act
	err := suite.service.Delete(ctx, suite.topic.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed deletion
	// Arrange
	suite.mockRepository.On("Delete", ctx, suite.topic.ID).Return(assert.AnError).Once()

	// Act
	err = suite.service.Delete(ctx, suite.topic.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestList() {
	// Test case 1: successful retrieval
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("List", ctx, nil).Return([]domain.Topic{suite.topic}, nil).Once()

	// Act
	topics, err := suite.service.List(ctx, nil)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), []domain.Topic{suite.topic}, topics, "Topics should be equal")

	// Test case 2: failed retrieval
	// Arrange
	suite.mockRepository.On("List", ctx, nil).Return([]domain.Topic{}, assert.AnError).Once()

	// Act
	topics, err = suite.service.List(ctx, nil)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), []domain.Topic{}, topics, "Topics should be empty")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestUpdate() {
	// Test case 1: successful update
	// Arrange
	ctx := context.Background()
	filters := &map[string]interface{}{
		"name": "new name",
	}
	suite.mockRepository.On("Update", ctx, suite.topic.ID, filters).Return(nil).Once()

	// Act
	err := suite.service.Update(ctx, suite.topic.ID, filters)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed update
	// Arrange
	suite.mockRepository.On("Update", ctx, suite.topic.ID, filters).Return(assert.AnError).Once()

	// Act
	err = suite.service.Update(ctx, suite.topic.ID, filters)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
