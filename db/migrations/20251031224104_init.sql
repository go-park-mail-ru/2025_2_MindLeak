-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE article_status AS ENUM ('draft', 'published', 'archived');
CREATE TYPE notification_type AS ENUM ('like', 'comment', 'follow');

CREATE TABLE "user" (
                        user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        login TEXT NOT NULL UNIQUE CHECK (LENGTH(login) >= 4 AND LENGTH(login) <= 32 AND login !~ ' '),
    password_hash TEXT NOT NULL CHECK (LENGTH(password_hash) <= 255),
    email TEXT NOT NULL UNIQUE CHECK (LENGTH(email) <= 320 AND email ~ '^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$'),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_profile (
                              user_id UUID PRIMARY KEY,
                              display_name TEXT NOT NULL CHECK (LENGTH(display_name) <= 32),
                              bio TEXT CHECK (LENGTH(bio) <= 1000),
                              avatar_url TEXT CHECK (avatar_url ~ '^https?://[A-Za-z0-9.-]+\\.[A-Za-z]{2,}/.*$'),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user"(user_id) ON DELETE CASCADE
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

CREATE TABLE article (
                         article_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         title TEXT NOT NULL CHECK (LENGTH(title) <= 200),
                         content TEXT NOT NULL CHECK (LENGTH(content) <= 100000),
                         author_id UUID NOT NULL,
                         published_at TIMESTAMPTZ,
                         status article_status NOT NULL,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES "user"(user_id) ON DELETE RESTRICT
);

CREATE TABLE comment (
                         comment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         article_id UUID NOT NULL,
                         user_id UUID NOT NULL,
                         content TEXT NOT NULL CHECK (LENGTH(content) <= 1000),
                         created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         CONSTRAINT fk_article FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE,
                         CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user"(user_id) ON DELETE RESTRICT
);

CREATE TABLE article_category (
                                  article_id UUID NOT NULL,
                                  category_id UUID NOT NULL,
                                  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  PRIMARY KEY (article_id, category_id),
                                  CONSTRAINT fk_article FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE,
                                  CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES category(category_id) ON DELETE RESTRICT
);

CREATE TABLE article_tag (
                             article_id UUID NOT NULL,
                             tag_id UUID NOT NULL,
                             created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             PRIMARY KEY (article_id, tag_id),
                             CONSTRAINT fk_article FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE,
                             CONSTRAINT fk_tag FOREIGN KEY (tag_id) REFERENCES tag(tag_id) ON DELETE RESTRICT
);

CREATE TABLE article_like (
                              user_id UUID NOT NULL,
                              article_id UUID NOT NULL,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (user_id, article_id),
                              CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user"(user_id) ON DELETE CASCADE,
                              CONSTRAINT fk_article FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE
);

CREATE TABLE comment_like (
                              user_id UUID NOT NULL,
                              comment_id UUID NOT NULL,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (user_id, comment_id),
                              CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user"(user_id) ON DELETE CASCADE,
                              CONSTRAINT fk_comment FOREIGN KEY (comment_id) REFERENCES comment(comment_id) ON DELETE CASCADE
);

CREATE TABLE notification (
                              notification_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              user_id UUID NOT NULL,
                              type notification_type NOT NULL,
                              content TEXT NOT NULL CHECK (LENGTH(content) <= 500),
                              is_read BOOLEAN NOT NULL DEFAULT FALSE,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user"(user_id) ON DELETE CASCADE
);

CREATE TABLE profile (
                         profile_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         user_id UUID NOT NULL UNIQUE,
                         phone TEXT CHECK (LENGTH(phone) <= 20),
                         country TEXT CHECK (LENGTH(country) <= 64),
                         language TEXT CHECK (LENGTH(language) <= 32),
                         sex TEXT CHECK (sex IN ('male', 'female', 'undefined')) DEFAULT 'undefined',
                         date_of_birth DATE,
                         age INT CHECK (age >= 0),
                         cover_url TEXT CHECK (cover_url ~ '^https?://[A-Za-z0-9.-]+\\.[A-Za-z]{2,}/.*$'),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_profile_user FOREIGN KEY (user_id)
        REFERENCES "user"(user_id) ON DELETE CASCADE
);

CREATE TRIGGER trigger_update_profile_updated_at
    BEFORE UPDATE ON profile
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TYPE media_type AS ENUM ('image', 'video', 'audio', 'document', 'other');

CREATE TABLE media (
                       media_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       article_id UUID REFERENCES article(article_id) ON DELETE CASCADE,
                       uploader_id UUID REFERENCES "user"(user_id) ON DELETE SET NULL,
                       type media_type NOT NULL,
                       mime TEXT NOT NULL,
                       url TEXT NOT NULL CHECK (url ~ '^https?://[A-Za-z0-9.-]+\\.[A-Za-z]{2,}/.*$'),
    size_bytes BIGINT CHECK (size_bytes >= 0),
    description TEXT CHECK (LENGTH(description) <= 500),
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

CREATE TRIGGER trigger_update_user_profile_updated_at
    BEFORE UPDATE ON user_profile
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trigger_update_category_updated_at
    BEFORE UPDATE ON category
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trigger_update_article_updated_at
    BEFORE UPDATE ON article
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trigger_update_comment_updated_at
    BEFORE UPDATE ON comment
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notification;
DROP TABLE IF EXISTS comment_like;
DROP TABLE IF EXISTS article_like;
DROP TABLE IF EXISTS article_tag;
DROP TABLE IF EXISTS article_category;
DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS article;
DROP TABLE IF EXISTS tag;
DROP TABLE IF EXISTS category;
DROP TABLE IF EXISTS user_profile;
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS media;
DROP TYPE IF EXISTS media_type;
DROP TABLE IF EXISTS profile;
DROP TYPE IF EXISTS article_status;
DROP TYPE IF EXISTS notification_type;
DROP EXTENSION IF EXISTS "pgcrypto";
-- +goose StatementEnd
