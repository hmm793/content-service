package content

type CreateContentInput struct {
	Data          string `json:"data" binding:"required"`
	IsPublished   bool   `json:"is_published" binding:"required"`
	ContentTypeId int    `json:"content_type_id" binding:"required"`
	CreatedByName string `json:"created_by_name" binding:"required"`
	CreatedBy     int    `json:"created_by" binding:"required"`
}

type GetContentByContentTypeIdInput struct {
	ID int `uri:"contentTypeId" binding:"required"`
}
type GetContentIdInput struct {
	ID int `uri:"id" binding:"required"`
}

type Content_Input struct {
	Data          string `json:"data" binding:"required"`
	IsPublished   bool   `json:"is_published" binding:"required"`
	ContentTypeId int    `json:"content_type_id" binding:"required"`
	UpdatedByName string `json:"updated_by_name"`
	UpdatedBy     int    `json:"updated_by"`
}

// type CreateContentInput struct {
// 	Data          string `json:"data" binding:"required"`
// 	IsPublished   bool   `json:"is_published" binding:"required"`
// 	ContentTypeId int    `json:"content_type_id" binding:"required"`
// 	CreatedByName string `json:"created_by_name"`
// 	DeletedByName string `json:"deleted_by_name"`
// 	UpdatedByName string `json:"updated_by_name"`
// 	CreatedBy     int    `json:"created_by"`
// 	UpdatedBy     int    `json:"updated_by"`
// 	DeletedBy     int    `json:"deleted_by"`
// }