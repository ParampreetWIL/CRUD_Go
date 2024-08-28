CREATE TABLE tasks (
	id BIGSERIAL PRIMARY KEY,
	name text NOT NULL,
	info text,
	isDone boolean NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email_token TEXT NOT NULL,
    jwt_token TEXT NOT NULL,
    name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
