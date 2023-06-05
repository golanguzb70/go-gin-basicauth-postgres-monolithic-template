package aboutrepo

type CreateReq struct {
	Title      string
	Intro      string
	LinkedIn   string
	Youtube    string
	Facebook   string
	Telegram   string
	ResumeLink string
	Photo      string
}

type UpdateReq struct {
	Id         int
	Title      string
	Intro      string
	LinkedIn   string
	Youtube    string
	Facebook   string
	Telegram   string
	ResumeLink string
	Photo      string
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
	Id         int
	Title      string
	Intro      string
	LinkedIn   string
	Youtube    string
	Facebook   string
	Telegram   string
	ResumeLink string
	Photo      string
}
