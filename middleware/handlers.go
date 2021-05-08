package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"math"
	"net/http"                 // used to access the request and response object of the api
	"os"                       // used to read the environment variable
	"vncities-distance/models" // models package where User schema is defined

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
type ErrorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
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
	cities := getCity(cityname)

	var responseData interface{}
	if len(cities) > 0 {
		responseData = cities
	} else {
		responseData = ErrorResponse{
			Code:    500,
			Message: "Can not find the city",
		}
	}

	// send the response
	json.NewEncoder(w).Encode(responseData)
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
	var distance float64
	var responseData interface{}

	// decode the json request to city
	err := json.NewDecoder(r.Body).Decode(&cityInput)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	cityInfo1 := getCity(cityInput.CityName1)
	cityInfo2 := getCity(cityInput.CityName2)

	//calculate distance
	if len(cityInfo1) > 0 && len(cityInfo2) > 0 {
		distance = Haversine(cityInfo1[0].Lng, cityInfo1[0].Lat, cityInfo2[0].Lng, cityInfo2[0].Lat)
		responseData = models.DistanceCities{
			Code:        200,
			Description: "Distance from " + cityInfo1[0].City + "," + *cityInfo1[0].AdminName + " to " + cityInfo2[0].City + "," + *cityInfo2[0].AdminName,
			Distance:    distance,
			Unit:        "kilometer",
		}
	} else {
		responseData = ErrorResponse{
			Code:    500,
			Message: "Can not find the distance because there is an invalid city",
		}
	}
	json.NewEncoder(w).Encode(responseData)
}

// get one city from the DB by its cityname
func getCity(cityName string) []models.City {
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
		log.Fatalf("Unable to scan the row. %v", err)
	} else {
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
	}
	return cities
}

func Haversine(lonFrom float64, latFrom float64, lonTo float64, latTo float64) (distance float64) {
	const earthRadius = float64(6371)
	var deltaLat = (latTo - latFrom) * (math.Pi / 180)
	var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)

	var a = math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(latFrom*(math.Pi/180))*math.Cos(latTo*(math.Pi/180))*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance = earthRadius * c //in meters

	return distance
}
