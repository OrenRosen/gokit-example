package inmem

import (
	"context"
	"fmt"

	"github.com/OrenRosen/gokit-example/article"
)

type articlesRepository struct {
	articles map[string]article.Article
}

func NewArticlesRepository() *articlesRepository {
	return &articlesRepository{
		articles: map[string]article.Article{},
	}
}

// not safe
func (r *articlesRepository) GetArticle(ctx context.Context, id string) (article.Article, error) {
	a, ok := r.articles[id]
	if !ok {
		return article.Article{}, fmt.Errorf("article wasn't found")
	}

	return a, nil
}

// not safe
func (r *articlesRepository) InsertArticle(ctx context.Context, article article.Article) error {
	r.articles[article.ID] = article
	return nil
}
