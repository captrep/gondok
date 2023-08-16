CREATE TABLE IF NOT EXISTS links (
    id varchar(255) NOT NULL,
    short_url varchar(50) NOT NULL,
    long_url varchar(255) NOT NULL,
    created_at timestamp,

    PRIMARY KEY(id)
);