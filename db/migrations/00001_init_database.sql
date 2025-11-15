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

CREATE TYPE appeal_status AS ENUM ('created', 'in_work', 'solved');

CREATE TABLE appeal_category (
                                 category_id SERIAL PRIMARY KEY,
                                 name TEXT UNIQUE NOT NULL
);

CREATE TABLE appeal (
                        appeal_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        creator_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                        email_registered TEXT NOT NULL,
                        category_id INT NOT NULL REFERENCES appeal_category(category_id),
                        status appeal_status NOT NULL DEFAULT 'created',
                        problem_description TEXT NOT NULL,
                        name TEXT NOT NULL,
                        email_for_connect TEXT NOT NULL,
                        screenshot_url TEXT NOT NULL,
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
                         comments_count INT,
                         reposts_count INT,
                         views_count INT,
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

CREATE TRIGGER trg_appeal_updated_at
    BEFORE UPDATE ON appeal
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS media CASCADE;
DROP TABLE IF EXISTS topic CASCADE;
DROP TABLE IF EXISTS article_like CASCADE;
DROP TABLE IF EXISTS comment CASCADE;
DROP TABLE IF EXISTS article CASCADE;
DROP TABLE IF EXISTS profile CASCADE;
DROP TABLE IF EXISTS subscription CASCADE;
DROP TABLE IF EXISTS "user" CASCADE;
DROP TYPE IF EXISTS media_type;
DROP TYPE IF EXISTS article_status;
DROP FUNCTION IF EXISTS update_updated_at();
DROP EXTENSION IF EXISTS "pgcrypto";
-- +goose StatementEnd
