CREATE TABLE aircraft (
    aircraft_id VARCHAR(36) PRIMARY KEY NOT NULL,
    registration_number INT NOT NULL UNIQUE,
    registration_name VARCHAR(100) NOT NULL
);

CREATE TABLE flight_detail (
    flight_id VARCHAR(36) PRIMARY KEY NOT NULL,
    flight_number INT NOT NULL,
    operating_flight_number INT NOT NULL,
    airline_code VARCHAR(5) NOT NULL,
    operating_airline_code VARCHAR(5)
);

CREATE TABLE segment (
    segment_ref VARCHAR(36) PRIMARY KEY NOT NULL,
    aircraft_id VARCHAR(36) NOT NULL,
    flight_id VARCHAR(36) NOT NULL,
    origin_airport_code VARCHAR(3) NOT NULL,
    destination_airport_code VARCHAR(3) NOT NULL,
    departure_time DATETIME NOT NULL,
    depature_timezone VARCHAR(50) NOT NULL,
    arrival_time DATETIME NOT NULL,
    arrival_timezone VARCHAR(50) NOT NULL,
    subject_to_goverment_approval BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (aircraft_id) REFERENCES aircraft(aircraft_id),
    FOREIGN KEY (flight_id) REFERENCES flight_detail(flight_id)
);

CREATE TABLE segment_stop_over (
    stop_over_id VARCHAR(36) PRIMARY KEY NOT NULL,
    segment_ref VARCHAR(36) NOT NULL,
    airport_code VARCHAR(3) NOT NULL,
    duration INT NOT NULL,
    FOREIGN KEY (segment_ref) REFERENCES segment(segment_ref) ON DELETE CASCADE ON UPDATE CASCADE
);