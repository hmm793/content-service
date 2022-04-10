package content_type

import "gorm.io/gorm"

type Repository interface {
	Save(content_type Content_Type) (Content_Type, error)
	FindById(id int) (Content_Type, error)
	FindAll() ([]Content_Type, error)
	Update(content_type Content_Type) (Content_Type, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(content_type Content_Type) (Content_Type, error) {
	err := r.db.Create(&content_type).Error
	if err != nil {
		return content_type, err
	}
	return content_type, nil
}

func (r *repository) FindById(id int) (Content_Type, error) {
	contentTypeData := Content_Type{}
	err := r.db.Where("id = ?", id).Find(&contentTypeData).Error
	if err != nil {
		return contentTypeData, err
	}
	return contentTypeData, nil
}

func (r *repository) FindAll() ([]Content_Type, error) {
	contentTypeDatas := []Content_Type{}
	err := r.db.Find(&contentTypeDatas).Error
	if err != nil {
		return contentTypeDatas, err
	}
	return contentTypeDatas, nil
}


func (r *repository) Update(content_type Content_Type) (Content_Type, error) {
	err := r.db.Save(&content_type).Error
	if err != nil {
		return content_type, err
	}
	return content_type, nil
}