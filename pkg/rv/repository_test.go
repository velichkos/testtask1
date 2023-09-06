package rv

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func testRepository() *RepositoryImpl {
	testConnectionString := "postgres://root:root@localhost:5434/testingwithrentals?sslmode=disable"
	return NewRepository(testConnectionString)
}

func TestRepository_FindById(t *testing.T) {
	repository := testRepository()
	rental, err := repository.FindById(context.Background(), 10)
	require.Nil(t, err)
	assert.Equal(t, int64(10), rental.ID)
	assert.Equal(t, "camper-van", rental.Type)
	assert.Equal(t, "Ben", rental.User.FirstName)
}

func TestRepository_FindById_NotFound(t *testing.T) {
	repository := testRepository()
	_, err := repository.FindById(context.Background(), 400)
	require.Equal(t, "sql: no rows in result set", err.Error())
}

func TestRepository_FindByFilter_Prise(t *testing.T) {
	repository := testRepository()
	filter := RentalFilter{
		PriceMax: 17000,
		Offset:   1,
		Limit:    5,
	}

	rentals, err := repository.FindByFilter(context.Background(), filter)
	require.Nil(t, err)
	require.Len(t, rentals, 5)
	assert.Equal(t, int64(2), rentals[0].ID)
	assert.Equal(t, int64(4), rentals[1].ID)
	assert.Equal(t, int64(5), rentals[2].ID)
	assert.Equal(t, int64(6), rentals[3].ID)
	assert.Equal(t, int64(7), rentals[4].ID)
}
func TestRepository_FindByFilter_IDs(t *testing.T) {
	repository := testRepository()
	filter := RentalFilter{
		IDs:   []int64{2, 4, 6, 10, 17},
		Limit: 5,
		Sort:  "id",
	}

	rentals, err := repository.FindByFilter(context.Background(), filter)
	require.Nil(t, err)
	require.Len(t, rentals, 5)
	assert.Equal(t, int64(2), rentals[0].ID)
	assert.Equal(t, int64(4), rentals[1].ID)
	assert.Equal(t, int64(6), rentals[2].ID)
	assert.Equal(t, int64(10), rentals[3].ID)
	assert.Equal(t, int64(17), rentals[4].ID)
}

func TestRepository_FindByFilter_Geo(t *testing.T) {
	repository := testRepository()
	filter := RentalFilter{
		Location: []float64{45.51, -122.68},
		Limit:    5,
	}

	rentals, err := repository.FindByFilter(context.Background(), filter)
	require.Nil(t, err)
	require.Len(t, rentals, 3)
	assert.Equal(t, int64(2), rentals[0].ID)
	assert.Equal(t, int64(16), rentals[1].ID)
	assert.Equal(t, int64(27), rentals[2].ID)
}
