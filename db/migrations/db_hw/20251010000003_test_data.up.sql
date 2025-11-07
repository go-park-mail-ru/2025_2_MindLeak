INSERT INTO "user" (user_id, login, password, email, name, avatar, created_at, updated_at)
VALUES
    (gen_random_uuid(), 'admin', 'hashed_password_1', 'admin@example.com', 'Admin', 'https://example.com/avatars/admin.jpg', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'user1', 'hashed_password_2', 'user1@example.com', 'User One', 'https://example.com/avatars/user1.jpg', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'user2', 'hashed_password_3', 'user2@example.com', 'User Two', 'https://example.com/avatars/user2.jpg', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'user3', 'hashed_password_4', 'user3@example.com', 'User Three', 'https://example.com/avatars/user3.jpg', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'user4', 'hashed_password_5', 'user4@example.com', 'User Four', 'https://example.com/avatars/user4.jpg', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);



INSERT INTO profile (profile_id, user_id, phone, country, language, sex, date_of_birth, age, description, cover_url, created_at, updated_at)
SELECT gen_random_uuid(), user_id, '+1234567890', 'USA', 'English', 'undefined', '1990-01-01', 30, 'Admin profile description', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'admin'
UNION
SELECT gen_random_uuid(), user_id, '+1234567891', 'USA', 'English', 'male', '1992-02-14', 28, 'Regular user description', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user1'
UNION
SELECT gen_random_uuid(), user_id, '+1234567892', 'USA', 'English', 'female', '1994-03-22', 26, 'Tech enthusiast', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user2'
UNION
SELECT gen_random_uuid(), user_id, '+1234567893', 'UK', 'English', 'male', '1985-07-19', 35, 'Business and marketing expert', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user3'
UNION
SELECT gen_random_uuid(), user_id, '+1234567894', 'Canada', 'English', 'female', '1989-05-09', 31, 'Designer and artist', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user4';



INSERT INTO category (category_id, name, description, created_at, updated_at)
VALUES
    (gen_random_uuid(), 'Technology', 'Articles about tech innovations', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Business', 'Business and startup insights', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Health', 'Articles on health and wellness', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Lifestyle', 'Tips and trends in lifestyle', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Education', 'Insights about education and learning', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);



INSERT INTO tag (tag_id, name, created_at)
VALUES
    (gen_random_uuid(), 'AI', CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Startups', CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'HealthTech', CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Design', CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Marketing', CURRENT_TIMESTAMP);



INSERT INTO article (article_id, title, content, author_id, media_url, topic_id, status, comments_count, reposts_count, views_count, created_at, updated_at)
SELECT gen_random_uuid(), 'Tech Trends 2025', 'Content about upcoming technology trends...', user_id, NULL, 1, 'published', 10, 3, 100, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'admin'
UNION
SELECT gen_random_uuid(), 'Business Strategies', 'How to scale your business...', user_id, NULL, 2, 'published', 5, 1, 50, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user1'
UNION
SELECT gen_random_uuid(), 'AI Revolution', 'Artificial intelligence advancements...', user_id, NULL, 1, 'published', 8, 2, 75, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user2'
UNION
SELECT gen_random_uuid(), 'Marketing 101', 'Understanding digital marketing...', user_id, NULL, 2, 'published', 12, 4, 150, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user3'
UNION
SELECT gen_random_uuid(), 'The Future of Design', 'Innovations in graphic design...', user_id, NULL, 3, 'published', 15, 6, 200, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user4';



INSERT INTO article_category (article_id, category_id, created_at)
SELECT article_id, category_id, CURRENT_TIMESTAMP
FROM article, category
WHERE article.title = 'Tech Trends 2025' AND category.name = 'Technology'
UNION
SELECT article_id, category_id, CURRENT_TIMESTAMP
FROM article, category
WHERE article.title = 'Business Strategies' AND category.name = 'Business'
UNION
SELECT article_id, category_id, CURRENT_TIMESTAMP
FROM article, category
WHERE article.title = 'AI Revolution' AND category.name = 'Technology'
UNION
SELECT article_id, category_id, CURRENT_TIMESTAMP
FROM article, category
WHERE article.title = 'Marketing 101' AND category.name = 'Business'
UNION
SELECT article_id, category_id, CURRENT_TIMESTAMP
FROM article, category
WHERE article.title = 'The Future of Design' AND category.name = 'Lifestyle';



INSERT INTO article_tag (article_id, tag_id, created_at)
SELECT article_id, tag_id, CURRENT_TIMESTAMP
FROM article, tag
WHERE article.title = 'Tech Trends 2025' AND tag.name = 'AI'
UNION
SELECT article_id, tag_id, CURRENT_TIMESTAMP
FROM article, tag
WHERE article.title = 'Business Strategies' AND tag.name = 'Startups'
UNION
SELECT article_id, tag_id, CURRENT_TIMESTAMP
FROM article, tag
WHERE article.title = 'AI Revolution' AND tag.name = 'AI'
UNION
SELECT article_id, tag_id, CURRENT_TIMESTAMP
FROM article, tag
WHERE article.title = 'Marketing 101' AND tag.name = 'Marketing'
UNION
SELECT article_id, tag_id, CURRENT_TIMESTAMP
FROM article, tag
WHERE article.title = 'The Future of Design' AND tag.name = 'Design';



INSERT INTO comment (comment_id, article_id, user_id, content, created_at, updated_at)
SELECT gen_random_uuid(), article_id, user_id, 'Great article!', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM article, "user"
WHERE article.title = 'Tech Trends 2025' AND "user".login = 'user1'
UNION
SELECT gen_random_uuid(), article_id, user_id, 'Very informative, thanks!', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM article, "user"
WHERE article.title = 'Business Strategies' AND "user".login = 'user2'
UNION
SELECT gen_random_uuid(), article_id, user_id, 'AI is changing everything!', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM article, "user"
WHERE article.title = 'AI Revolution' AND "user".login = 'user3'
UNION
SELECT gen_random_uuid(), article_id, user_id, 'Great marketing tips!', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM article, "user"
WHERE article.title = 'Marketing 101' AND "user".login = 'user4'
UNION
SELECT gen_random_uuid(), article_id, user_id, 'Design tips are helpful!', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM article, "user"
WHERE article.title = 'The Future of Design' AND "user".login = 'user1';



INSERT INTO notification (notification_id, user_id, type, content, is_read, created_at)
SELECT gen_random_uuid(), user_id, 'comment', 'User1 commented on your article', FALSE, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'admin'
UNION
SELECT gen_random_uuid(), user_id, 'comment', 'User2 commented on your article', FALSE, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user1'
UNION
SELECT gen_random_uuid(), user_id, 'comment', 'User3 commented on your article', FALSE, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user2'
UNION
SELECT gen_random_uuid(), user_id, 'comment', 'User4 commented on your article', FALSE, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user3'
UNION
SELECT gen_random_uuid(), user_id, 'comment', 'User1 commented on your article', FALSE, CURRENT_TIMESTAMP
FROM "user" WHERE login = 'user4';
