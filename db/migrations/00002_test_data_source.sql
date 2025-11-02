-- +goose Up
-- +goose StatementBegin
-- Пользователи
INSERT INTO "user" (email, password_hash, name, avatar)
VALUES
    ('test1@example.com', 'hashed_password_1', 'Test User 1', 'http://62.109.19.84:9000/mindleak-bucket/defaultAvatar.jpg'),
    ('test2@example.com', 'hashed_password_2', 'Test User 2', 'http://62.109.19.84:9000/mindleak-bucket/defaultAvatar.jpg'),
    ('admin@example.com', 'hashed_password_admin', 'Admin', 'http://62.109.19.84:9000/mindleak-bucket/defaultAvatar.jpg');

-- Профили
INSERT INTO profile (user_id, phone, country, language, sex, age, cover_url)
SELECT user_id, '+1234567890', 'Russia', 'ru', 'male', 25, 'http://62.109.19.84:9000/mindleak-bucket/cover-pic.jpg'
FROM "user" WHERE name = 'Test User 1';

INSERT INTO profile (user_id, phone, country, language, sex, age, cover_url)
SELECT user_id, '+9876543210', 'Germany', 'de', 'female', 30, 'http://62.109.19.84:9000/mindleak-bucket/cover-pic.jpg'
FROM "user" WHERE name = 'Test User 2';

-- Статьи
INSERT INTO article (author_id, title, content, image, status)
SELECT user_id, 'Welcome to MindLeak!', 'This is the first published article from Test User 1.',
       'http://62.109.19.84:9000/mindleak-bucket/articles/test_image1.jpg', 'published'
FROM "user" WHERE name = 'Test User 1';

INSERT INTO article (author_id, title, content, status)
SELECT user_id, 'Draft article example', 'This article is in draft mode.', 'draft'
FROM "user" WHERE name = 'Test User 2';

-- Комментарий
INSERT INTO comment (article_id, user_id, content)
SELECT a.article_id, u.user_id, 'Nice article! Keep up the good work.'
FROM article a
         JOIN "user" u ON u.name = 'Test User 2'
WHERE a.status = 'published'
    LIMIT 1;

-- Лайки
INSERT INTO article_like (user_id, article_id)
SELECT u.user_id, a.article_id
FROM "user" u
         JOIN article a ON a.status = 'published'
WHERE u.name IN ('Test User 2', 'Admin');

-- Медиафайлы
INSERT INTO media (article_id, uploader_id, type, mime, url, size_bytes, description)
SELECT a.article_id, u.user_id, 'image', 'image/jpeg',
       'http://62.109.19.84:9000/mindleak-bucket/articles/sample_attachment.jpg',
       256000, 'Attached image for demo'
FROM article a
         JOIN "user" u ON a.author_id = u.user_id
WHERE a.status = 'published'
    LIMIT 1;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DELETE FROM media;
DELETE FROM article_like;
DELETE FROM comment;
DELETE FROM article;
DELETE FROM profile;
DELETE FROM "user";
-- +goose StatementEnd
