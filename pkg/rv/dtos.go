package rv

type RentalDTO struct {
	ID              int64       `json:"id"`
	User            UserDTO     `json:"user"`
	Name            string      `json:"name"`
	Type            string      `json:"type"`
	Description     string      `json:"description"`
	Sleeps          int64       `json:"sleeps"`
	Price           PriseDTO    `json:"prise"`
	Location        LocationDTO `json:"location"`
	Make            string      `json:"make"`
	Model           string      `json:"model"`
	Year            int64       `json:"year"`
	Length          float64     `json:"length"`
	PrimaryImageUrl string      `json:"primary_image_url"`
}

type UserDTO struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type PriseDTO struct {
	Day int64 `json:"day"`
}

type LocationDTO struct {
	City    string  `json:"city"`
	State   string  `json:"state"`
	Zip     string  `json:"zip"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

func ToDTOs(rentals []Rental) []RentalDTO {
	result := make([]RentalDTO, 0, len(rentals))
	for _, model := range rentals {
		result = append(result, ToDTO(model))
	}
	return result
}
func ToDTO(model Rental) RentalDTO {
	return RentalDTO{
		ID: model.ID,
		User: UserDTO{
			ID:        model.User.ID,
			FirstName: model.User.FirstName,
			LastName:  model.User.LastName,
		},
		Name:        model.Name,
		Type:        model.Type,
		Description: model.Description,
		Sleeps:      model.Sleeps,
		Price: PriseDTO{
			Day: model.PricePerDay,
		},
		Location: LocationDTO{
			City:    model.HomeCity,
			State:   model.HomeState,
			Zip:     model.HomeZip,
			Country: model.HomeCountry,
			Lat:     model.Lat,
			Lng:     model.Lng,
		},
		Make:            model.VehicleMake,
		Model:           model.VehicleModel,
		Year:            model.VehicleYear,
		Length:          model.VehicleLength,
		PrimaryImageUrl: model.PrimaryImageUrl,
	}
}
