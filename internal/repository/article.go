package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ArticleRepository interface {
	CreateArticle(authorId uuid.UUID, title, content string) (*Article, error)
	GetArticleById(id uuid.UUID) (*Article, error)
	GetArticlesByAuthorId(authorId uuid.UUID) ([]*Article, error)
	GetAllArticles() ([]*Article, error)
	DeleteArticle(id uuid.UUID) (bool, error)
}

type Article struct {
	Id           uuid.UUID `json:"-"`
	AuthorId     uuid.UUID `json:"-"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"-"`
	Image        string    `json:"image"`
	AuthorName   string    `json:"author_name"`
	AuthorAvatar string    `json:"author_avatar"`
}

type InMemoryArticle struct {
	Articles []Article
	mu       sync.RWMutex
}

func NewInMemoryArticle() *InMemoryArticle {
	articles := &InMemoryArticle{
		Articles: make([]Article, 0),
	}
	authorID := uuid.New()

	_, _ = articles.CreateArticle(authorID,
		"ИИ в 2025: Как нейросети меняют бизнес-процессы",
		"Искусственный интеллект в 2025 году стал неотъемлемой частью бизнеса...")

	_, _ = articles.CreateArticle(authorID,
		"Как российский стартап привлёк $10M на рынке SaaS",
		"Российский стартап CloudPeak разработал SaaS-платформу...")

	_, _ = articles.CreateArticle(authorID,
		"Тренды контент-маркетинга: Что работает в 2025 году",
		"Контент-маркетинг в 2025 году переживает новый виток...")

	_, _ = articles.CreateArticle(authorID,
		"Почему 80% стартапов терпят неудачу в первый год",
		"Запуск стартапа — это всегда риск...")

	_, _ = articles.CreateArticle(authorID,
		"Как мы увеличили конверсию на 30% с помощью UX",
		"Компания BrightPath переработала интерфейс...")

	_, _ = articles.CreateArticle(authorID,
		"Экспериментальный сверхдлинный заголовок статьи, в котором мы попробуем уместить сразу и суть, и интригу, и даже немного юмора, чтобы проверить, как фронтенд справится с рендерингом текста...",
		`Это тестовое содержимое статьи, которое специально сделано очень длинным, чтобы проверить работу фронтенда с большими объёмами текста... (длинный текст)`)

	return articles
}

func (mem *InMemoryArticle) CreateArticle(authorID uuid.UUID, title, content string) (*Article, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	for _, article := range mem.Articles {
		if article.Title == title && article.AuthorId == authorID {

			return nil, errors.New("article with this title already exists for this author")
		}
	}

	article := Article{
		Id:           uuid.New(),
		AuthorId:     authorID,
		Title:        title,
		Content:      content,
		CreatedAt:    time.Now(),
		AuthorName:   "Алексей Владимиров",
		AuthorAvatar: "https://sun9-88.userapi.com/s/v1/ig2/P_e5HW2lWX3ZxayBg73NnzbHzyhxFCXtBseRjSrN_NbemNC78OpkeYfJeXcTOXqyR8NhSwizZKqJEq_R8PhQo607.jpg?quality=95&as=32x40,48x60,72x90,108x135,160x200,240x300,360x450,480x600,540x675,640x800,720x900,1080x1350,1280x1600,1440x1800,1620x2025&from=bu&cs=1620x0",
		Image:        "https://st4.depositphotos.com/36740986/38337/i/450/depositphotos_383375990-stock-photo-collection-hundred-dollar-banknotes-female.jpg",
	}
	mem.Articles = append(mem.Articles, article)
	copyArticle := article
	return &copyArticle, nil
}

func (mem *InMemoryArticle) GetArticleById(articleID uuid.UUID) (*Article, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	for i := range mem.Articles {
		if mem.Articles[i].Id == articleID {
			copyArticle := mem.Articles[i]
			return &copyArticle, nil
		}
	}

	return nil, errors.New("article not found")
}

func (mem *InMemoryArticle) GetArticlesByAuthorId(authorId uuid.UUID) ([]*Article, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	var result []*Article
	for i := range mem.Articles {
		if mem.Articles[i].AuthorId == authorId {
			temp := mem.Articles[i]
			result = append(result, &temp)
		}
	}

	return result, nil
}

func (mem *InMemoryArticle) GetAllArticles() ([]*Article, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	articlesCopy := make([]*Article, len(mem.Articles))
	for i := range mem.Articles {
		temp := mem.Articles[i]
		articlesCopy[i] = &temp
	}
	return articlesCopy, nil
}

func (mem *InMemoryArticle) DeleteArticle(articleID uuid.UUID) (bool, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	for idx, article := range mem.Articles {
		if article.Id == articleID {
			mem.Articles[idx] = mem.Articles[len(mem.Articles)-1]
			mem.Articles = mem.Articles[:len(mem.Articles)-1]

			return true, nil
		}
	}

	return false, errors.New("article not found")
}
