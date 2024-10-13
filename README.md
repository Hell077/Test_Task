## Перед запуском программы убедитесь что у вас есть пользователь postgres с паролем postgres 

## Также создайте базу данных 

````
musiclibrary
````

### [Ссылка на swagger](http://localhost:8000/swagger/index.html#) 


Запуск 
````bash
go run main.go
````


## Тестовые данные 

````sql
-- Заполнение таблицы songs тестовыми данными
INSERT INTO songs (group_name, song_name, release_date, text) VALUES
    ('Queen', 'Bohemian Rhapsody', '1975-10-31', 'Is this the real life? Is this just fantasy?'),
    ('The Beatles', 'Hey Jude', '1968-08-26', 'Hey Jude, dont make it bad.'),
    ('Adele', 'Someone Like You', '2011-01-24', 'Never mind, Ill find someone like you.'),
    ('Nirvana', 'Smells Like Teen Spirit', '1991-09-10', 'With the lights out, it’s less dangerous.'),
    ('Linkin Park', 'In the End', '2000-10-12', 'I tried so hard and got so far.');

-- Заполнение таблицы song_lyrics тестовыми данными
INSERT INTO song_lyrics (song_id, verse_number, text) VALUES
    (1, 1, 'Is this the real life? Is this just fantasy?'),
    (1, 2, 'Caught in a landslide, no escape from reality.'),
    (2, 1, 'Hey Jude, dont make it bad.'),
    (2, 2, 'Take a sad song and make it better.'),
    (3, 1, 'Never mind, Ill find someone like you.'),
    (3, 2, 'I wish nothing but the best for you, too.'),
    (4, 1, 'With the lights out, it’s less dangerous.'),
    (4, 2, 'Here we are now, entertain us.'),
    (5, 1, 'I tried so hard and got so far.'),
    (5, 2, 'But in the end, it doesnt even matter.');
````