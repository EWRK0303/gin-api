package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Connect() *minio.Core {
	core, _ := minio.NewCore("127.0.0.1:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("admin", "admin12345", ""),
		Secure: false,
	})
	return core
}
