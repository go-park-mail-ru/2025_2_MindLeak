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
	Id        uuid.UUID
	AuthorId  uuid.UUID
	Title     string
	Content   string
	CreatedAt time.Time
}

type InMemoryArticle struct {
	Articles []*Article
	mu       sync.RWMutex
}

func NewInMemoryArticle() *InMemoryArticle {
	return &InMemoryArticle{
		Articles: make([]*Article, 0),
	}
}

func (mem *InMemoryArticle) CreateArticle(authorId uuid.UUID, title, content string) (*Article, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	for _, article := range mem.Articles {
		if article.Title == title && article.AuthorId == authorId {
			return nil, errors.New("article with this title already exists for this author")
		}
	}

	article := &Article{
		Id:        uuid.New(),
		AuthorId:  authorId,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
	mem.Articles = append(mem.Articles, article)
	return article, nil
}

func (mem *InMemoryArticle) GetArticleById(id uuid.UUID) (*Article, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	for _, article := range mem.Articles {
		if article.Id == id {
			return article, nil
		}
	}
	return nil, errors.New("article not found")
}

func (mem *InMemoryArticle) GetArticlesByAuthorId(authorId uuid.UUID) ([]*Article, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	var result []*Article
	for _, article := range mem.Articles {
		if article.AuthorId == authorId {
			result = append(result, article)
		}
	}
	return result, nil
}

func (mem *InMemoryArticle) GetAllArticles() ([]*Article, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	return mem.Articles, nil
}

func (mem *InMemoryArticle) DeleteArticle(id uuid.UUID) (bool, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	for idx, article := range mem.Articles {
		if article.Id == id {
			mem.Articles[idx] = mem.Articles[len(mem.Articles)-1]
			mem.Articles = mem.Articles[:len(mem.Articles)-1]
			return true, nil
		}
	}
	return false, errors.New("article not found")
}
