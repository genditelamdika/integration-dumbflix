package categoriesdto

type CreateCategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
	// Email    string `json:"email" form:"email" validate:"required"`
	// Password string `json:"password" form:"password" validate:"required"`
	// Phone    string `json:"phone" form:"password" validate:"required"`
	// Gender   string `json:"gender" form:"password" validate:"required"`
	// Address  string `json:"addres" form:"password" validate:"required"`
	// Subcribe bool   `json:"subcribe" form:"password" `
}

type UpdateCategoryRequest struct {
	Name string `json:"name" form:"name"`
	// Email    string `json:"email" form:"email"`
	// Password string `json:"password" form:"password"`
}
