package config

import (
	"let-you-cook/utils/helper"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioUseSSL    bool
)

func InitMinio() (*minio.Client, error) {
	MinioEndpoint = helper.GetENV("MINIO_URI", "127.0.0.1:9000")
	MinioAccessKey = helper.GetENV("MINIO_ACCESS", "minio")
	MinioSecretKey = helper.GetENV("MINIO_SECRET", "minio123")
	MinioUseSSL = false

	minioClient, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKey, MinioSecretKey, ""),
		Secure: MinioUseSSL,
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}