package services_test

import (
	"testing"
	"tet/internal/mocks/repomocks"
	"tet/internal/models"
	"tet/internal/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductService_Insert(t *testing.T) {
	repo := &repomocks.ProductRepositoryInterface{}
	repo.On("Add", mock.AnythingOfType("models.Product")).
		Return(nil).
		Once()

	service := services.NewProductService(repo)

	err := service.Insert("2f1afe98-63c4-4f59-bcaf-1df835602bdb", models.InsertProductDTO{
		Name:  "Macbook",
		Price: 20500,
		Stock: 10,
	})

	assert.Nil(t, err)
}
