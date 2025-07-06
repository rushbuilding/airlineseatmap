CREATE TABLE frequent_flyer (
    frequent_flyer_id     VARCHAR(36) PRIMARY KEY,
    passenger_id          VARCHAR(36) NOT NULL,
    airline_code          VARCHAR(10) NOT NULL,
    membership_number     VARCHAR(50) NOT NULL,
    tier_number           INT,

    FOREIGN KEY (passenger_id) REFERENCES passengers(passenger_id)
        ON DELETE CASCADE
);