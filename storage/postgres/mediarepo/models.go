package mediarepo

type CreateReq struct {
	Name string
}

type UpdateReq struct {
	Id   int
	Name string
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
	Id        int
	Name      string
	CreatedAt string
	UpdatedAt string
}
