package minio

import (
	"fmt"
	"os"
)

//const (
//	DefaultAvatarURL = "https://62.109.19.84:9000/mindleak-bucket/defaultAvatar.jpg"
//	DefaultCoverURL  = "https://62.109.19.84:9000/mindleak-bucket/cover-pic.jpg"
//)

type MinioConfig struct {
	MinioUser           string
	MinioPass           string
	MinioEndpoint       string
	MinioBucket         string
	MinioPublicEndpoint string
}

func NewMinioConfig() *MinioConfig {
	return &MinioConfig{
		MinioUser:           os.Getenv("MINIO_ROOT_USER"),
		MinioPass:           os.Getenv("MINIO_ROOT_PASSWORD"),
		MinioEndpoint:       os.Getenv("MINIO_ENDPOINT"),
		MinioBucket:         os.Getenv("MINIO_BUCKET"),
		MinioPublicEndpoint: os.Getenv("MINIO_PUBLIC_ENDPOINT"),
	}
}

func (c *MinioConfig) GetDefaultAvatar() string {
	return fmt.Sprintf("%s/%s/defaultAvatar.jpg", c.MinioPublicEndpoint, c.MinioBucket)
}

func (c *MinioConfig) GetDefaultCover() string {
	return fmt.Sprintf("%s/%s/cover-pic.jpg", c.MinioPublicEndpoint, c.MinioBucket)
}
