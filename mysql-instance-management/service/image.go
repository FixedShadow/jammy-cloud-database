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
	return ""
}
