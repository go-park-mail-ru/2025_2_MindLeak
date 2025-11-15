-- +goose Up
-- +goose StatementBegin

-- Наполнение таблицы topic
INSERT INTO topic (topic_id, title) VALUES
                                        (7, 'Инвестиции'),
                                        (6, 'Личный опыт'),
                                        (9, 'Образование'),
                                        (5, 'AI'),
                                        (3, 'Деньги'),
                                        (0, 'Без темы'),
                                        (4, 'Путешествия'),
                                        (8, 'Карьера'),
                                        (2, 'Маркетинг'),
                                        (1, 'Сервисы')
ON CONFLICT (topic_id) DO NOTHING;

-- Наполнение таблицы appeal_category
INSERT INTO appeal_category (name) VALUES
                                       ('Баг или техническая проблема'),
                                       ('Проблема с аккаунтом/авторизацией'),
                                       ('Предложение по функционалу'),
                                       ('Вопрос по использованию сервиса'),
                                       ('Жалоба или обратная связь'),
                                       ('Другое')
ON CONFLICT (name) DO NOTHING;

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

-- Удаляем только те топики, которые добавили этой миграцией
DELETE FROM topic
WHERE topic_id IN (0,1,2,3,4,5,6,7,8,9);

-- Удаляем только наши категории обращений
DELETE FROM appeal_category
WHERE name IN (
               'Баг или техническая проблема',
               'Проблема с аккаунтом/авторизацией',
               'Предложение по функционалу',
               'Вопрос по использованию сервиса',
               'Жалоба или обратная связь',
               'Другое'
    );

-- +goose StatementEnd
