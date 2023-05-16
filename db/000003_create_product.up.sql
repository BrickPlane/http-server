BEGIN;

CREATE TABLE IF NOT EXISTS "product" (
    id SERIAL PRIMARY KEY,
    name TEXT,
    brand TEXT,
    price DOUBLE PRECISION 
);

COMMIT;