ALTER TABLE segment
DROP COLUMN duration;

ALTER TABLE segment
CHANGE COLUMN departure_timezone depature_timezone VARCHAR(50) NOT NULL;