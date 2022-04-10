package content

import "gorm.io/gorm"

type Repository interface {
	Save(content Content) (Content, error)
	FindByContentTypeId(id int) ([]Content, error)
	FindById(id int) (Content, error)
	Update(content Content) (Content, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(content Content) (Content, error) {
	err := r.db.Create(&content).Error
	if err != nil {
		return content, err
	}
	return content, nil
}


func (r *repository) FindByContentTypeId(id int) ([]Content, error) {
	var contents []Content
	err := r.db.Where("content_type_id = ?", id).Find(&contents).Error
	if err != nil {
		return contents, err
	}
	return contents, nil
}

func (r *repository) FindById(id int) (Content, error) {
	var content Content
	err := r.db.Where("id = ?", id).Find(&content).Error

	if err != nil {
		return content, err
	}

	return content, nil
} 


func (r *repository) Update(content Content) (Content, error) {
	err := r.db.Save(&content).Error
	if err != nil {
		return content, err
	}
	return content, nil
}