-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE "user" (
                        user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        email TEXT NOT NULL UNIQUE,
                        password TEXT NOT NULL,
                        name TEXT NOT NULL,
                        avatar TEXT,
                        created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
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
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE article_status AS ENUM ('draft', 'published', 'archived');

CREATE TABLE topic (
                       topic_id INT PRIMARY KEY,
                       title TEXT NOT NULL UNIQUE
);

CREATE TABLE article (
                         article_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         author_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                         title TEXT NOT NULL,
                         content TEXT NOT NULL,
                         media_url TEXT,
                         topic_id INT NOT NULL REFERENCES topic(topic_id) ON DELETE NO ACTION,
                         status article_status NOT NULL DEFAULT 'draft',
                         comments_count INT NOT NULL DEFAULT 0,
                         reposts_count INT NOT NULL DEFAULT 0,
                         views_count INT NOT NULL DEFAULT 0,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comment (
                         comment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         article_id UUID NOT NULL REFERENCES article(article_id) ON DELETE NO ACTION,
                         user_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE NO ACTION,
                         content TEXT NOT NULL,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         reply_to UUID REFERENCES comment(comment_id) ON DELETE NO ACTION
);

CREATE TABLE article_like (
                              user_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                              article_id UUID NOT NULL REFERENCES article(article_id) ON DELETE CASCADE,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (user_id, article_id)
);

CREATE TYPE media_type AS ENUM ('image', 'video', 'audio', 'document', 'other');

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


-- ----------------------
-- UPDATE TIMESTAMPS TRIGGERS
-- ----------------------
CREATE OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_user_updated_at
    BEFORE UPDATE ON "user"
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trg_profile_updated_at
    BEFORE UPDATE ON profile
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trg_article_updated_at
    BEFORE UPDATE ON article
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trg_comment_updated_at
    BEFORE UPDATE ON comment
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();


-- ----------------------
-- COMMENTS COUNTER TRIGGERS
-- ----------------------

-- increment after INSERT
CREATE OR REPLACE FUNCTION inc_comments_count()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE article
    SET comments_count = comments_count + 1
    WHERE article_id = NEW.article_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_inc_comments
    AFTER INSERT ON comment
    FOR EACH ROW EXECUTE FUNCTION inc_comments_count();


-- decrement after DELETE
CREATE OR REPLACE FUNCTION dec_comments_count()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE article
    SET comments_count = comments_count - 1
    WHERE article_id = OLD.article_id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_dec_comments
    AFTER DELETE ON comment
    FOR EACH ROW EXECUTE FUNCTION dec_comments_count();

-- +goose StatementEnd




-- +goose Down
-- +goose StatementBegin

-- Remove comment count triggers
DROP TRIGGER IF EXISTS trg_inc_comments ON comment CASCADE;
DROP TRIGGER IF EXISTS trg_dec_comments ON comment CASCADE;

DROP FUNCTION IF EXISTS inc_comments_count CASCADE;
DROP FUNCTION IF EXISTS dec_comments_count CASCADE;

-- Remove updated_at triggers
DROP TRIGGER IF EXISTS trg_user_updated_at ON "user" CASCADE;
DROP TRIGGER IF EXISTS trg_profile_updated_at ON profile CASCADE;
DROP TRIGGER IF EXISTS trg_article_updated_at ON article CASCADE;
DROP TRIGGER IF EXISTS trg_comment_updated_at ON comment CASCADE;

-- Remove tables
DROP TABLE IF EXISTS media CASCADE;
DROP TABLE IF EXISTS topic CASCADE;
DROP TABLE IF EXISTS article_like CASCADE;
DROP TABLE IF EXISTS comment CASCADE;
DROP TABLE IF EXISTS article CASCADE;
DROP TABLE IF EXISTS profile CASCADE;
DROP TABLE IF EXISTS subscription CASCADE;
DROP TABLE IF EXISTS "user" CASCADE;

-- Remove types
DROP TYPE IF EXISTS media_type;
DROP TYPE IF EXISTS article_status;

-- Remove update function
DROP FUNCTION IF EXISTS update_updated_at CASCADE;

-- Remove pgcrypto extension
DROP EXTENSION IF EXISTS "pgcrypto";

-- +goose StatementEnd
