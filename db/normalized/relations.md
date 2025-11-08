# Схема данных веб-сервиса (аналог vc.ru)

## Отношения и зависимости:

### Пользователи и профили:

#### Relation: `user`
**Описание:** Таблица хранения учетных записей пользователей для авторизации и регистрации.  
**Отношение:**  
USER (  
user_id UUID PK DEFAULT gen_random_uuid(),
login TEXT NOT NULL UNIQUE,
password TEXT NOT NULL,
email TEXT NOT NULL UNIQUE,
name TEXT NOT NULL,
avatar TEXT,
created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)  
**Функциональные зависимости:**  
`{user_id} → {login, password, email, name, avatar, created_at, updated_at}`  
`{login} → {user_id, password, email, name, avatar, created_at, updated_at}`  
`{email} → {user_id, login, password, name, avatar, created_at, updated_at}`

#### Relation: `profile`
**Описание:** Таблица хранения профилей пользователей для отображения публичной информации.  
**Отношение:**  
PROFILE (  
profile_id UUID PK DEFAULT gen_random_uuid(),
user_id UUID NOT NULL UNIQUE REFERENCES user(user_id) ON DELETE CASCADE,
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
)  
**Функциональные зависимости:**  
`{user_id} → {profile_id, phone, country, language, sex, date_of_birth, age, description, cover_url, created_at, updated_at}`  
`{profile_id} → {user_id, ...}`

### Статьи и контент:

#### Relation: `article`
**Описание:** Таблица хранения статей, создаваемых авторами.  
**Отношение:**  
ARTICLE (  
article_id UUID PK DEFAULT gen_random_uuid(),
title TEXT NOT NULL,
content TEXT NOT NULL,
author_id UUID NOT NULL REFERENCES user(user_id) ON DELETE CASCADE,
media_url TEXT,
topic_id INT NOT NULL REFERENCES topic(topic_id) ON DELETE NO ACTION,
status article_status NOT NULL DEFAULT 'draft',
comments_count INT DEFAULT 0,
reposts_count INT DEFAULT 0,
views_count INT DEFAULT 0,
created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)  
**Функциональные зависимости:**  
`{article_id} → {title, content, author_id, media_url, topic_id, status, comments_count, reposts_count, views_count, created_at, updated_at}`

#### Relation: `tag`
**Описание:** Таблица хранения тегов для пометки статей.  
**Отношение:**  
TAG (
tag_id UUID PK DEFAULT gen_random_uuid(),
name TEXT NOT NULL UNIQUE,
created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)
**Функциональные зависимости:**  
`{tag_id} → {name, created_at}`  
`{name} → {tag_id, created_at}`

#### Relation: `topic`
**Описание:** Таблица хранения топиков статей.  
**Отношение:**
TOPIC (
topic_id INT PK,
title TEXT NOT NULL UNIQUE
)
**Функциональные зависимости:**  
`{topic_id} → {title}`  
`{title} → {topic_id}`

#### Relation: `comment`
**Описание:** Таблица хранения комментариев к статьям.  
**Отношение:**  
COMMENT (  
comment_id UUID PK DEFAULT gen_random_uuid(),
article_id UUID NOT NULL REFERENCES article(article_id) ON DELETE CASCADE,
user_id UUID NOT NULL REFERENCES user(user_id) ON DELETE RESTRICT,
content TEXT NOT NULL,
created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
reply_to UUID REFERENCES comment(comment_id) ON DELETE NO ACTION
)  
**Функциональные зависимости:**  
`{comment_id} → {article_id, user_id, content, created_at, updated_at, reply_to}`

#### Relation: `media`
**Описание:** Таблица для хранения медиа-файлов(сами медиа лежат в minIO).  
**Отношение:**  
MEDIA (
media_id UUID PK DEFAULT gen_random_uuid(),
article_id UUID NOT NULL REFERENCES article(article_id) ON DELETE CASCADE,
uploader_id UUID REFERENCES user(user_id) ON DELETE SET NULL,
type media_type NOT NULL,
mime TEXT NOT NULL,
url TEXT NOT NULL,
size_bytes BIGINT,
description TEXT,
created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)
**Функциональные зависимости:**  
`{media_id} → {article_id, uploader_id, type, mime, url, size_bytes, description, created_at}`

