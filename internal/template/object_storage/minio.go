package objectstorage_template

const MinioConfigTemplate string = `
package storage

import 	(
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioConnection() *minio.Client {
	hostPort := viper.GetString("minio.host") + ":" + viper.GetString("minio.port")
	minioClient, err := minio.New(hostPort, &minio.Options{
		Creds: credentials.NewStaticV4(
			viper.GetString("minio.credential.user"),
			viper.GetString("minio.credential.password"),
			"",
		),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}
`
