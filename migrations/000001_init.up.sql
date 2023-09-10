CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name varchar(255)
);

CREATE TABLE segments
(
    id SERIAL PRIMARY KEY,
    slug varchar(255) NOT NULL
);

CREATE TABLE users_segments
(
    user_id INTEGER REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
    segment_id INTEGER REFERENCES segments (id) ON UPDATE CASCADE,
    CONSTRAINT users_segments_pk PRIMARY KEY (user_id, segment_id)
);


