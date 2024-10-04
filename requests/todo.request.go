package requests

type TodoRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsCompleted bool   `json:"is_completed"`
}

type TodoUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}
