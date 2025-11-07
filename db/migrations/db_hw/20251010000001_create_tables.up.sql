CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE article_status AS ENUM ('draft', 'published', 'archived');
CREATE TYPE notification_type AS ENUM ('like', 'comment', 'follow');
CREATE TYPE media_type AS ENUM ('image', 'video', 'audio', 'document', 'other');

CREATE TABLE "user" (
                        user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        login TEXT NOT NULL UNIQUE CHECK (LENGTH(login) >= 4 AND LENGTH(login) <= 32 AND login !~ ' '),
                        password TEXT NOT NULL CHECK (LENGTH(password) <= 255),
                        email TEXT NOT NULL UNIQUE CHECK (LENGTH(email) <= 320 AND email ~ '^[^\s@]+@[^\s@]+\.[^\s@]+$'),
                        name TEXT NOT NULL,
                        avatar TEXT,
                        created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        CONSTRAINT chk_email_format CHECK (email ~ '^[^\s@]+@[^\s@]+\.[^\s@]+$'),
                        CONSTRAINT chk_login_format CHECK (LENGTH(login) >= 4 AND LENGTH(login) <= 32 AND login !~ ' ')
);

CREATE TABLE profile (
                         profile_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         user_id UUID NOT NULL UNIQUE REFERENCES "user"(user_id) ON DELETE CASCADE,
                         phone TEXT,
                         country TEXT,
                         language TEXT,
                         sex TEXT DEFAULT 'undefined',
                         date_of_birth DATE,
                         age INT,
                         description TEXT,
                         cover_url TEXT,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         CONSTRAINT chk_age CHECK (age >= 0),
                         CONSTRAINT chk_sex CHECK (sex IN ('undefined', 'male', 'female', 'other')),
                         CONSTRAINT chk_phone_format CHECK (phone ~ '^\+?[1-9]\d{1,14}$'),
                         CONSTRAINT chk_date_of_birth CHECK (date_of_birth <= CURRENT_DATE),
                         CONSTRAINT chk_country CHECK (country IS NULL OR LENGTH(country) <= 100),
                         CONSTRAINT chk_language CHECK (language IS NULL OR LENGTH(language) <= 50)
);


CREATE TABLE category (
                          category_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          name TEXT NOT NULL UNIQUE CHECK (LENGTH(name) >= 4 AND LENGTH(name) <= 32),
                          description TEXT,
                          created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tag (
                     tag_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                     name TEXT NOT NULL UNIQUE CHECK (LENGTH(name) >= 4 AND LENGTH(name) <= 32),
                     created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE topic (
                       topic_id INT PRIMARY KEY,
                       title TEXT NOT NULL UNIQUE
);

CREATE TABLE article (
                         article_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         title TEXT NOT NULL CHECK (LENGTH(title) <= 200),
                         content TEXT NOT NULL CHECK (LENGTH(content) <= 100000),
                         author_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                         media_url TEXT,
                         topic_id INT NOT NULL REFERENCES topic(topic_id) ON DELETE NO ACTION,
                         status article_status NOT NULL DEFAULT 'draft',
                         comments_count INT DEFAULT 0,
                         reposts_count INT DEFAULT 0,
                         views_count INT DEFAULT 0,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         CONSTRAINT chk_status CHECK (status IN ('draft', 'published', 'archived')),
                         CONSTRAINT chk_comments_count CHECK (comments_count >= 0),
                         CONSTRAINT chk_reposts_count CHECK (reposts_count >= 0),
                         CONSTRAINT chk_views_count CHECK (views_count >= 0)
);

CREATE TABLE comment (
                         comment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         article_id UUID NOT NULL REFERENCES article(article_id) ON DELETE CASCADE,
                         user_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE RESTRICT,
                         content TEXT NOT NULL CHECK (LENGTH(content) <= 1000),
                         created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         reply_to UUID REFERENCES comment(comment_id) ON DELETE NO ACTION,
                         CONSTRAINT chk_content CHECK (LENGTH(content) > 0)
);

CREATE TABLE article_like (
                              user_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                              article_id UUID NOT NULL REFERENCES article(article_id) ON DELETE CASCADE,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (user_id, article_id)
);

CREATE TABLE comment_like (
                              user_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                              comment_id UUID NOT NULL REFERENCES comment(comment_id) ON DELETE CASCADE,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (user_id, comment_id)
);

CREATE TABLE media (
                       media_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       article_id UUID NOT NULL REFERENCES article(article_id) ON DELETE CASCADE,
                       uploader_id UUID REFERENCES "user"(user_id) ON DELETE SET NULL,
                       type media_type NOT NULL,
                       mime TEXT NOT NULL,
                       url TEXT NOT NULL,
                       size_bytes BIGINT,
                       description TEXT,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE subscription (
                              follower_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                              followed_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (follower_id, followed_id),
                              CHECK (follower_id <> followed_id)
);

CREATE TABLE notification (
                              notification_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              user_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                              type notification_type NOT NULL,
                              content TEXT NOT NULL CHECK (LENGTH(content) <= 500),
                              is_read BOOLEAN NOT NULL DEFAULT FALSE,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_user_updated_at
    BEFORE UPDATE ON "user"
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trigger_update_profile_updated_at
    BEFORE UPDATE ON profile
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trigger_update_article_updated_at
    BEFORE UPDATE ON article
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trigger_update_comment_updated_at
    BEFORE UPDATE ON comment
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();
