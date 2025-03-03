CREATE TABLE artists(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    birth_date DATE,
    nationality VARCHAR(50)
);

CREATE TABLE albums(
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    release_date DATE NOT NULL,
    artist_id INT NOT NULL,
    CONSTRAINT fk_album_artist FOREIGN KEY (artist_id) REFERENCES artist(id) ON DELETE CASCADE,
    UNIQUE (title, artist_id) -- Ensuring a artist doesn't release two albums with the same title
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
CREATE INDEX idx_album_artist ON albums(artist_id);
CREATE INDEX idx_song_album ON songs(album_id);