CREATE TABLE destination(
	id SERIAL PRIMARY KEY,
	destination VARCHAR(255)
);

CREATE TABLE results(
	id SERIAL PRIMARY KEY,
	target_id smallint REFERENCES destination (id),
	destination VARCHAR(255),
	status VARCHAR(30),
	duration integer,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
CREATE TABLE testDestination(
	id SERIAL PRIMARY KEY,
	destination VARCHAR(255)
);

CREATE TABLE testResults(
	id SERIAL PRIMARY KEY,
	target_id smallint REFERENCES destination (id),
	destination VARCHAR(255),
	status VARCHAR(30),
	duration integer,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

