package category

type EditCategoryRequest struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}
type AddCategoryRequest struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}
