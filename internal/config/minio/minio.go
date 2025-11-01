package minio

import (
	"os"
)

const (
	DefaultAvatarURL = "http://62.109.19.84:9000/mindleak-bucket/defaultAvatar.jpg"
	DefaultCoverURL  = "http://62.109.19.84:9000/mindleak-bucket/cover-pic.jpg"
)

type MinioConfig struct {
	MinioUser     string
	MinioPass     string
	MinioEndpoint string
	MinioBucket   string
}

func NewMinioConfig() *MinioConfig {
	return &MinioConfig{
		MinioUser:     os.Getenv("MINIO_ROOT_USER"),
		MinioPass:     os.Getenv("MINIO_ROOT_PASSWORD"),
		MinioEndpoint: os.Getenv("MINIO_ENDPOINT"),
		MinioBucket:   os.Getenv("MINIO_BUCKET"),
	}
}
