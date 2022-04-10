package handler

import (
	"content-service/content_type"
	"content-service/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contentTypeHandler struct {
	service content_type.Service
}

func NewHandler(service content_type.Service) *contentTypeHandler {
	return &contentTypeHandler{service}
}


func (h *contentTypeHandler) CreateContentType(c *gin.Context) {
	input := content_type.Content_Type_Input{}
	
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors" : errors,
		}
		response := helper.APIResponse("Create Content Type Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	content_type_data, err := h.service.Create(input)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Create Content Type Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatContentTypeResponse := content_type.FormatContentTypeResponse(content_type_data)
	
	response := helper.APIResponse("Create Content Type Success", http.StatusOK, "success", formatContentTypeResponse)
	c.JSON(http.StatusOK, response)
}


func (h *contentTypeHandler) GetContentTypeById(c *gin.Context){
	var input content_type.GetContentTypeInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Get Content Type Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	content_type_data, err := h.service.GetById(input.ID)
	
	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Get Content Type Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatContentTypeResponse := content_type.FormatContentTypeResponse(content_type_data)
	
	response := helper.APIResponse("Get Content Type Success", http.StatusOK, "success", formatContentTypeResponse)
	c.JSON(http.StatusOK, response)
}



func (h *contentTypeHandler) GetContentTypes(c *gin.Context){
	content_type_datas, err := h.service.GetAll()
	
	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Get All Content Type Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatContentTypeResponse := content_type.FormatContentTypesResponse(content_type_datas)
	
	response := helper.APIResponse("Get All Content Type Success", http.StatusOK, "success", formatContentTypeResponse)
	c.JSON(http.StatusOK, response)
}

func (h *contentTypeHandler) UpdateContentType(c *gin.Context){
	var inputId content_type.GetContentTypeInput
	err := c.ShouldBindUri(&inputId)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Update Content Type Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	inputData := content_type.Content_Type_Input{}
	
	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors" : errors,
		}
		response := helper.APIResponse("Update Content Type Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	contentTypeUpdated, err := h.service.Update(inputId.ID, inputData)

	if err != nil {
		errorMessage := gin.H{
			"errors" : err.Error(),
		}
		response := helper.APIResponse("Update Content Type Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatContentTypeResponse := content_type.FormatContentTypeResponse(contentTypeUpdated)
	
	response := helper.APIResponse("Update Content Type Success", http.StatusOK, "success", formatContentTypeResponse)
	c.JSON(http.StatusOK, response)

}


