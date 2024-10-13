DO $$
BEGIN
    -- Создаем таблицу songs, если она не существует
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'songs') THEN
        CREATE TABLE songs (
            id SERIAL PRIMARY KEY,
            group_name VARCHAR(255) NOT NULL,
            song_name VARCHAR(255) NOT NULL,
            release_date DATE,
            text TEXT
        );

        CREATE INDEX idx_song_name ON songs(song_name);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'song_lyrics') THEN
        CREATE TABLE song_lyrics (
            id SERIAL PRIMARY KEY,
            song_id INTEGER NOT NULL,
            verse_number INTEGER NOT NULL,
            text TEXT
        );

        CREATE INDEX idx_song_lyrics_song_id ON song_lyrics(song_id);
        CREATE INDEX idx_song_lyrics_verse_number ON song_lyrics(verse_number);
    END IF;

    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'song_lyrics') THEN
        ALTER TABLE song_lyrics
        ADD CONSTRAINT fk_song_id FOREIGN KEY (song_id) REFERENCES songs(id) ON DELETE CASCADE;
    END IF;
END $$;
