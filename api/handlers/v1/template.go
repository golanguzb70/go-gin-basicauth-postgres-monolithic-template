package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/models"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/logger"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/storage"
)

type TemplateI interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Find(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type TemplateHandler struct {
	log      *logger.Logger
	cfg      config.Config
	postgres storage.StorageI
}

func NewTemplateHandler(c *HandlerV1Config) TemplateI {
	return &TemplateHandler{
		log:      c.Logger,
		cfg:      c.Cfg,
		postgres: c.Postgres,
	}
}

// @Router		/template [POST]
// @Summary		Create template
// @Tags        Template
// @Description	Here template can be created.
// @Security    BearerAuth
// @Accept      json
// @Produce		json
// @Param       post   body       models.TemplateCreateReq true "post info"
// @Success		200 	{object}  models.TemplateApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *TemplateHandler) Create(c *gin.Context) {
	body := &models.TemplateCreateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.postgres.Template().Create(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.postgres.Template().Create()") {
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
func (h *TemplateHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	res, err := h.postgres.Template().Get(context.Background(), &models.TemplateGetReq{
		Id: id,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.postgres.Template().Get()") {
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
func (h *TemplateHandler) Find(c *gin.Context) {
	page, err := ParsePageQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "helper.ParsePageQueryParam(c)") {
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "helper.ParseLimitQueryParam(c)") {
		return
	}

	res, err := h.postgres.Template().Find(context.Background(), &models.TemplateFindReq{
		Page:  page,
		Limit: limit,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.postgres.Template().Find()") {
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
// @Security    BearerAuth
// @Accept      json
// @Produce		json
// @Param       post   body       models.TemplateUpdateReq true "post info"
// @Success		200 	{object}  models.TemplateApiResponse
// @Failure     default {object}  models.DefaultResponse
// @Router		/template [PUT]
func (h *TemplateHandler) Update(c *gin.Context) {
	body := &models.TemplateUpdateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.postgres.Template().Update(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.postgres.Template().Update()") {
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
// @Security    BearerAuth
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.DefaultResponse
// @Failure     default {object}  models.DefaultResponse
func (h *TemplateHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	err = h.postgres.Template().Delete(context.Background(), &models.TemplateDeleteReq{Id: id})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.postgres.Template().Delete()") {
		return
	}

	c.JSON(http.StatusOK, models.DefaultResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "Successfully deleted",
	})
}
