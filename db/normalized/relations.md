# Схема данных веб-сервиса (аналог vc.ru)

## Отношения и зависимости:

### Пользователи и профили:

#### Relation: `user`
**Описание:** Таблица хранения учетных записей пользователей для авторизации и регистрации.  
**Отношение:**  
USER (  
    USER_ID UUID PRIMARY KEY,  
    LOGIN TEXT NOT NULL UNIQUE,  
    PASSWORD_HASH TEXT NOT NULL,  
    EMAIL TEXT NOT NULL UNIQUE,  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP  
)  
**Функциональные зависимости:**  
{USER_ID} → LOGIN, PASSWORD_HASH, EMAIL, CREATED_AT, UPDATED_AT  
{LOGIN} → USER_ID, PASSWORD_HASH, EMAIL, CREATED_AT, UPDATED_AT  
{EMAIL} → USER_ID, LOGIN, PASSWORD_HASH, CREATED_AT, UPDATED_AT  

#### Relation: `user_profile`
**Описание:** Таблица хранения профилей пользователей для отображения публичной информации.  
**Отношение:**  
USER_PROFILE (  
    USER_ID UUID PRIMARY KEY,  
    DISPLAY_NAME TEXT NOT NULL,  
    BIO TEXT,  
    AVATAR_URL TEXT,  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    FOREIGN KEY (USER_ID) REFERENCES USER(USER_ID) ON DELETE CASCADE  
)  
**Функциональные зависимости:**  
{USER_ID} → DISPLAY_NAME, BIO, AVATAR_URL, CREATED_AT, UPDATED_AT  

### Статьи и контент:

#### Relation: `article`
**Описание:** Таблица хранения статей, создаваемых авторами.  
**Отношение:**  
ARTICLE (  
    ARTICLE_ID UUID PRIMARY KEY,  
    TITLE TEXT NOT NULL,  
    CONTENT TEXT NOT NULL,  
    AUTHOR_ID UUID NOT NULL,  
    PUBLISHED_AT TIMESTAMPTZ,  
    STATUS TEXT NOT NULL CHECK (STATUS IN ('draft', 'published', 'archived')),  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    FOREIGN KEY (AUTHOR_ID) REFERENCES USER(USER_ID) ON DELETE RESTRICT  
)  
**Функциональные зависимости:**  
{ARTICLE_ID} → TITLE, CONTENT, AUTHOR_ID, PUBLISHED_AT, STATUS, CREATED_AT, UPDATED_AT  

#### Relation: `category`
**Описание:** Таблица хранения категорий для классификации статей.  
**Отношение:**  
CATEGORY (  
    CATEGORY_ID UUID PRIMARY KEY,  
    NAME TEXT NOT NULL UNIQUE,  
    DESCRIPTION TEXT,  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP  
)  
**Функциональные зависимости:**  
{CATEGORY_ID} → NAME, DESCRIPTION, CREATED_AT, UPDATED_AT  
{NAME} → CATEGORY_ID, DESCRIPTION, CREATED_AT, UPDATED_AT  

#### Relation: `tag`
**Описание:** Таблица хранения тегов для пометки статей.  
**Отношение:**  
TAG (  
    TAG_ID UUID PRIMARY KEY,  
    NAME TEXT NOT NULL UNIQUE,  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP  
)  
**Функциональные зависимости:**  
{TAG_ID} → NAME, CREATED_AT  
{NAME} → TAG_ID, CREATED_AT  

#### Relation: `comment`
**Описание:** Таблица хранения комментариев к статьям.  
**Отношение:**  
COMMENT (  
    COMMENT_ID UUID PRIMARY KEY,  
    ARTICLE_ID UUID NOT NULL,  
    USER_ID UUID NOT NULL,  
    CONTENT TEXT NOT NULL,  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    FOREIGN KEY (ARTICLE_ID) REFERENCES ARTICLE(ARTICLE_ID) ON DELETE CASCADE,  
    FOREIGN KEY (USER_ID) REFERENCES USER(USER_ID) ON DELETE RESTRICT  
)  
**Функциональные зависимости:**  
{COMMENT_ID} → ARTICLE_ID, USER_ID, CONTENT, CREATED_AT, UPDATED_AT  

### Связующие таблицы:

