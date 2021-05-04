CREATE TABLE vn_cities (
    id SERIAL PRIMARY KEY NOT NULL,
    city VARCHAR(255) NOT NULL,
    lat REAL NOT NULL,
    lng REAL NOT NULL,
    country VARCHAR(255) NOT NULL,
    iso2 VARCHAR(255) NOT NULL,
    admin_name VARCHAR(255) NOT NULL,
    capital VARCHAR(255) NOT NULL,
    population BIGINT NOT NULL,
    population_proper BIGINT NOT NULL,
    
);