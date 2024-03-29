CREATE TABLE Posts (
    ID CHAR(26) NOT NULL,
    UserID CHAR(26) NOT NULL,
    Picture VARCHAR(255) NOT NULL,
    PlantType VARCHAR(255) NOT NULL,
    PlantAgeInMonth INTEGER NOT NULL,
    PlantName VARCHAR(255) NOT NULL,
    Coordinate POINT NOT NULL,
    Desciprion TEXT NOT NULL,
    CreatedAt DATETIME NOT NULL DEFAULT NOW(),
    PRIMARY KEY (ID),
    FOREIGN KEY FKPostsUsers (UserID) REFERENCES Users (ID) 
    ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE = INNODB DEFAULT CHARSET = UTF8;