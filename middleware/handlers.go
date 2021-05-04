package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http"              // used to access the request and response object of the api
	"os"                    // used to read the environment variable
	"simple-restapi/models" // models package where User schema is defined

	// package used to covert string into int type
	"github.com/gorilla/mux"   // used to get the params from the route
	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

func GetCityInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	cityname := params["cityname"]

	city, err := getCity(cityname)

	if err != nil {
		log.Fatalf("Unable to get city. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(city)
}

func GetDistance(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	var cityInput models.CityInput

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&cityInput)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	cityInfo1, err := getCity(cityInput.CityName1)
	cityInfo2, err := getCity(cityInput.CityName2)

	// format a response object
	res := response{
		ID:      123,
		Message: *cityInfo1[0].AdminName + " " + *cityInfo2[0].AdminName,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// get one city from the DB by its cityname
func getCity(cityName string) ([]models.City, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var cities []models.City

	// create the select sql query
	sqlStatement := `SELECT * FROM vn_cities WHERE city ILIKE $1`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, "%"+cityName+"%")

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var city models.City

		// unmarshal the row object to user
		err = rows.Scan(&city.ID, &city.City, &city.Lat, &city.Lng, &city.Country, &city.Iso2, &city.AdminName, &city.Capital, &city.Population, &city.PopulationProper)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		cities = append(cities, city)

	}

	// return empty user on error
	return cities, err
}
