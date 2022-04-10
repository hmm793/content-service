package content

import "errors"

type Service interface {
	Create(input CreateContentInput) (Content, error)
	FindByContentTypeId(id int) ([]Content, error)
	FindById(id int) (Content, error)
	UpdateContentById(id int, input Content_Input) (Content, error)
}


type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input CreateContentInput) (Content, error) {
	contentData := Content{}
	contentData.Data = input.Data
	contentData.IsPublished = input.IsPublished
	contentData.CreatedByName = input.CreatedByName
	contentData.ContentTypeId = input.ContentTypeId
	contentData.CreatedBy = input.CreatedBy

	newContentData, err := s.repository.Save(contentData)

	if err != nil {
		return newContentData, err
	}
	return newContentData, nil
}

func (s *service) FindByContentTypeId(id int) ([]Content, error) {
	contents, err := s.repository.FindByContentTypeId(id)
	if err != nil {
		return contents, err
	}

	if len(contents) == 0 {
		return contents, errors.New("Content Not Found")
	}

	return contents, nil

}


func (s *service) FindById(id int) (Content, error) {
	content, err := s.repository.FindById(id)
	if err != nil {
		return content, err
	}
	if content.ID == 0 {
		return content, errors.New("Content Not Found")
	}
	return content, nil
}

func (s *service) UpdateContentById(id int, input Content_Input) (Content, error) {
	contentData, err := s.repository.FindById(id)

	if err != nil {
		return contentData, err
	}

	if contentData.ID == 0 {
		return contentData, errors.New("Content Not Found")
	}

	contentData.Data = input.Data
	contentData.IsPublished = input.IsPublished
	contentData.ContentTypeId = input.ContentTypeId
	contentData.UpdatedByName = input.UpdatedByName
	contentData.UpdatedBy = input.UpdatedBy

	contentDataUpdated, err := s.repository.Update(contentData)
	
	if err != nil {
		return contentData, err
	}

	return contentDataUpdated, nil
}