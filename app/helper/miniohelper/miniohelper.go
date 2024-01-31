package miniohelper

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

type Minio struct {
	Client     *minio.Client
	BucketName string
}

func GetMinio(config *viper.Viper, logger *logger.Logger) *Minio {
	endpoint := config.GetString("minio.endpoint")
	accessKeyID := config.GetString("minio.accessKeyId")
	secretAccessKey := config.GetString("minio.secretAccessKey")
	useSSL := config.GetBool("minio.useSSL")
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logger.Info("Minio client initialized failed")
		logger.Fatal(err.Error())
	}
	logger.Info("Minio client initialized successfully")

	return &Minio{
		Client:     minioClient,
		BucketName: config.GetString("minio.bucketName"),
	}
}
