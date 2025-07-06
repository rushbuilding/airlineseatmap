ALTER TABLE segment
ADD COLUMN duration INT NOT NULL;

ALTER TABLE segment
CHANGE COLUMN depature_timezone departure_timezone VARCHAR(50) NOT NULL;