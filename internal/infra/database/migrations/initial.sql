-- Inserting Musician Data
INSERT INTO Musician (name, birth_date, nationality) VALUES 
    ('Dido', '1971-12-25', 'British'),
    ('Lauryn Hill', '1975-05-26', 'American'),
    ('Lenny Kravitz', '1964-05-26', 'American');

-- Inserting Album Data
INSERT INTO Album (title, release_date, musician_id) VALUES 
    ('No Angel', '2000-06-19', 1),
    ('The Miseducation of Lauryn Hill', '1998-08-25', 2)
    ('Are You Gonna Go My Way', '1993-03-03', 3);

INSERT INTO Song (title, duration, album_id, track_number) VALUES 
    ('Here with Me', 255, 1, 1),
    ('Hunter', 271, 1, 2),
    ('Don''t Think of Me', 241, 1, 3),
    ('Thank You', 278, 1, 4),
    ('All You Want', 265, 1, 5);

INSERT INTO Song (title, duration, album_id, track_number) VALUES 
    ('Doo Wop (That Thing)', 280, 2, 1),
    ('Ex-Factor', 348, 2, 2),
    ('Everything Is Everything', 264, 2, 3),
    ('To Zion', 388, 2, 4),
    ('Lost Ones', 296, 2, 5);

-- Inserting Song Data for Lenny Kravitz - Are You Gonna Go My Way (1993)
INSERT INTO Song (title, duration, album_id, track_number) VALUES 
    ('Are You Gonna Go My Way', 231, 3, 1),
    ('Mr. Cab Driver', 271, 3, 2),
    ('Believe', 267, 3, 3),
    ('I Build This Garden for Us', 266, 3, 4),
    ('Sister', 299, 3, 5);
