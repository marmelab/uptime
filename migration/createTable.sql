CREATE TABLE destination(
	id SERIAL PRIMARY KEY,
	destination VARCHAR(255)
);

CREATE TABLE results(
	id SERIAL PRIMARY KEY,
	destination VARCHAR(255),
	status VARCHAR(30),
	time integer,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

