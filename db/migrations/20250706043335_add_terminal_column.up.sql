ALTER TABLE segment
ADD COLUMN departure_terminal VARCHAR(10) NOT NULL;

ALTER TABLE segment
ADD COLUMN arrival_terminal VARCHAR(10) NOT NULL;