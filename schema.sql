CREATE TABLE tasks (
	id BIGSERIAL PRIMARY KEY,
	name text NOT NULL,
	info text,
	isDone boolean NOT NULL
);
