package rv

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *Server) HandleGetRental(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rentalId, convErr := strconv.Atoi(mux.Vars(r)["id"])

	if convErr != nil {
		writeError(http.StatusBadRequest, "invalid rental id", w)
		return
	}

	rental, err := s.repository.FindById(ctx, int64(rentalId))

	if err != nil {
		if err.Error() == noRowsErr {
			writeError(http.StatusNotFound, "rental not found", w)
			return
		} else {
			writeError(http.StatusInternalServerError, "error while getting rental", w)
			return
		}
	}

	writeJSON(ToDTO(rental), w)
}

func (s *Server) HandleSearchRentals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var filter RentalFilter
	urlErr := s.decoder.Decode(&filter, r.URL.Query())
	if urlErr != nil {
		writeError(http.StatusBadRequest, fmt.Sprintf("invalid filter format: %s", urlErr.Error()), w)
		return
	}

	rentals, err := s.repository.FindByFilter(ctx, filter)

	if err != nil {
		if errors.Is(err, invalidFilterValuesErr) {
			writeError(http.StatusBadRequest, err.Error(), w)
			return
		} else {
			writeError(http.StatusInternalServerError, "error while getting rental", w)
			return
		}
	}

	writeJSON(ToDTOs(rentals), w)
}
