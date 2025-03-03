CREATE TABLE musicians(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    birth_date DATE,
    nationality VARCHAR(50)
);

CREATE TABLE albums(
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    release_date DATE NOT NULL,
    musician_id INT NOT NULL,
    CONSTRAINT fk_album_musician FOREIGN KEY (musician_id) REFERENCES musicians(id) ON DELETE CASCADE,
    UNIQUE (title, musician_id) -- Ensuring a musician doesn't release two albums with the same title
);

CREATE TABLE songs(
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    duration INT NOT NULL CHECK (duration >= 0),  -- Ensuring duration positive values
    album_id INT NOT NULL,
    track_number SMALLINT NOT NULL CHECK (track_number > 0),
    CONSTRAINT fk_song_album FOREIGN KEY (album_id) REFERENCES albums(id) ON DELETE CASCADE,
    UNIQUE (album_id, track_number), -- Ensuring no duplicate track numbers within an album
    UNIQUE (album_id, title) -- Ensuring no duplicate song titles within an album
);

-- Indexes for performance optimization
CREATE INDEX idx_album_musician ON albums(musician_id);
CREATE INDEX idx_song_album ON songs(album_id);