package rv

import (
	"context"
	"fmt"
)

type repositoryMock struct {
	Rental     *Rental
	RentalList []Rental
	CalledWith RentalFilter
}

func (m *repositoryMock) FindById(_ context.Context, _ int64) (Rental, error) {
	if m.Rental != nil {
		return *m.Rental, nil
	} else {
		return Rental{}, fmt.Errorf(noRowsErr)
	}

}

func (m *repositoryMock) FindByFilter(_ context.Context, filter RentalFilter) ([]Rental, error) {
	m.CalledWith = filter
	return m.RentalList, nil
}
