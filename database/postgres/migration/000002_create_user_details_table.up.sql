CREATE TABLE IF NOT EXISTS user_details (
    user_id INTEGER  NOT NULL,
    nrp varchar(9) NOT NULL,
    prodi varchar(255) NOT NULL,
    program varchar(255) NOT NULL,
    company varchar(255) NULL,
    batch smallint NOT NULL
);