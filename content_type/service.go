package content_type

import "errors"

type Service interface {
	Create(input Content_Type_Input) (Content_Type, error)
	GetById(id int) (Content_Type, error)
	GetAll() ([]Content_Type, error)
	Update(id int, input Content_Type_Input) (Content_Type, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input Content_Type_Input) (Content_Type, error) {
	contentType := Content_Type{}
	contentType.Name = input.Name

	NewContentType, err := s.repository.Save(contentType)

	if err != nil {
		return NewContentType, err
	}

	return NewContentType, nil
}

func (s *service) GetById(id int) (Content_Type, error) {
	contentTypeData, err := s.repository.FindById(id)

	if err != nil {
		return contentTypeData, err
	}

	if contentTypeData.ID == 0 {
		return contentTypeData, errors.New("Content Type Not Found")
	}
	
	return contentTypeData, nil
}

func (s *service) GetAll() ([]Content_Type, error) {
	contentTypeDatas, err :=  s.repository.FindAll()

	if err != nil {
		return contentTypeDatas, err
	}

	if len(contentTypeDatas) == 0 {
		return contentTypeDatas, errors.New("There is no data")
	}

	return contentTypeDatas, nil

}

func (s *service) Update(id int, input Content_Type_Input) (Content_Type, error) {
	contentTypeData, err := s.repository.FindById(id)

	if err != nil {
		return contentTypeData, err
	}
	if contentTypeData.ID == 0 {
		return contentTypeData, errors.New("Content Type Not Found")
	}

	contentTypeData.Name = input.Name

	contentTypeDataUpdated, err := s.repository.Update(contentTypeData)
	
	if err != nil {
		return contentTypeData, err
	}

	return contentTypeDataUpdated, nil

}