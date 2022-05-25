package s3

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/monferon/fsm/loader/config"
)

type s3client struct {
	Client *minio.Client
}

func New(cfg config.S3) (*s3client, error) {
	minioClient, err := minio.New(cfg.URL, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	return &s3client{Client: minioClient}, nil
}
