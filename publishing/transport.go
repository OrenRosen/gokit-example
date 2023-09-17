package publishing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

type Router interface {
	Handle(method, path string, handler http.Handler)
}

func RegisterRoutes(router *httprouter.Router, s Service) {
	getArticleHandler := kithttp.NewServer(
		MakeEndpointGetArticle(s),
		decodeGetArticleRequest,
		encodeGetArticleResponse,
	)

	createArticleHandler := kithttp.NewServer(
		MakeEndpointCreateArticle(s),
		decodeCreateArticleRequest,
		encodeCreateArticleResponse,
	)

	router.Handler(http.MethodGet, "/articles/:id", getArticleHandler)
	router.Handler(http.MethodPost, "/articles", createArticleHandler)
}

// get article

func decodeGetArticleRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	params := httprouter.ParamsFromContext(ctx)
	return GetArticleRequestModel{
		ID: params.ByName("id"),
	}, nil
}

func encodeGetArticleResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(GetArticleResponseModel)
	if !ok {
		return fmt.Errorf("encodeGetArticleResponse failed cast response")
	}

	formatted := formatGetArticleResponse(res)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(formatted)
}

// create article

func decodeCreateArticleRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateArticleRequestModel
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("decodeCreateArticleRequest %s", err)
	}
	return req, nil
}

func encodeCreateArticleResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(CreateArticleResponseModel)
	if !ok {
		return fmt.Errorf("encodeCreateArticleResponse failed cast response")
	}

	formatted := formatCreateArticleResponse(res)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(formatted)
}
