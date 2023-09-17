package publishing

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"

	"github.com/OrenRosen/gokit-example/article"
)

type Service interface {
	GetArticle(ctx context.Context, id string) (article.Article, error)
	CreateArticle(ctx context.Context, thing article.Article) (id string, err error)
}

// get article

type GetArticleRequestModel struct {
	ID string
}

type GetArticleResponseModel struct {
	Article article.Article
}

func MakeEndpointGetArticle(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetArticleRequestModel)
		if !ok {
			return nil, errors.New("MakeEndpointGetArticle failed cast request")
		}

		a, err := s.GetArticle(ctx, req.ID)
		if err != nil {
			return nil, fmt.Errorf("MakeEndpointGetArticle: %w", err)
		}

		return GetArticleResponseModel{
			Article: a,
		}, nil
	}
}

// create article

type CreateArticleRequestModel struct {
	Title string
	Text  string
}

func (r CreateArticleRequestModel) ToArticle() article.Article {
	return article.Article{
		Title: r.Title,
		Text:  r.Text,
	}
}

type CreateArticleResponseModel struct {
	ID string
}

func MakeEndpointCreateArticle(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(CreateArticleRequestModel)
		if !ok {
			return nil, errors.New("MakeEndpointCreateArticle failed cast request")
		}

		id, err := s.CreateArticle(ctx, req.ToArticle())
		if err != nil {
			return nil, fmt.Errorf("MakeEndpointCreateArticle: %w", err)
		}

		return CreateArticleResponseModel{
			ID: id,
		}, nil
	}
}
