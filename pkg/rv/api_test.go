package rv

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockedRepository repositoryMock

func testServer() *Server {
	mockedRepository = repositoryMock{}
	return NewServer(&mockedRepository)
}

func TestServer_HandleGetRental(t *testing.T) {
	server := testServer()
	mockedRepository.Rental = &Rental{
		ID:   10,
		Name: "test rental",
		Type: "rental",
	}

	req, err := http.NewRequest(http.MethodGet, "/rentals/10", nil)
	require.Nil(t, err)
	resp := httptest.NewRecorder()

	server.router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseRental RentalDTO
	err = json.Unmarshal(resp.Body.Bytes(), &responseRental)
	require.Nil(t, err)
	assert.Equal(t, int64(10), responseRental.ID)
	assert.Equal(t, "test rental", responseRental.Name)
	assert.Equal(t, "rental", responseRental.Type)
}

func TestServer_HandleGetRental_NotFound(t *testing.T) {
	server := testServer()

	req, err := http.NewRequest(http.MethodGet, "/rentals/10", nil)
	require.Nil(t, err)
	resp := httptest.NewRecorder()

	server.router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestServer_HandleSearchRentals(t *testing.T) {
	server := testServer()

	mockedRepository.RentalList = []Rental{
		{
			ID:   1,
			Name: "test rental 1",
		},
		{
			ID:   2,
			Name: "test rental 2",
		},
	}

	req, err := http.NewRequest(http.MethodGet, "/rentals?limit=5&near=45.51,-122.68", nil)
	require.Nil(t, err)
	resp := httptest.NewRecorder()

	server.router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	assert.Len(t, mockedRepository.CalledWith.Location, 2)
	assert.Equal(t, 5, mockedRepository.CalledWith.Limit)

	var responseRentals []RentalDTO
	err = json.Unmarshal(resp.Body.Bytes(), &responseRentals)
	require.Nil(t, err)
	require.Len(t, responseRentals, 2)
	assert.Equal(t, int64(1), responseRentals[0].ID)
	assert.Equal(t, "test rental 1", responseRentals[0].Name)
	assert.Equal(t, int64(2), responseRentals[1].ID)
	assert.Equal(t, "test rental 2", responseRentals[1].Name)
}
