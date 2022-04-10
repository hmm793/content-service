package content

import (
	"content-service/content_type"
	"time"
)

type Content struct {
	ID            int
	Data          string
	IsPublished   bool
	CreatedByName string
	DeletedByName string
	UpdatedByName string
	CreatedBy     int
	UpdatedBy     int
	DeletedBy     int
	ContentTypeId int
	CreatedAt 	  time.Time
	UpdatedAt 	  time.Time
	DeletedAt 	  time.Time
	ContentType   content_type.Content_Type
}
