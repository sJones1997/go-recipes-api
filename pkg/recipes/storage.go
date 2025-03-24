package recipes

//Data Storage
type store interface {
	Add(name string, recipe Recipe) error
	Get(name string) (Recipe, error)
	Update(name string, recipe Recipe) error
	List() (map[string]Recipe, error)
	Remove(name string) error
}