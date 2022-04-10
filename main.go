package main

import (
	"content-service/config"
	"content-service/content"
	"content-service/content_type"
	"content-service/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Content Service")
	db, Sql, err  := config.ConnectPostgres()

	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	contentTypeRepository := content_type.NewRepository(db)
	contentTypeService := content_type.NewService(contentTypeRepository)
	contentTypeHandler := handler.NewHandler(contentTypeService)
	
	contentRepository := content.NewRepository(db)
	contentService := content.NewService(contentRepository)
	contentHandler := handler.NewContentHandler(contentService)



	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/content_types", contentTypeHandler.CreateContentType)
	api.GET("/content_types", contentTypeHandler.GetContentTypes)
	api.GET("/content_types/:id", contentTypeHandler.GetContentTypeById)
	api.PUT("/content_types/:id", contentTypeHandler.UpdateContentType)
	
	api.POST("/contents", contentHandler.CreateContent)
	api.GET("/contents/byContentTypeId/:contentTypeId", contentHandler.FindByContentTypeId)
	api.GET("/contents/byId/:id", contentHandler.FindById)
	api.PUT("/contents/byId/:id", contentHandler.GetContentIdInput)
	router.Run(":8080")
}