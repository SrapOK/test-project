CREATE TABLE IF NOT EXISTS song (
    song_id UUID PRIMARY KEY,
    song VARCHAR(100) NOT NULL,
    author VARCHAR(100) NOT NULL,
    release_date DATE NOT NULL,
    text TEXT NOT NULL,
    link VARCHAR(200) NOT NULL
);