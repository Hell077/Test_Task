DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'song_lyrics') THEN
        DROP TABLE song_lyrics;
    END IF;


    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'songs') THEN
        DROP TABLE songs;
    END IF;
END $$;
