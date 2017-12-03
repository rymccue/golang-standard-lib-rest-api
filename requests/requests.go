package requests

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateJobRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateJobRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
