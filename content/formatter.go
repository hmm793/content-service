package content

type CreateContentResponseFormatter struct {
	ID            int    `json:"id"`
	Data          string `json:"data"`
	IsPublished   bool   `json:"is_published"`
	ContentTypeId int    `json:"content_type_id"`
	CreatedByName string `json:"created_by_name"`
	CreatedBy     int    `json:"created_by"`
}

type GetContentByContentTypeIdResponseFormatter struct {
	ID            int    `json:"id"`
	Data          string `json:"data"`
	IsPublished   bool   `json:"is_published"`
	ContentTypeId int    `json:"content_type_id"`
	CreatedBy     int    `json:"created_by"`
	DeletedBy     int    `json:"deleted_by"`
	UpdatedBy     int    `json:"updated_by"`
	CreatedByName string `json:"created_by_name"`
	DeletedByName string `json:"deleted_by_name"`
	UpdatedByName string `json:"updated_by_name"`
}
type GetContentByIdResponseFormatter struct {
	ID            int    `json:"id"`
	Data          string `json:"data"`
	IsPublished   bool   `json:"is_published"`
	ContentTypeId int    `json:"content_type_id"`
	CreatedBy     int    `json:"created_by"`
	DeletedBy     int    `json:"deleted_by"`
	UpdatedBy     int    `json:"updated_by"`
	CreatedByName string `json:"created_by_name"`
	DeletedByName string `json:"deleted_by_name"`
	UpdatedByName string `json:"updated_by_name"`
}

type UpdateContentByIdResponseFormatter struct {
	ID            int    `json:"id"`
	Data          string `json:"data"`
	IsPublished   bool   `json:"is_published"`
	ContentTypeId int    `json:"content_type_id"`
	CreatedBy     int    `json:"created_by"`
	DeletedBy     int    `json:"deleted_by"`
	UpdatedBy     int    `json:"updated_by"`
	CreatedByName string `json:"created_by_name"`
	DeletedByName string `json:"deleted_by_name"`
	UpdatedByName string `json:"updated_by_name"`
}

func FormatCreateContentResponse(content Content) CreateContentResponseFormatter {
	createContentFormatter := CreateContentResponseFormatter{}
	createContentFormatter.ID = content.ID
	createContentFormatter.Data = content.Data
	createContentFormatter.IsPublished = content.IsPublished
	createContentFormatter.ContentTypeId = content.ContentTypeId
	createContentFormatter.CreatedByName = content.CreatedByName
	createContentFormatter.CreatedBy = content.CreatedBy

	return createContentFormatter
}

func FormatterGetContentByContentTypeIdResponse(content Content) GetContentByContentTypeIdResponseFormatter {
	getContentFormatter := GetContentByContentTypeIdResponseFormatter{}
	getContentFormatter.ID = content.ID
	getContentFormatter.Data = content.Data
	getContentFormatter.IsPublished = content.IsPublished
	getContentFormatter.ContentTypeId = content.ContentTypeId

	getContentFormatter.CreatedByName = content.CreatedByName
	getContentFormatter.DeletedByName = content.DeletedByName
	getContentFormatter.UpdatedByName = content.UpdatedByName

	getContentFormatter.CreatedBy = content.CreatedBy
	getContentFormatter.DeletedBy = content.DeletedBy
	getContentFormatter.UpdatedBy = content.UpdatedBy

	return getContentFormatter
}


func FormatGetContentsByContentTypeIdResponse(contents []Content) []GetContentByContentTypeIdResponseFormatter {
	var result []GetContentByContentTypeIdResponseFormatter
	for _, res := range contents {
		result = append(result, FormatterGetContentByContentTypeIdResponse(res))
	}
	return result
}

func FormatGetContentResponse(content Content) GetContentByIdResponseFormatter {
	getContentFormatter := GetContentByIdResponseFormatter{}
	getContentFormatter.ID = content.ID
	getContentFormatter.Data = content.Data
	getContentFormatter.IsPublished = content.IsPublished
	getContentFormatter.ContentTypeId = content.ContentTypeId

	getContentFormatter.CreatedByName = content.CreatedByName
	getContentFormatter.DeletedByName = content.DeletedByName
	getContentFormatter.UpdatedByName = content.UpdatedByName

	getContentFormatter.CreatedBy = content.CreatedBy
	getContentFormatter.DeletedBy = content.DeletedBy
	getContentFormatter.UpdatedBy = content.UpdatedBy

	return getContentFormatter
}
func FormatUpdateContentResponse(content Content) UpdateContentByIdResponseFormatter {
	updateContentFormatter := UpdateContentByIdResponseFormatter{}
	updateContentFormatter.ID = content.ID
	updateContentFormatter.Data = content.Data
	updateContentFormatter.IsPublished = content.IsPublished
	updateContentFormatter.ContentTypeId = content.ContentTypeId

	updateContentFormatter.CreatedByName = content.CreatedByName
	updateContentFormatter.DeletedByName = content.DeletedByName
	updateContentFormatter.UpdatedByName = content.UpdatedByName

	updateContentFormatter.CreatedBy = content.CreatedBy
	updateContentFormatter.DeletedBy = content.DeletedBy
	updateContentFormatter.UpdatedBy = content.UpdatedBy

	return updateContentFormatter
}
