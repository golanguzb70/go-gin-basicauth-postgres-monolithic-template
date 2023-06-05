package postsrepo

type CreateReq struct {
	Title       string `json:"title"`
	Image       string `json:"image"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

type UpdateReq struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

type FindOneReq struct {
	Id int
}

type FindListReq struct {
	Page  int
	Limit int
}

type DeleteReq struct {
	Id int
}

type FullResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Content     string `json:"content"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
