
CREATE TABLE passengers (
    passenger_id VARCHAR(36) PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    birth_of_date DATE,
    gender ENUM('Male', 'Female', 'Other'),
    passenger_type ENUM('ADT', 'CHD', 'INF')
);

CREATE TABLE passenger_document (
    document_id VARCHAR(36) PRIMARY KEY,
    passenger_id VARCHAR(36) NOT NULL,
    document_type VARCHAR(50),
    issuing_country CHAR(3),
    country_of_birth CHAR(3),
    nationality CHAR(3),
    FOREIGN KEY (passenger_id) REFERENCES passengers(passenger_id)
);

CREATE TABLE passenger_email (
    email_id VARCHAR(36) PRIMARY KEY,
    passenger_id VARCHAR(36) NOT NULL,
    email_address VARCHAR(255) UNIQUE NOT NULL,
    FOREIGN KEY (passenger_id) REFERENCES passengers(passenger_id)
);

CREATE TABLE passenger_phone (
    phone_id VARCHAR(36) PRIMARY KEY,
    passenger_id VARCHAR(36) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    FOREIGN KEY (passenger_id) REFERENCES passengers(passenger_id)
);

CREATE TABLE passenger_address (
    address_id VARCHAR(36) PRIMARY KEY,
    passenger_id VARCHAR(36) NOT NULL,
    first_street_line VARCHAR(100),
    second_street_line VARCHAR(100),
    postcode VARCHAR(10),
    state VARCHAR(50),
    country CHAR(3),
    address_type ENUM('HOME', 'OFFICE'),
    FOREIGN KEY (passenger_id) REFERENCES passengers(passenger_id)
);