package models

type TemplateCreateReq struct {
	TemplateName string `json:"template_name"`
}

type TemplateUpdateReq struct {
	Id           int    `json:"id"`
	TemplateName string `json:"template_name"`
}

type TemplateGetReq struct {
	Id int `json:"id"`
}

type TemplateFindReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type TemplateDeleteReq struct {
	Id int `json:"id"`
}

type TemplateFindResponse struct {
	Templates []*TemplateResponse `json:"templates"`
	Count     int                 `json:"count"`
}

type TemplateResponse struct {
	Id           int    `json:"id"`
	TemplateName string `json:"template_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type TemplateApiResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *TemplateResponse
}

type TemplateApiFindResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Body         *TemplateFindResponse
}
