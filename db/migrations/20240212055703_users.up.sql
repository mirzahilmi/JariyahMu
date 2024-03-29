CREATE TABLE Users (
    ID CHAR(26) NOT NULL,
    FullName VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL UNIQUE,
    HashedPassword CHAR(60) NOT NULL,
    ProfilePicture VARCHAR(255) NOT NULL DEFAULT "",
    Active BOOLEAN NOT NULL DEFAULT FALSE,
    CreatedAt DATETIME NOT NULL DEFAULT NOW(),
    UpdatedAt DATETIME NOT NULL ON UPDATE NOW(),
    PRIMARY KEY (ID)
) ENGINE = INNODB DEFAULT CHARSET = UTF8;