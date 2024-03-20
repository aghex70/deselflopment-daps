package note

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type mockNoteRepository struct {
	mock.Mock
}

func (m *mockNoteRepository) Create(ctx context.Context, n domain.Note) (domain.Note, error) {
	args := m.Called(ctx, n)
	return args.Get(0).(domain.Note), args.Error(1)
}

func (m *mockNoteRepository) Get(ctx context.Context, id uint) (domain.Note, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Note), args.Error(1)
}

func (m *mockNoteRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockNoteRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Note, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).([]domain.Note), args.Error(1)
}

func (m *mockNoteRepository) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	args := m.Called(ctx, id, filters)
	return args.Error(0)
}

func (m *mockNoteRepository) Share(ctx context.Context, id uint, u domain.User) error {
	args := m.Called(ctx, id, u)
	return args.Error(0)
}

func (m *mockNoteRepository) Unshare(ctx context.Context, id uint, u domain.User) error {
	args := m.Called(ctx, id, u)
	return args.Error(0)
}

type ServiceTestSuite struct {
	suite.Suite
	service        Service
	mockRepository *mockNoteRepository
	note           domain.Note
	user           domain.User
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.mockRepository = new(mockNoteRepository)
	suite.service = NewNoteService(suite.mockRepository, nil)
	suite.note = domain.Note{
		ID:      1,
		Content: "Test",
	}
	suite.user = domain.User{
		ID: 1,
	}
}

func (suite *ServiceTestSuite) TestCreate() {
	// Test case 1: successful creation
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Create", ctx, suite.note).Return(suite.note, nil).Once()

	// Act
	n, err := suite.service.Create(ctx, suite.note)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), suite.note, n, "Note should be equal")

	// Test case 2: failed creation
	// Arrange
	suite.mockRepository.On("Create", ctx, suite.note).Return(domain.Note{}, assert.AnError).Once()

	// Act
	n, err = suite.service.Create(ctx, suite.note)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), domain.Note{}, n, "Note should be empty")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestGet() {
	// Test case 1: successful retrieval
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Get", ctx, suite.note.ID).Return(suite.note, nil).Once()

	// Act
	n, err := suite.service.Get(ctx, suite.note.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), suite.note, n, "Note should be equal")

	// Test case 2: failed retrieval
	// Arrange
	suite.mockRepository.On("Get", ctx, suite.note.ID).Return(domain.Note{}, assert.AnError).Once()

	// Act
	n, err = suite.service.Get(ctx, suite.note.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), domain.Note{}, n, "Note should be empty")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestDelete() {
	// Test case 1: successful deletion
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Delete", ctx, suite.note.ID).Return(nil).Once()

	// Act
	err := suite.service.Delete(ctx, suite.note.ID)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed deletion
	// Arrange
	suite.mockRepository.On("Delete", ctx, suite.note.ID).Return(assert.AnError).Once()

	// Act
	err = suite.service.Delete(ctx, suite.note.ID)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestList() {
	// Test case 1: successful retrieval
	// Arrange
	ctx := context.Background()
	notes := []domain.Note{suite.note}
	suite.mockRepository.On("List", ctx, &map[string]interface{}{}).Return(notes, nil).Once()

	// Act
	n, err := suite.service.List(ctx, &map[string]interface{}{})

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), notes, n, "Notes should be equal")

	// Test case 2: failed listing
	// Arrange
	suite.mockRepository.On("List", ctx, &map[string]interface{}{}).Return([]domain.Note{}, assert.AnError).Once()

	// Act
	n, err = suite.service.List(ctx, &map[string]interface{}{})

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), []domain.Note{}, n, "Notes should be empty")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestUpdate() {
	// Test case 1: successful update
	// Arrange
	ctx := context.Background()
	id := uint(1)
	fields := map[string]interface{}{"content": "Updated"}
	suite.mockRepository.On("Update", ctx, id, &fields).Return(nil).Once()

	// Act
	err := suite.service.Update(ctx, id, &fields)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed update
	// Arrange
	suite.mockRepository.On("Update", ctx, id, &fields).Return(assert.AnError).Once()

	// Act
	err = suite.service.Update(ctx, id, &fields)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestShare() {
	// Test case 1: successful sharing
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Share", ctx, suite.note.ID, suite.user).Return(nil).Once()

	// Act
	err := suite.service.Share(ctx, suite.note.ID, suite.user)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed sharing
	// Arrange
	suite.mockRepository.On("Share", ctx, suite.note.ID, suite.user).Return(assert.AnError).Once()

	// Act
	err = suite.service.Share(ctx, suite.note.ID, suite.user)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestUnshare() {
	// Test case 1: successful unsharing
	// Arrange
	ctx := context.Background()
	suite.mockRepository.On("Unshare", ctx, suite.note.ID, suite.user).Return(nil).Once()

	// Act
	err := suite.service.Unshare(ctx, suite.note.ID, suite.user)

	// Assert
	assert.NoError(suite.T(), err, "Error should be nil")

	// Test case 2: failed unsharing
	// Arrange
	suite.mockRepository.On("Unshare", ctx, suite.note.ID, suite.user).Return(assert.AnError).Once()

	// Act
	err = suite.service.Unshare(ctx, suite.note.ID, suite.user)

	// Assert
	assert.Error(suite.T(), err, "Error should not be nil")

	// Check if all expectations were met
	suite.mockRepository.AssertExpectations(suite.T())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
