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
  topic_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title TEXT NOT NULL UNIQUE
);

CREATE TABLE article (
                         article_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         author_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                         title TEXT NOT NULL,
                         content TEXT NOT NULL,
                         topic_id UUID NOT NULL REFERENCES topic(topic_id) ON DELETE NO ACTION,
                         status article_status NOT NULL DEFAULT 'draft',
                         created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comment (
                         comment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         article_id UUID NOT NULL REFERENCES article(article_id) ON DELETE CASCADE,
                         user_id UUID NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
                         content TEXT NOT NULL,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
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

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS media;
DROP TABLE IF EXISTS article_like;
DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS article;
DROP TABLE IF EXISTS profile;
DROP TABLE IF EXISTS "user";
DROP TYPE IF EXISTS media_type;
DROP TYPE IF EXISTS article_status;
DROP FUNCTION IF EXISTS update_updated_at();
DROP EXTENSION IF EXISTS "pgcrypto";
-- +goose StatementEnd
