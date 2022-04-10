package handler

import (
	"content-service/content"
	"content-service/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contentHandler struct {
	contentService content.Service
}

func NewContentHandler(contentService content.Service) *contentHandler{
	return &contentHandler{contentService}
}

func (h *contentHandler) CreateContent(c *gin.Context){
	input := content.CreateContentInput{}
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors" : errors,
		}
		response := helper.APIResponse("Create Content Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	content_data, err := h.contentService.Create(input)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Create Content Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatCreateContentResponse := content.FormatCreateContentResponse(content_data)
	response := helper.APIResponse("Create Content Success", http.StatusOK, "success", formatCreateContentResponse)
	c.JSON(http.StatusOK, response)
}


func (h *contentHandler) FindByContentTypeId(c *gin.Context) {
	input := content.GetContentByContentTypeIdInput{}
	err := c.ShouldBindUri(&input)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Get Content By Content Type Id Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	contents_data, err := h.contentService.FindByContentTypeId(input.ID)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Get Content By Content Type Id Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatGetContentsByContentTypeIdResponse := content.FormatGetContentsByContentTypeIdResponse(contents_data)
	
	response := helper.APIResponse("Get Content By Content Type Id Success", http.StatusOK, "success", formatGetContentsByContentTypeIdResponse)
	c.JSON(http.StatusOK, response)

}


func (h *contentHandler) FindById(c *gin.Context) {
	input := content.GetContentIdInput{}
	err := c.ShouldBindUri(&input)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Get Content By Id Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	content_data, err := h.contentService.FindById(input.ID)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Get Content By Id Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatGetContentByIdResponse := content.FormatGetContentResponse(content_data)
	
	response := helper.APIResponse("Get Content By Id Success", http.StatusOK, "success", formatGetContentByIdResponse)
	c.JSON(http.StatusOK, response)
}

func (h *contentHandler) GetContentIdInput(c *gin.Context) {
	inputId := content.GetContentIdInput{}
	err := c.ShouldBindUri(&inputId)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Update Content Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputData := content.Content_Input{}
	
	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors" : errors,
		}
		response := helper.APIResponse("Update Content Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}


	contentUpdated, err := h.contentService.UpdateContentById(inputId.ID, inputData)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Update Content Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatContentResponse := content.FormatUpdateContentResponse(contentUpdated)
	
	response := helper.APIResponse("Update Content Type Success", http.StatusOK, "success", formatContentResponse)
	c.JSON(http.StatusOK, response)

}