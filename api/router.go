package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/api/docs" // docs
	v1 "github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/api/handlers/v1"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/config"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/pkg/logger"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/storage"
	"github.com/golanguzb70/middleware/gin/basicauth"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf     config.Config
	Logger   *logger.Logger
	Postgres storage.StorageI
}

// New ...
// @title           Template project API Endpoints
// @version         1.0
// @description     Here QA can test and frontend or mobile developers can get information of API endpoints.

// @BasePath  /v1

// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
func New(log *logger.Logger, cfg config.Config, strg storage.StorageI) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	h := v1.New(&v1.HandlerV1Config{
		Logger:   log,
		Cfg:      cfg,
		Postgres: strg,
	})

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	authConf := basicauth.Config{
		Users: []basicauth.User{
			{
				UserName: cfg.AdminUsername,
				Password: cfg.AdminPassword,
			},
		},
		RequireAuthForAll: true,
	}

	router.Use(basicauth.New(&authConf).Middleware)
	api := router.Group("/v1")

	template := api.Group("/template")
	template.POST("", h.TemplateCreate)
	template.GET("/:id", h.TemplateGet)
	template.GET("/list", h.TemplateFind)
	template.PUT("", h.TemplateUpdate)
	template.DELETE(":id", h.TemplateDelete)

	// Don't delete this line, it is used to modify the file automatically

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