#### Relation: `article_like`
**Описание:** Таблица хранения лайков пользователей к статьям.  
**Отношение:**  
ARTICLE_LIKE (  
user_id UUID NOT NULL REFERENCES user(user_id) ON DELETE CASCADE,
article_id UUID NOT NULL REFERENCES article(article_id) ON DELETE CASCADE,
created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (user_id, article_id)
)  
**Функциональные зависимости:**  
`{user_id, article_id} → {created_at}`

#### Relation: `comment_like`
**Описание:** Таблица хранения лайков пользователей к комментариям.  
**Отношение:**  
COMMENT_LIKE (  
user_id UUID NOT NULL REFERENCES user(user_id) ON DELETE CASCADE,
comment_id UUID NOT NULL REFERENCES comment(comment_id) ON DELETE CASCADE,
created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (user_id, comment_id) 
)  
**Функциональные зависимости:**  
`{user_id, comment_id} → {created_at}`

#### Relation: `notification`
**Описание:** Таблица хранения уведомлений для пользователей (например, о лайках или комментариях).  
**Отношение:**  
NOTIFICATION (  
Nnotification_id UUID PK DEFAULT gen_random_uuid(),
user_id UUID NOT NULL REFERENCES user(user_id) ON DELETE CASCADE,
type notification_type NOT NULL,
content TEXT NOT NULL,
is_read BOOLEAN NOT NULL DEFAULT FALSE,
created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)  
**Функциональные зависимости:**  
`{notification_id} → {user_id, type, content, is_read, created_at}`

#### Relation: `subscription`
**Описание:** Таблица хранения подписок.  
**Отношение:**  
SUBSCRIPTION (
follower_id UUID NOT NULL REFERENCES user(user_id) ON DELETE CASCADE,
followed_id UUID NOT NULL REFERENCES user(user_id) ON DELETE CASCADE,
created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (follower_id, followed_id),
CHECK (follower_id <> followed_id)
)
**Функциональные зависимости:**  
`{follower_id, followed_id} → {created_at}`

### Дополнительные хранилища:
- **Redis**: Используется для хранения сессий пользователей. Данные хранятся в формате ключ-значение, где ключ — это `session_id` (UUID), а значение — JSON с информацией о пользователе и сроке действия сессии. Не отображается в реляционной схеме, так как не является частью PostgreSQL.
- **minIO**: Используется для хранения изображений (аватары, изображения в статьях). Ссылки на файлы хранятся в полях `AVATAR_URL` (таблицы `user_profile`, `article`) в формате TEXT.

## Соответствие требованиям:

### Соответствие 1НФ:
Схема соответствует 1 нормальной форме, так как:
- Каждый из атрибутов атомарен (нет составных типов данных, таких как массивы или JSON).
- Каждое значение в одном столбце имеет один и тот же тип данных (UUID, TEXT, TIMESTAMPTZ, BOOLEAN).
- Порядок строк и столбцов не имеет значения.

### Соответствие 2НФ:
Схема соответствует 2 нормальной форме, так как:
- Она соответствует 1 нормальной форме.
- Неключевые атрибуты функционально зависят от всего первичного ключа (в таблицах `user`, `user_profile`, `article`, `category`, `tag`, `comment`, `notification`).
- В таблицах с составным ключом (`article_category`, `article_tag`, `article_like`, `comment_like`) неключевые атрибуты (`CREATED_AT`) зависят от всего составного ключа.

### Соответствие 3НФ:
Схема соответствует 3 нормальной форме, так как:
- Соответствует 1 и 2 нормальным формам.
- Нет транзитивных зависимостей между неключевыми атрибутами (например, в таблице `article` атрибут `CONTENT` зависит только от `ARTICLE_ID`, а не от `AUTHOR_ID`).

### Соответствие НФБК:
Схема соответствует нормальной форме Бойса-Кодда, так как:
- Соответствует 1, 2 и 3 нормальным формам.
- Все функциональные зависимости имеют в качестве детерминанта кандидат ключа (например, в таблице `user` поля `LOGIN` и `EMAIL` являются уникальными ключами, а зависимости от них покрывают все атрибуты).
