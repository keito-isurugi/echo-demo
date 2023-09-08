TRUNCATE TABLE todos RESTART IDENTITY CASCADE;
INSERT INTO todos (title, content)
VALUES ('テストToDo1', 'これはテストToDo1です！'),
       ('テストToDo2', 'これはテストToDo2です！'),
       ('テストToDo3', 'これはテストToDo3です！'),
       ('テストToDo4', 'これはテストToDo4です！');