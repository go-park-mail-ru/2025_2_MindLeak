package minio

import (
	"fmt"
	"os"
)

type MinioConfig struct {
	User      string
	Password  string
	Endpoint  string
	Bucket    string
	PublicURL string
}

func NewMinioConfig() *MinioConfig {
	return &MinioConfig{
		User:      os.Getenv("MINIO_ROOT_USER"),
		Password:  os.Getenv("MINIO_ROOT_PASSWORD"),
		Endpoint:  os.Getenv("MINIO_ENDPOINT"),
		Bucket:    os.Getenv("MINIO_BUCKET"),
		PublicURL: os.Getenv("MINIO_PUBLIC_ENDPOINT"),
	}
}

func (c *MinioConfig) GetDefaultAvatar() string {
	return fmt.Sprintf("%s/%s/defaultAvatar.jpg", c.PublicURL, c.Bucket)
}

func (c *MinioConfig) GetDefaultCover() string {
	return fmt.Sprintf("%s/%s/cover-pic.jpg", c.PublicURL, c.Bucket)
}
