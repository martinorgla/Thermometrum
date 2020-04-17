USE thermometrum;

CREATE TABLE temperatures (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    room VARCHAR(30) NOT NULL,
    temperature float,
    humidity float,
    time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)