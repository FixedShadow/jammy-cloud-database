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
	return "67c922efe9030a23c57dd86eccd0dfef03e54565126e79f88bb8cd9b34ecd992"
}
