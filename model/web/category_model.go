package web

type CategoryCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type CategoryUpdateRequest struct {
	Id   uint64 `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}

type CategoryResponse struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
