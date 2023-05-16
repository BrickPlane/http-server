CREATE TABLE IF NOT EXISTS "users" (
	id SERIAL PRIMARY KEY,
    email text UNIQUE,
	login text UNIQUE,
	password text
);

/* migrate -path ./db -database "postgresql://postgres:222@localhost:5432/http-prjct?sslmode=disable" up