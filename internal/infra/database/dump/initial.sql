-- Sample Data
INSERT INTO Musician (name, birth_date, nationality) VALUES 
    ('John Mayer', '1977-10-16', 'American'),
    ('Adele', '1988-05-05', 'British');

INSERT INTO Album (title, release_date, musician_id) VALUES 
    ('Continuum', '2006-09-12', 1),
    ('Sob Rock', '2021-07-16', 1),
    ('25', '2015-11-20', 2),
    ('30', '2021-11-19', 2);

INSERT INTO Song (title, duration, album_id, track_number) VALUES 
    ('Gravity', '00:04:06', 1, 1),
    ('Slow Dancing in a Burning Room', '00:04:02', 1, 2),
    ('New Light', '00:03:37', 2, 1),
    ('Easy On Me', '00:03:45', 3, 1),
    ('Hello', '00:04:55', 3, 2),
    ('I Drink Wine', '00:06:16', 4, 1);
