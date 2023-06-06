package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/models"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/logger"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/storage"
)

type TemplateI interface {
	Create(c *gin.Context)
	// Get(c *gin.Context)
	// Find(c *gin.Context)
	// Update(c *gin.Context)
	// Delete(c *gin.Context)
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

// @Summary		Create template
// @Tags        Template
// @Description	Here template can be created.
// @Security    BearerAuth
// @Accept      json
// @Produce		json
// @Param       post   body      models.TemplateCreateReq true "post info"
// @Success		200 	{object}  models.TemplateApiResponse
// @Failure     default {object}  models.DefaultResponse
// @Router		/template [POST]
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
