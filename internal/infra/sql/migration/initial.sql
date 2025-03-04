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
    CONSTRAINT fk_album_artist FOREIGN KEY (artist_id) REFERENCES artists(id) ON DELETE CASCADE,
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

-- Data dump
INSERT INTO artists(name, birth_date, nationality) VALUES
    ('Dido', '1971-12-25', 'British'),
    ('Lauryn Hill', '1975-05-26', 'American'),
    ('Daft Punk', '1993-01-01', 'French'),
    ('Linkin Park', '1996-01-01', 'American');

INSERT INTO albums(title, release_date, artist_id) VALUES
    ('No Angel', '2000-06-19', 1),
    ('The Miseducation of Lauryn Hill', '1998-08-25', 2),
    ('Discovery', '2001-03-12', 3),
    ('Random Access Memories', '2013-05-17', 3),
    ('Hybrid Theory', '2000-10-24', 4);

INSERT INTO songs(title, duration, album_id, track_number) VALUES
    ('Here with Me', 255, 1, 1),
    ('Hunter', 271, 1, 2),
    ('Don''t Think of Me', 241, 1, 3),
    ('Thank You', 278, 1, 4),
    ('All You Want', 265, 1, 5),
    ('My Lover''s Gone', 252, 1, 6),
    ('Take My Hand', 366, 1, 7),
    ('I''m No Angel', 231, 1, 8),
    ('My Life', 208, 1, 9),
    ('Isobel', 261, 1, 10),
    ('Honestly OK', 257, 1, 11),
    ('Slide', 263, 1, 12);

INSERT INTO songs(title, duration, album_id, track_number) VALUES
    ('Intro', 47, 2, 1),
    ('Lost Ones', 334, 2, 2),
    ('Ex-Factor', 327, 2, 3),
    ('To Zion (feat. Carlos Santana)', 369, 2, 4),
    ('Doo Wop (That Thing)', 320, 2, 5),
    ('Superstar', 297, 2, 6),
    ('Final Hour', 256, 2, 7),
    ('When It Hurts So Bad', 342, 2, 8),
    ('I Used to Love Him (feat. Mary J. Blige)', 340, 2, 9),
    ('Forgive Them Father', 315, 2, 10),
    ('Every Ghetto, Every City', 315, 2, 11),
    ('Nothing Even Matters', 351, 2, 12),
    ('Everything Is Everything', 293, 2, 13),
    ('The Miseducation of Lauryn Hill', 235, 2, 14),
    ('Can''t Take My Eyes Off of You (I Love You Baby)', 221, 2, 15),
    ('Tell Him', 281, 2, 16);

INSERT INTO songs(title, duration, album_id, track_number) VALUES
    ('One More Time', 320, 3, 1),
    ('Aerodynamic', 227, 3, 2),
    ('Digital Love', 300, 3, 3),
    ('Harder, Better, Faster, Stronger', 224, 3, 4),
    ('Crescendolls', 231, 3, 5),
    ('Nightvision', 106, 3, 6),
    ('Superheroes', 217, 3, 7),
    ('High Life', 202, 3, 8),
    ('Something About Us', 232, 3, 9),
    ('Voyager', 227, 3, 10),
    ('Veridis Quo', 323, 3, 11),
    ('Short Circuit', 202, 3, 12),
    ('Face to Face', 240, 3, 13),
    ('Too Long', 600, 3, 14);

INSERT INTO songs(title, duration, album_id, track_number) VALUES
    ('Give Life Back to Music', 272, 4, 1),
    ('The Game of Love', 323, 4, 2),
    ('Giorgio by Moroder', 569, 4, 3),
    ('Within', 223, 4, 4),
    ('Instant Crush', 337, 4, 5),
    ('Lose Yourself to Dance', 335, 4, 6),
    ('Touch', 512, 4, 7),
    ('Get Lucky', 369, 4, 8),
    ('Beyond', 292, 4, 9),
    ('Motherboard', 345, 4, 10),
    ('Fragments of Time', 288, 4, 11),
    ('Doin'' It Right', 253, 4, 12),
    ('Contact', 379, 4, 13);

INSERT INTO songs(title, duration, album_id, track_number) VALUES
    ('Papercut', 184, 5, 1),
    ('One Step Closer', 156, 5, 2),
    ('With You', 203, 5, 3),
    ('Points of Authority', 200, 5, 4),
    ('Crawling', 209, 5, 5),
    ('Runaway', 184, 5, 6),
    ('By Myself', 190, 5, 7),
    ('In the End', 216, 5, 8),
    ('A Place for My Head', 185, 5, 9),
    ('Forgotten', 194, 5, 10),
    ('Cure for the Itch', 157, 5, 11),
    ('Pushing Me Away', 192, 5, 12);
