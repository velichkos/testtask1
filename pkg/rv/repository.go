package rv

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const nearDiscanceMeters = 160934 // 100 miles
const defaultLimit = 20

const noRowsErr = "sql: no rows in result set"

var invalidFilterValuesErr = errors.New("invalid filter values")

var sortColumns = map[string]string{
	"id":    "id",
	"price": "price_per_day",
	"name":  "name",
	"type":  "type",
	"model": "vehicle_make",
	"year":  "vehicle_year",
}

type RepositoryImpl struct {
	db *bun.DB
}

func NewRepository(connectionString string) *RepositoryImpl {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))

	db := bun.NewDB(sqldb, pgdialect.New())

	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) FindById(ctx context.Context, id int64) (Rental, error) {
	result := Rental{}

	err := r.db.NewSelect().Model(&result).Relation("User").Where("r.id = ?", id).Scan(ctx)
	if err != nil {
		return Rental{}, err
	}

	return result, nil
}

func (r *RepositoryImpl) FindByFilter(ctx context.Context, filter RentalFilter) ([]Rental, error) {
	var result []Rental

	if err := filter.Validate(); err != nil {
		return nil, err
	}

	query := r.db.NewSelect().Model(&result).Relation("User")
	query = applyFilter(query, filter)

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func applyFilter(query *bun.SelectQuery, filter RentalFilter) *bun.SelectQuery {
	if filter.PriceMin != 0 {
		query = query.Where("r.price_per_day >= ?", filter.PriceMin)
	}
	if filter.PriceMax != 0 {
		query = query.Where("r.price_per_day <= ?", filter.PriceMax)
	}
	if filter.IDs != nil {
		query = query.Where("r.id in (?)", bun.In(filter.IDs))
	}
	if filter.Location != nil {
		query = query.Where("ST_DWithin(ST_GeogFromText(format('POINT(%s %s)', r.lat, r.lng)), ST_GeogFromText('POINT(? ?)'), ?)",
			filter.Location[0], filter.Location[1], nearDiscanceMeters)
	}

	if filter.Offset != 0 {
		query = query.Offset(filter.Offset)
	}

	if filter.Limit != 0 {
		query = query.Limit(filter.Limit)
	} else {
		query = query.Limit(defaultLimit)
	}

	if filter.Sort != "" {
		query.Order(sortColumns[filter.Sort])
	}

	return query
}

func (f RentalFilter) Validate() error {
	if f.PriceMax < f.PriceMin {
		return fmt.Errorf("%w: price max should be higher than price min", invalidFilterValuesErr)
	}
	if f.Location != nil && len(f.Location) != 2 {
		return fmt.Errorf("%w: near search should have lat and lng", invalidFilterValuesErr)
	}
	if f.Sort != "" {
		_, isSupported := sortColumns[f.Sort]
		if !isSupported {
			return fmt.Errorf("%w: unsupported sort filter", invalidFilterValuesErr)
		}
	}
	return nil
}
