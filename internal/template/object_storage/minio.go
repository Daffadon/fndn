package objectstorage_template

const MinioConfigTemplate string = `
package storage

import 	(
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
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

const DockerComposeMinioConfigTemplate string = `
# minio
  {{.ProjectName}}_storage:
    image: minio/minio:RELEASE.2025-05-24T17-08-30Z-cpuv1
    container_name: {{.ProjectName}}_storage
    environment:
      MINIO_ROOT_USER: ${MINIO_ROOT_USER}
      MINIO_ROOT_PASSWORD: ${MINIO_ROOT_PASSWORD}
    restart: unless-stopped
    ports:
      - "9000:80"
      - "9001:9001"
    volumes:
      - {{.ProjectName}}_minio_data:/data
    command: server /data  --address ":80" --console-address ":9001"
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:80/minio/health/live"]
      interval: 30s
      timeout: 10s
      retries: 5
`

const DockerComposeMinioVolumeTemplate string = `
  {{.ProjectName}}_minio_data: {}`
