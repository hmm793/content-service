package content_type

type ContentTypeFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatContentTypeResponse(contentType Content_Type) ContentTypeFormatter {
	contentTypeFormatter := ContentTypeFormatter{}
	contentTypeFormatter.ID = contentType.ID
	contentTypeFormatter.Name = contentType.Name

	return contentTypeFormatter
}
func FormatContentTypesResponse(contentTypes []Content_Type) []ContentTypeFormatter {
	var contentTypesFormatter []ContentTypeFormatter
	for _, contentType := range contentTypes {
		contentTypesFormatter = append(contentTypesFormatter, FormatContentTypeResponse(contentType))
	}
	return contentTypesFormatter
}