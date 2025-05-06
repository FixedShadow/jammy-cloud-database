package service

type ImageService struct{}

type IImageService interface {
	GetImageIdByType(imageType string) string
}

func NewImageService() IImageService {
	return &ImageService{}
}

// TODO get image id from db.
func (i *ImageService) GetImageIdByType(imageType string) string {
	return "913ae47e658993fb7d5e89995ffe91b2194c14843aaa9032e69ec8737055db75"
}
