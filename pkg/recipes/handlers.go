package recipes

import (
	"encoding/json"
	"github.com/gosimple/slug"
	"log"
	"net/http"
	"regexp"
	"recipe_rest_api/pkg/httphelpers"
)

type Handler struct {
	store store
}

func NewHandler(s store) *Handler {
	return &Handler{
		store: s,
	}
}

var (
	RecipeRe 	   = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithId = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method)
	log.Println(r.URL.Path)

	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipes(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithId.MatchString(r.URL.Path):
		h.GetRecipe(w, r)
		return
	case r.Method == http.MethodPut && RecipeReWithId.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithId.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return
	}
}

func (h *Handler) CreateRecipe(w http.ResponseWriter, r *http.Request) {


	recipe, err := httphelpers.DecodeBody[Recipe](w, r)
	if err != nil {
		return
	}

	resourceID := slug.Make(recipe.Name)
	if err := h.store.Add(resourceID, recipe); err != nil {
		httphelpers.InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (h *Handler) ListRecipes(w http.ResponseWriter, r *http.Request) {

	recipes, err := h.store.List(); if err != nil {
		httphelpers.InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(recipes); if err != nil {
		httphelpers.InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
func (h *Handler) GetRecipe(w http.ResponseWriter, r *http.Request) {

	matches := RecipeReWithId.FindStringSubmatch(r.URL.Path)

	if len(matches) < 2 {
		httphelpers.InternalServerErrorHandler(w, r)
		return
	}

	recipe, err := h.store.Get(matches[1])

	if err != nil {
		if err == NotFoundErr {
			httphelpers.NotFoundHandler(w, r)
			return
		}

		httphelpers.InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(recipe); if err != nil {
		httphelpers.InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
func (h *Handler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {

	matches := RecipeReWithId.FindStringSubmatch(r.URL.Path)

	if len(matches) < 2 {
		httphelpers.InternalServerErrorHandler(w, r)
		return
	}

	recipe, err := httphelpers.DecodeBody[Recipe](w, r)
	if err != nil {
		return
	}

	if err := h.store.Update(matches[1], recipe); err  != nil {

		if(err == NotFoundErr){
			httphelpers.NotFoundHandler(w, r)
			return
		}

		httphelpers.InternalServerErrorHandler(w, r)
		return

	}

	jsonBytes, err := json.Marshal(recipe); if err != nil {
		httphelpers.InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
func (h *Handler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {

	matches := RecipeReWithId.FindStringSubmatch(r.URL.Path)

	if len(matches) < 2 {
		httphelpers.InternalServerErrorHandler(w, r)
		return
	}

	if err := h.store.Remove(matches[1]); err != nil {
		httphelpers.InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)

}