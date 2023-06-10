package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/models"
)

// @Router		/template [POST]
// @Summary		Create template
// @Tags        Template
// @Description	Here template can be created.
// @Security    BasicAuth
// @Accept      json
// @Produce		json
// @Param       post   body       models.TemplateCreateReq true "post info"
// @Success		200 	{object}  models.TemplateApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) TemplateCreate(c *gin.Context) {
	body := &models.TemplateCreateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.storage.Postgres().TemplateCreate(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().TemplateCreate()") {
		return
	}

	c.JSON(http.StatusOK, &models.TemplateApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/template/{id} [GET]
// @Summary		Get template by key
// @Tags        Template
// @Description	Here template can be got.
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.TemplateApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) TemplateGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	res, err := h.storage.Postgres().TemplateGet(context.Background(), &models.TemplateGetReq{
		Id: id,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().TemplateGet()") {
		return
	}

	c.JSON(http.StatusOK, models.TemplateApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/template/list [GET]
// @Summary		Get templates list
// @Tags        Template
// @Description	Here all templates can be got.
// @Accept      json
// @Produce		json
// @Param       filters query models.TemplateFindReq true "filters"
// @Success		200 	{object}  models.TemplateApiFindResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) TemplateFind(c *gin.Context) {
	page, err := ParsePageQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "helper.ParsePageQueryParam(c)") {
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "helper.ParseLimitQueryParam(c)") {
		return
	}

	res, err := h.storage.Postgres().TemplateFind(context.Background(), &models.TemplateFindReq{
		Page:  page,
		Limit: limit,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().TemplateFind()") {
		return
	}

	c.JSON(http.StatusOK, &models.TemplateApiFindResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Summary		Update template
// @Tags        Template
// @Description	Here template can be updated.
// @Security    BasicAuth
// @Accept      json
// @Produce		json
// @Param       post   body       models.TemplateUpdateReq true "post info"
// @Success		200 	{object}  models.TemplateApiResponse
// @Failure     default {object}  models.DefaultResponse
// @Router		/template [PUT]
func (h *handlerV1) TemplateUpdate(c *gin.Context) {
	body := &models.TemplateUpdateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.storage.Postgres().TemplateUpdate(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().TemplateUpdate()") {
		return
	}

	c.JSON(http.StatusOK, &models.TemplateApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/template/{id} [DELETE]
// @Summary		Delete template
// @Tags        Template
// @Description	Here template can be deleted.
// @Security    BasicAuth
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.DefaultResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) TemplateDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	err = h.storage.Postgres().TemplateDelete(context.Background(), &models.TemplateDeleteReq{Id: id})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().TemplateDelete()") {
		return
	}

	c.JSON(http.StatusOK, models.DefaultResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "Successfully deleted",
	})
}
