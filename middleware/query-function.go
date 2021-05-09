package middleware

import (

	// package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"math" // used to access the request and response object of the api

	// used to read the environment variable
	"vncities-distance/models" // models package where User schema is defined

	// package used to covert string into int type
	// used to get the params from the route
	// package used to read the .env file
	_ "github.com/lib/pq" // postgres golang driver
)

// get city from the DB by its cityname
func getCity(cityName string) []models.City {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var cities []models.City

	// create the select sql query
	sqlStatement := `SELECT * FROM vn_cities 
	WHERE city ILIKE $1 
	OR city ILIKE unaccent($1) 
	OR unaccent(city) ILIKE $1` //unaccent: drop vietnamese characters

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

func insertDistanceHistory(history models.StoredDistance) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO search_distance_history (description, distance, unit_measurement) VALUES ($1, $2, $3) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, history.Description, history.Distance, history.Unit).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func getDistanceHistory() []models.StoredDistance {
	db := createConnection()

	// close the db connection
	defer db.Close()
	// create a user of models.User type
	var history []models.StoredDistance

	// create the select sql query
	sqlStatement := `SELECT * FROM search_distance_history`
	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to scan the row. %v", err)
	} else {
		// close the statement
		defer rows.Close()

		// iterate over the rows
		for rows.Next() {
			var distanceHistory models.StoredDistance

			// unmarshal the row object to user
			err = rows.Scan(&distanceHistory.ID, &distanceHistory.Description, &distanceHistory.Distance, &distanceHistory.Unit, &distanceHistory.DateCreated)

			if err != nil {
				log.Fatalf("Unable to scan the row. %v", err)
			}

			// append the user in the users slice
			history = append(history, distanceHistory)

		}
	}
	return history
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