#### Relation: `article_category`
**Описание:** Таблица реализации связи многие-ко-многим между статьями и категориями.  
**Отношение:**  
ARTICLE_CATEGORY (  
    ARTICLE_ID UUID NOT NULL,  
    CATEGORY_ID UUID NOT NULL,  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    PRIMARY KEY (ARTICLE_ID, CATEGORY_ID),  
    FOREIGN KEY (ARTICLE_ID) REFERENCES ARTICLE(ARTICLE_ID) ON DELETE CASCADE,  
    FOREIGN KEY (CATEGORY_ID) REFERENCES CATEGORY(CATEGORY_ID) ON DELETE RESTRICT  
)  
**Функциональные зависимости:**  
{ARTICLE_ID, CATEGORY_ID} → CREATED_AT  

#### Relation: `article_tag`
**Описание:** Таблица реализации связи многие-ко-многим между статьями и тегами.  
**Отношение:**  
ARTICLE_TAG (  
    ARTICLE_ID UUID NOT NULL,  
    TAG_ID UUID NOT NULL,  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    PRIMARY KEY (ARTICLE_ID, TAG_ID),  
    FOREIGN KEY (ARTICLE_ID) REFERENCES ARTICLE(ARTICLE_ID) ON DELETE CASCADE,  
    FOREIGN KEY (TAG_ID) REFERENCES TAG(TAG_ID) ON DELETE RESTRICT  
)  
**Функциональные зависимости:**  
{ARTICLE_ID, TAG_ID} → CREATED_AT  

#### Relation: `article_like`
**Описание:** Таблица хранения лайков пользователей к статьям.  
**Отношение:**  
ARTICLE_LIKE (  
    USER_ID UUID NOT NULL,  
    ARTICLE_ID UUID NOT NULL,  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    PRIMARY KEY (USER_ID, ARTICLE_ID),  
    FOREIGN KEY (USER_ID) REFERENCES USER(USER_ID) ON DELETE CASCADE,  
    FOREIGN KEY (ARTICLE_ID) REFERENCES ARTICLE(ARTICLE_ID) ON DELETE CASCADE  
)  
**Функциональные зависимости:**  
{USER_ID, ARTICLE_ID} → CREATED_AT  

#### Relation: `comment_like`
**Описание:** Таблица хранения лайков пользователей к комментариям.  
**Отношение:**  
COMMENT_LIKE (  
    USER_ID UUID NOT NULL,  
    COMMENT_ID UUID NOT NULL,  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    PRIMARY KEY (USER_ID, COMMENT_ID),  
    FOREIGN KEY (USER_ID) REFERENCES USER(USER_ID) ON DELETE CASCADE,  
    FOREIGN KEY (COMMENT_ID) REFERENCES COMMENT(COMMENT_ID) ON DELETE CASCADE  
)  
**Функциональные зависимости:**  
{USER_ID, COMMENT_ID} → CREATED_AT  

#### Relation: `notification`
**Описание:** Таблица хранения уведомлений для пользователей (например, о лайках или комментариях).  
**Отношение:**  
NOTIFICATION (  
    NOTIFICATION_ID UUID PRIMARY KEY,  
    USER_ID UUID NOT NULL,  
    TYPE TEXT NOT NULL CHECK (TYPE IN ('like', 'comment', 'follow')),  
    CONTENT TEXT NOT NULL,  
    IS_READ BOOLEAN NOT NULL DEFAULT FALSE,  
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,  
    FOREIGN KEY (USER_ID) REFERENCES USER(USER_ID) ON DELETE CASCADE  
)  
**Функциональные зависимости:**  
{NOTIFICATION_ID} → USER_ID, TYPE, CONTENT, IS_READ, CREATED_AT  

### Дополнительные хранилища:
- **Redis**: Используется для хранения сессий пользователей. Данные хранятся в формате ключ-значение, где ключ — это `session_id` (UUID), а значение — JSON с информацией о пользователе и сроке действия сессии. Не отображается в реляционной схеме, так как не является частью PostgreSQL.  
- **IO Storage**: Используется для хранения изображений (аватары, изображения в статьях). Ссылки на файлы хранятся в полях `AVATAR_URL` (таблицы `user_profile`, `article`) в формате TEXT.  

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