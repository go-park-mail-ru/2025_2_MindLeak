package article

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestArticle(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T, mem *InMemoryArticle)
	}{
		{
			name: "CreateArticle creates new article",
			run: func(t *testing.T, mem *InMemoryArticle) {
				authorID := uuid.New()
				a, err := mem.CreateArticle(authorID, "Test Title", "Test Content")
				assert.NoError(t, err)
				assert.NotNil(t, a)
				assert.NotEqual(t, uuid.Nil, a.Id)
				assert.Equal(t, authorID, a.AuthorId)
				assert.Equal(t, "Test Title", a.Title)
				assert.Equal(t, "Test Content", a.Content)
				assert.NotEmpty(t, a.CreatedAt)
				assert.Equal(t, "https://st4.depositphotos.com/36740986/38337/i/450/depositphotos_383375990-stock-photo-collection-hundred-dollar-banknotes-female.jpg", a.Image)
				assert.Equal(t, "Алексей Владимиров", a.AuthorName)
				assert.Equal(t, "https://sun9-88.userapi.com/s/v1/ig2/P_e5HW2lWX3ZxayBg73NnzbHzyhxFCXtBseRjSrN_NbemNC78OpkeYfJeXcTOXqyR8NhSwizZKqJEq_R8PhQo607.jpg?quality=95&as=32x40,48x60,72x90,108x135,160x200,240x300,360x450,480x600,540x675,640x800,720x900,1080x1350,1280x1600,1440x1800,1620x2025&from=bu&cs=1620x0", a.AuthorAvatar)

				assert.Len(t, mem.Articles, 7)
				assert.Equal(t, a.Id, mem.Articles[6].Id)
			},
		},
		{
			name: "CreateArticle returns error if title already exists for author",
			run: func(t *testing.T, mem *InMemoryArticle) {
				authorID := uuid.New()
				_, _ = mem.CreateArticle(authorID, "Test Title", "Test Content")
				_, err := mem.CreateArticle(authorID, "Test Title", "New Content")
				assert.EqualError(t, err, "article with this title already exists for this author")
			},
		},
		{
			name: "GetArticleById returns existing article",
			run: func(t *testing.T, mem *InMemoryArticle) {
				authorID := uuid.New()
				a, _ := mem.CreateArticle(authorID, "Test Title", "Test Content")
				got, err := mem.GetArticleById(a.Id)
				assert.NoError(t, err)
				assert.Equal(t, a.Id, got.Id)
				assert.Equal(t, a.AuthorId, got.AuthorId)
				assert.Equal(t, a.Title, got.Title)
				assert.Equal(t, a.Content, got.Content)
				assert.Equal(t, a.CreatedAt, got.CreatedAt)
				assert.Equal(t, a.Image, got.Image)
				assert.Equal(t, a.AuthorName, got.AuthorName)
				assert.Equal(t, a.AuthorAvatar, got.AuthorAvatar)
			},
		},
		{
			name: "GetArticleById returns error if not found",
			run: func(t *testing.T, mem *InMemoryArticle) {
				_, err := mem.GetArticleById(uuid.New())
				assert.EqualError(t, err, "article not found")
			},
		},
		{
			name: "GetArticlesByAuthorId returns articles for author",
			run: func(t *testing.T, mem *InMemoryArticle) {
				authorID1 := uuid.New()
				authorID2 := uuid.New()
				_, _ = mem.CreateArticle(authorID1, "Title1", "Content1")
				_, _ = mem.CreateArticle(authorID1, "Title2", "Content2")
				_, _ = mem.CreateArticle(authorID2, "Title3", "Content3")
				result, err := mem.GetArticlesByAuthorId(authorID1)
				assert.NoError(t, err)
				assert.Len(t, result, 2)
				assert.Equal(t, "Title1", result[0].Title)
				assert.Equal(t, "Title2", result[1].Title)
			},
		},
		{
			name: "GetArticlesByAuthorId returns empty if no articles",
			run: func(t *testing.T, mem *InMemoryArticle) {
				result, err := mem.GetArticlesByAuthorId(uuid.New())
				assert.NoError(t, err)
				assert.Empty(t, result)
			},
		},
		{
			name: "GetAllArticles returns all articles",
			run: func(t *testing.T, mem *InMemoryArticle) {
				authorID := uuid.New()
				_, _ = mem.CreateArticle(authorID, "Title1", "Content1")
				_, _ = mem.CreateArticle(authorID, "Title2", "Content2")
				all, err := mem.GetAllArticles()
				assert.NoError(t, err)
				assert.Len(t, all, 2+6) // 6 mock articles from NewInMemoryArticle + 2 new
			},
		},
		{
			name: "DeleteArticle deletes existing article",
			run: func(t *testing.T, mem *InMemoryArticle) {
				authorID := uuid.New()
				a, _ := mem.CreateArticle(authorID, "Test Title", "Test Content")
				ok, err := mem.DeleteArticle(a.Id)
				assert.True(t, ok)
				assert.NoError(t, err)
				assert.Len(t, mem.Articles, 6)
			},
		},
		{
			name: "DeleteArticle returns error if not found",
			run: func(t *testing.T, mem *InMemoryArticle) {
				ok, err := mem.DeleteArticle(uuid.New())
				assert.False(t, ok)
				assert.EqualError(t, err, "article not found")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mem := NewInMemoryArticle()
			test.run(t, mem)
		})
	}
}
