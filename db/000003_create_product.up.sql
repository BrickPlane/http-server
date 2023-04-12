BEGIN;

CREATE TABLE IF NOT EXIST "product" (
    id SERIAL PRIMARY KEY,
    name TEXT,
    brand TEXT,
    price money
);

COMMIT;