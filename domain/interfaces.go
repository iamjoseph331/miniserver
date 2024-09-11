package domain

type ReceptorObject struct {
	UserID  string `json:"user_id"`
	Role    string `json:"role"`
	Message string `json:"message"`
}

type SigninQueryObject struct {
	UserID string `json:"user_id"`
	Query  string `json:"query"`
	Format string `json:"format"`
}
