package content_type

type Content_Type_Input struct {
	Name string `json:"name" binding:"required"`
}

type GetContentTypeInput struct {
	ID int `uri:"id" binding:"required"`
}