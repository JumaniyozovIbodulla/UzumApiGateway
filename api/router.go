package api

import (
	"net/http"

	_ "uzumapi/api/docs" //for swagger
	"uzumapi/api/handler"
	"uzumapi/config"
	"uzumapi/pkg/grpc_client"
	"uzumapi/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// @title           Uzum Clone Api
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})

	// orders
	r.POST("/api/v1/order", handler.CreateOrder)
	r.GET("/api/v1/order/:id", handler.GetByIdOrder)
	r.PUT("/api/v1/order", handler.UpdateOrder)
	r.DELETE("/api/v1/order/:id", handler.DeleteOrder)
	r.GET("/api/v1/orders", handler.GetAllOrder)

	// order-products
	r.POST("/api/v1/orderProduct", handler.CreateOrderProduct)
	r.GET("/api/v1/orderProduct/:id", handler.GetByIdOrderProduct)
	r.PUT("/api/v1/orderProduct", handler.UpdateOrderProduct)
	r.DELETE("/api/v1/orderProduct/:id", handler.DeleteOrderProduct)
	r.GET("/api/v1/orderProducts", handler.GetAllOrderProducts)

	// order-notes
	r.POST("/api/v1/orderNote", handler.CreateOrderNotes)
	r.GET("/api/v1/orderNote/:id", handler.GetByIdOrderNotes)
	r.PUT("/api/v1/orderNote", handler.UpdateOrderNotes)
	r.DELETE("/api/v1/orderNote/:id", handler.DeleteOrderNotes)
	r.GET("/api/v1/orderNotes", handler.GetAllOrderNotes)

	// Shipper endpoints
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
