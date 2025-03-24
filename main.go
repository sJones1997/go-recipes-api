package main

import (
	"log"
	"net/http"
    "recipe_rest_api/pkg/recipes"
)

func main() {

	mux := setupServer()
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func setupServer() *http.ServeMux {
	store := recipes.NewMemStore()
	recipesHandler := recipes.NewHandler(store)

	mux := http.NewServeMux()
	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes",  recipesHandler)
	mux.Handle("/recipes/", recipesHandler)
	log.Println("Running server on port 8080")

	return mux
}


//Home
type homeHandler struct{}
func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("This is my homepage"))
	if err != nil {
		return
	}
}