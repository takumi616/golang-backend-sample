CREATE TABLE IF NOT EXISTS vocabularies (
    vocabulary_no SERIAL PRIMARY KEY,
    title VARCHAR(20) NOT NULL,
    meaning TEXT NOT NULL,
    sentence TEXT NOT NULL
);