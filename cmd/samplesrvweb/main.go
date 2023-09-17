package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/OrenRosen/gokit-example/inmem"
	"github.com/OrenRosen/gokit-example/publishing"
)

func main() {
	inmemRepo := inmem.NewArticlesRepository()
	articleService := publishing.NewService(inmemRepo)

	router := httprouter.New()

	publishing.RegisterRoutes(router, articleService)

	fmt.Println("Listening on :9999")
	http.ListenAndServe(":9999", router)
}
