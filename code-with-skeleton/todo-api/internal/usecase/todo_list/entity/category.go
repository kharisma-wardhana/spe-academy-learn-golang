package entity

type CategoryReq struct {
	ID          int64  `json:"id,omitempty" swaggerignore:"true"`
	Name        string `json:"name,omitempty" validate:"required" name:"Nama Kategori"`
	Description string `json:"description,omitempty" validate:"required" name:"Deskripsi Kategori"`
}

type CategoryResponse struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
