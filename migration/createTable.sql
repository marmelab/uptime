CREATE TABLE destination(
	id SERIAL PRIMARY KEY,
	destination VARCHAR(255)
);
CREATE TABLE results(
	destination VARCHAR(255),
	status VARCHAR(30),
	time smallint
);

