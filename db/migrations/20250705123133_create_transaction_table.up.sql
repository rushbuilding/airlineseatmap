CREATE TABLE seat_availability (
    seat_availability_id VARCHAR(36) PRIMARY KEY,
    segment_ref VARCHAR(36) NOT NULL,
    seat_id VARCHAR(36) NOT NULL,
    fee_waived BOOLEAN DEFAULT FALSE,
    fee_waiver_rule_id VARCHAR(36),
    refund_indicator CHAR(1),
    free_of_charge BOOLEAN DEFAULT FALSE,
    is_available BOOLEAN DEFAULT FALSE,

    FOREIGN KEY (segment_ref) REFERENCES segment(segment_ref) ON DELETE CASCADE,
    FOREIGN KEY (seat_id) REFERENCES seat_detail(seat_id) ON DELETE CASCADE,
    UNIQUE (segment_ref, seat_id)
);

CREATE TABLE seat_pricing_option (
    pricing_id VARCHAR(36) PRIMARY KEY,
    seat_availability_id VARCHAR(36) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    price_amount DECIMAL(10, 2) NOT NULL,
    tax_amount DECIMAL(10, 2),

    FOREIGN KEY (seat_availability_id) REFERENCES seat_availability(seat_availability_id)
        ON DELETE CASCADE
);

CREATE TABLE segment_offering (
    segment_offering_id VARCHAR(36) PRIMARY KEY,
    segment_ref VARCHAR(36) NOT NULL,
    service_class ENUM('ECONOMY', 'BUSINESS', 'FIRST_CLASS') NOT NULL,
    fare_basis VARCHAR(10),
    flight_miles INT,
    award_fare BOOLEAN DEFAULT FALSE,
    booking_class CHAR(1),

    FOREIGN KEY (segment_ref) REFERENCES segment(segment_ref) ON DELETE CASCADE
);

CREATE TABLE segment_passenger (
    segment_passenger_id VARCHAR(36) PRIMARY KEY,
    segment_offering_id VARCHAR(36) NOT NULL,
    passenger_id VARCHAR(36) NOT NULL,
    seat_availability_id VARCHAR(36),
    passenger_index INT,
    passenger_name_number VARCHAR(10),
    seat_selection_enabled BOOLEAN DEFAULT TRUE,
    meal_preference VARCHAR(50),
    seat_preference VARCHAR(50),

    FOREIGN KEY (segment_offering_id) REFERENCES segment_offering(segment_offering_id) ON DELETE CASCADE,
    FOREIGN KEY (passenger_id) REFERENCES passengers(passenger_id) ON DELETE CASCADE,
    FOREIGN KEY (seat_availability_id) REFERENCES seat_availability(seat_availability_id) ON DELETE SET NULL
);

CREATE TABLE special_request (
    special_request_id VARCHAR(36) PRIMARY KEY,
    segment_passenger_id VARCHAR(36) NOT NULL,
    request_type ENUM('SPECIAL_SERVICE', 'SPECIAL_REQUEST') NOT NULL,
    description VARCHAR(225),

    FOREIGN KEY (segment_passenger_id) REFERENCES segment_passenger(segment_passenger_id) ON DELETE CASCADE
);