CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(64) NOT NULL CHECK (role IN ('client', 'moderator'))
);


CREATE TABLE pvz (
    id UUID PRIMARY KEY,
    registration_date TIMESTAMP NOT NULL DEFAULT NOW(),
    city VARCHAR(64) NOT NULL CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань'))
);

CREATE TABLE receptions (
    id UUID PRIMARY KEY,
    date_time TIMESTAMP NOT NULL DEFAULT NOW(),
    pvz_id UUID REFERENCES pvz(id) ON DELETE CASCADE,
    status VARCHAR(64) NOT NULL CHECK (status IN ('in_progress', 'close'))
);

CREATE TABLE products (
    id UUID PRIMARY KEY,
    date_time TIMESTAMP NOT NULL DEFAULT NOW(),
    type VARCHAR(64) NOT NULL CHECK (type IN ('электроника', 'одежда', 'обувь')),
    reception_id UUID REFERENCES receptions(id) ON DELETE CASCADE
);

ALTER TABLE receptions ADD COLUMN product_id UUID[];


