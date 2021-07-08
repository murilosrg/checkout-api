package products

import (
	"github.com/murilosrg/checkout-api/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testDB *db.DB
)

func TestRepository_FindProduct(t *testing.T) {
	loadDB()
	repo := NewRepository(testDB, logrus.New())

	actual, err := repo.Find(1)
	assert.NotNil(t, actual)
	assert.Equal(t, 1, actual.ID)
	assert.Nil(t, err)
}

func TestRepository_FindProduct_NotFound(t *testing.T) {
	loadDB()
	repo := NewRepository(testDB, logrus.New())

	actual, err := repo.Find(9999)
	assert.Equal(t, 0, actual.ID)
	assert.NotNil(t, err)
}

func TestRepository_FindProduct_GetGift(t *testing.T) {
	loadDB()
	repo := NewRepository(testDB, logrus.New())

	actual, err := repo.GetGift()
	assert.NotNil(t, actual)
	assert.Equal(t, true, actual.IsGift)
	assert.Nil(t, err)
}

func loadDB() {
	if testDB == nil {
		testDB, _ = db.InitializeDb("../../products.json")
	}
}
