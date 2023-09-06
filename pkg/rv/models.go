package rv

import (
	"github.com/uptrace/bun"
	"time"
)

type Rental struct {
	bun.BaseModel   `bun:"table:rentals,alias:r"`
	ID              int64 `bun:"id,pk"`
	UserId          int64
	User            User    `bun:"rel:has-one,join:user_id=id"`
	Name            string  `bun:"name"`
	Type            string  `bun:"type"`
	Description     string  `bun:"description"`
	Sleeps          int64   `bun:"sleeps"`
	PricePerDay     int64   `bun:"price_per_day"`
	HomeCity        string  `bun:"home_city"`
	HomeState       string  `bun:"home_state"`
	HomeZip         string  `bun:"home_zip"`
	HomeCountry     string  `bun:"home_country"`
	VehicleMake     string  `bun:"vehicle_make"`
	VehicleModel    string  `bun:"vehicle_model"`
	VehicleYear     int64   `bun:"vehicle_year"`
	VehicleLength   float64 `bun:"vehicle_length"`
	Created         time.Time
	Updated         time.Time
	Lat             float64 `bun:"lat"`
	Lng             float64 `bun:"lng"`
	PrimaryImageUrl string  `bun:"primary_image_url"`
}

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64  `bun:"id,pk"`
	FirstName     string `bun:"first_name"`
	LastName      string `bun:"last_name"`
}
