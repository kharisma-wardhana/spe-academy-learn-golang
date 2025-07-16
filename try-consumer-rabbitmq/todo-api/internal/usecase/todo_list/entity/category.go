package entity

// CategoryReq represents the request structure for creating or updating a Todo List Category.
type CategoryReq struct {
	ID          int64  `json:"id,omitempty" swaggerignore:"true"`
	Name        string `json:"name,omitempty" validate:"required" name:"Nama Kategori"`
	Description string `json:"description,omitempty" validate:"required" name:"Deskripsi Kategori"`
	CreatedBy   int64  `json:"created_by,omitempty" swaggerignore:"true"`
}

// CategoryResponse represents the response structure for a Todo List Category.
type CategoryResponse struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   int64  `json:"created_by"`
}

func (r *CategoryReq) SetID(ID int64) {
	r.ID = ID
}

func (r *CategoryReq) SetUserID(UserID int64) {
	r.CreatedBy = UserID
}
