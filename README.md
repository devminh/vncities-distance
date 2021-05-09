# vncities-distance
A simple restapi to calculate distance between 2 cities in VN.

This project uses a remote PostgreSQL DB on https://www.elephantsql.com/

VN cities csv: https://simplemaps.com/data/vn-cities

go run main.go to run

Test API:

curl --location --request GET 'http://localhost:10000/api/city?cityname=ho%20chi%20minh'


curl --location --request POST 'http://localhost:10000/api/get_distance' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "city_name_1":"ho chi minh",
    "city_name_2":"vung tau"
}'



curl --location --request GET 'http://localhost:10000/api/get_distance_history'

