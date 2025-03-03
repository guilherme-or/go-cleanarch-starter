INSERT INTO musicians(name, birth_date, nationality) VALUES 
    ('Dido', '1971-12-25', 'British'),
    ('Lauryn Hill', '1975-05-26', 'American'),
    ('Daft Punk', '1993-01-01', 'French');

INSERT INTO albums(title, release_date, musician_id) VALUES 
    ('No Angel', '2000-06-19', 1),
    ('The Miseducation of Lauryn Hill', '1998-08-25', 2),
    ('Discovery', '2001-03-12', 3),
    ('Random Access Memories', '2013-05-17', 3);

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
    ('Doo Wop (That Thing)', 280, 2, 1),
    ('Ex-Factor', 348, 2, 2),
    ('Everything Is Everything', 264, 2, 3),
    ('To Zion', 388, 2, 4),
    ('Lost Ones', 296, 2, 5),
    ('Superstar', 294, 2, 6),
    ('Final Hour', 240, 2, 7),
    ('When It Hurts So Bad', 300, 2, 8),
    ('I Used to Love Him', 324, 2, 9),
    ('Forgive Them Father', 330, 2, 10),
    ('Every Ghetto, Every City', 309, 2, 11),
    ('Nothing Even Matters', 337, 2, 12),
    ('Everything Is Everything', 263, 2, 13),
    ('The Miseducation of Lauryn Hill', 278, 2, 14);

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
