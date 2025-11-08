INSERT INTO user (user_id, login, password_hash, email, created_at, updated_at)
VALUES
    (gen_random_uuid(), 'admin', 'hashed_password_1', 'admin@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'user1', 'hashed_password_2', 'user1@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO user_profile (user_id, display_name, bio, avatar_url, created_at, updated_at)
SELECT user_id, 'Admin User', 'Administrator account', 'https://example.com/avatars/admin.jpg', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM user WHERE login = 'admin'
UNION
SELECT user_id, 'User One', 'Regular user account', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM user WHERE login = 'user1';

INSERT INTO category (category_id, name, description, created_at, updated_at)
VALUES
    (gen_random_uuid(), 'Technology', 'Articles about tech innovations', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Business', 'Business and startup insights', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO tag (tag_id, name, created_at)
VALUES
    (gen_random_uuid(), 'AI', CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Startups', CURRENT_TIMESTAMP);

INSERT INTO article (article_id, title, content, author_id, published_at, status, created_at, updated_at)
SELECT gen_random_uuid(), 'First Tech Article', 'Content about AI advancements...', user_id, CURRENT_TIMESTAMP, 'published', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM user WHERE login = 'admin';

INSERT INTO article_category (article_id, category_id, created_at)
SELECT article_id, category_id, CURRENT_TIMESTAMP
FROM article, category
WHERE article.title = 'First Tech Article' AND category.name = 'Technology';

INSERT INTO article_tag (article_id, tag_id, created_at)
SELECT article_id, tag_id, CURRENT_TIMESTAMP
FROM article, tag
WHERE article.title = 'First Tech Article' AND tag.name = 'AI';

INSERT INTO comment (comment_id, article_id, user_id, content, created_at, updated_at)
SELECT gen_random_uuid(), article_id, user_id, 'Great article!', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM article, user
WHERE article.title = 'First Tech Article' AND user.login = 'user1';

INSERT INTO notification (notification_id, user_id, type, content, is_read, created_at)
SELECT gen_random_uuid(), user_id, 'comment', 'User1 commented on your article', FALSE, CURRENT_TIMESTAMP
FROM user WHERE login = 'admin';