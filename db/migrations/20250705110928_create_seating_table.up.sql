CREATE TABLE cabin (
    cabin_id VARCHAR(36) PRIMARY KEY NOT NULL,
    aircraft_id VARCHAR(36) NOT NULL,
    service_class ENUM('ECONOMY', 'BUSINESS', 'FIRST_CLASS') NOT NULL,
    deck ENUM('MAIN', 'UPPER') DEFAULT 'MAIN',
    first_row_number INT NOT NULL,
    last_row_number INT NOT NULL,

    FOREIGN KEY (aircraft_id) REFERENCES aircraft(aircraft_id)
        ON DELETE RESTRICT
);

CREATE TABLE cabin_configuration (
    cabin_layout_id VARCHAR(36) PRIMARY KEY NOT NULL,
    cabin_id VARCHAR(36) NOT NULL,
    column_code VARCHAR(20) NOT NULL,
    column_number INT NOT NULL,

    FOREIGN KEY (cabin_id) REFERENCES cabin(cabin_id)
        ON DELETE CASCADE,

    UNIQUE (cabin_id, column_code)
);

CREATE TABLE cabin_row (
    row_id VARCHAR(36) PRIMARY KEY NOT NULL,
    cabin_id VARCHAR(36) NOT NULL,
    row_number INT NOT NULL,
    is_enable BOOLEAN DEFAULT TRUE,
    disable_cause VARCHAR(255),
    row_type ENUM('EXIT', 'WING_BEGIN', 'WING_END') DEFAULT NULL,
    row_designation ENUM('FRONT_OF_CABIN', 'VACANT_OR_OFFERED_LAST') DEFAULT NULL,

    FOREIGN KEY (cabin_id) REFERENCES cabin(cabin_id)
        ON DELETE CASCADE,

    UNIQUE (cabin_id, row_number)
);

CREATE TABLE seat_detail (
    seat_id VARCHAR(36) PRIMARY KEY NOT NULL,
    cabin_id VARCHAR(36) NOT NULL,
    row_id VARCHAR(36) NOT NULL,
    cabin_layout_id VARCHAR(36) NOT NULL,
    store_front_slot_code VARCHAR(10) NOT NULL,
    is_selectable BOOLEAN DEFAULT TRUE,

    FOREIGN KEY (cabin_id) REFERENCES cabin(cabin_id)
        ON DELETE CASCADE,
    FOREIGN KEY (row_id) REFERENCES cabin_row(row_id)
        ON DELETE CASCADE,
    FOREIGN KEY (cabin_layout_id) REFERENCES cabin_configuration(cabin_layout_id)
        ON DELETE CASCADE
);


CREATE TABLE seat_characteristic (
    characteristic_id VARCHAR(36) PRIMARY KEY,
    limitation_code VARCHAR(30)
);

CREATE TABLE seat_characteristic_map (
    seat_id VARCHAR(36) NOT NULL,
    characteristic_id VARCHAR(36) NOT NULL,
    is_display BOOLEAN DEFAULT FALSE,

    PRIMARY KEY (seat_id, characteristic_id),
    FOREIGN KEY (seat_id) REFERENCES seat_detail(seat_id)
        ON DELETE CASCADE,
    FOREIGN KEY (characteristic_id) REFERENCES seat_characteristic(characteristic_id)
        ON DELETE CASCADE
);