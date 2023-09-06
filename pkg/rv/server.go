package rv

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"net/http"
)

type Server struct {
	repository Repository
	router     *mux.Router
	decoder    *schema.Decoder
}

type Repository interface {
	FindById(ctx context.Context, id int64) (Rental, error)
	FindByFilter(ctx context.Context, filter RentalFilter) ([]Rental, error)
}

type RentalFilter struct {
	PriceMin int64     `schema:"price_min"`
	PriceMax int64     `schema:"price_max"`
	IDs      []int64   `schema:"ids"`
	Location []float64 `schema:"near"`

	Limit  int `schema:"limit"`
	Offset int `schema:"offset"`

	Sort string `schema:"ids"`
}

func NewServer(repository Repository) *Server {
	server := Server{
		repository: repository,
		router:     mux.NewRouter(),
		decoder:    schema.NewDecoder(),
	}

	server.router.HandleFunc("/rentals/{id}", server.HandleGetRental).Methods(http.MethodGet)
	server.router.HandleFunc("/rentals", server.HandleSearchRentals).Methods(http.MethodGet)

	return &server
}

func (s *Server) Start(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.router)
}

func writeJSON(payload interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytePayload, _ := json.Marshal(payload)
	w.Write(bytePayload)
}

func writeError(statusCode int, message string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
