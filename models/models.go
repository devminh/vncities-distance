package models

// User schema of the user table
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Age      int64  `json:"age"`
}

type City struct {
	ID               int64   `json:"id"`
	City             string  `json:"city"`
	Lat              float64 `json:"lat"`
	Lng              float64 `json:"lng"`
	Country          string  `json:"country"`
	Iso2             string  `json:"iso2"`
	AdminName        *string `json:"admin_name"`
	Capital          *string `json:"capital"`
	Population       *int64  `json:"population"`
	PopulationProper *int64  `json:"population_proper"`
}

type DistanceCities struct {
	Code        int32   `json:"code"`
	Description string  `json:"description"`
	Distance    float64 `json:"distance"`
	Unit        string  `json:"unit_measure"`
}

type StoredDistance struct {
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Distance    float64 `json:"distance"`
	Unit        string  `json:"unit_measurement"`
	DateCreated string  `json:"date_created"`
}

type CityInput struct {
	CityName1 string `json:"city_name_1"`
	CityName2 string `json:"city_name_2"`
}
