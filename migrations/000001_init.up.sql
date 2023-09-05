CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name varchar(255),
    segments integer ARRAY
);

CREATE TABLE segments
(
    id SERIAL PRIMARY KEY,
    slug varchar(255)
);
