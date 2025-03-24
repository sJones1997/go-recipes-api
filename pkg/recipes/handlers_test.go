package recipes

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func setupTestHandler(t *testing.T) (*MemStore, http.Handler) {
	store := NewMemStore()
	handler := NewHandler(store)

	return store, handler
}

func performRequest(t *testing.T, handler http.Handler, method string, target string, body io.Reader) *http.Response  {
	req := httptest.NewRequest(method, target, body)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	res := w.Result()
	t.Cleanup(func() { res.Body.Close() })
	return res
}

func readTestData(t *testing.T, name string) []byte {
	t.Helper()
	content, err := os.ReadFile("../../testdata/" + name)

	if err != nil {
		t.Errorf("Could not read %v", name)
	}

	return content
}

func TestPostHandler(t *testing.T){

	_, handler := setupTestHandler(t)

	hamAndCheese := readTestData(t, "ham_and_cheese_recipe.json")
	hamAndCheeseReader := bytes.NewReader(hamAndCheese)

	res := performRequest(t, handler, http.MethodPost, "/recipes", hamAndCheeseReader)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

}

func TestGetAllRecipesHandler(t *testing.T) {

	store, handler := setupTestHandler(t)

	name := "ham_and_cheese_toastie"

	store.Add(name, Recipe{
		Name: name,
		Ingredients: []Ingredient{
			{Name: "ham"},
		},
	})

	res := performRequest(t, handler, http.MethodGet, "/recipes", nil)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetSingleRecipeHandler(t *testing.T) {

	store, handler := setupTestHandler(t)
	name := "ham_and_cheese_toastie"

	store.Add(name, Recipe{
		Name: name,
		Ingredients: []Ingredient{
			{Name: "ham"},
		},
	})

	res := performRequest(t, handler, http.MethodGet, "/recipes/" + name, nil)
	assert.Equal(t, http.StatusOK, res.StatusCode)

}

func TestUpdateRecipeHandler(t *testing.T) {

	store, handler := setupTestHandler(t)
	name := "ham_and_cheese_toastie"

	store.Add(name, Recipe{
		Name: name,
		Ingredients: []Ingredient{
			{Name: "ham"},
		},
	})

	hamAndCheeseWithButter := readTestData(t, "ham_and_cheese_with_butter_recipe.json")
	hamAndCheeseWithButterReader := bytes.NewReader(hamAndCheeseWithButter)

	res := performRequest(t, handler, http.MethodGet, "/recipes/" + name, hamAndCheeseWithButterReader)
	assert.Equal(t, http.StatusOK, res.StatusCode)

}

func TestDeleteRecipeHandler(t *testing.T) {

	store, handler := setupTestHandler(t)
	name := "ham_and_cheese_toastie"

	store.Add(name, Recipe{
		Name: name,
		Ingredients: []Ingredient{
			{Name: "ham"},
		},
	})


	res := performRequest(t, handler, http.MethodDelete, "/recipes/" + name, nil)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}