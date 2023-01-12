CREATE TABLE IF NOT EXISTS main_data
(
    id  SERIAL UNIQUE  NOT NULL,
    uuid    VARCHAR(250) NOT NULL,
    headline VARCHAR(250) NOT NULL,
    description VARCHAR(250) NOT NULL,
    keywords VARCHAR(500) NOT NULL,
    snippet VARCHAR(250) NOT NULL,
    url VARCHAR(250) NOT NULL,
    CONSTRAINT pk_MainData PRIMARY KEY (id)
    );
