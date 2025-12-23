package objectstorage_template

const RustfsConfigTemplate string = `
package storage

import 	(
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func NewRustfsConnection() *minio.Client {
	hostPort := viper.GetString("rustfs.host") + ":" + viper.GetString("rustfs.port")
	rustfsClient, err := minio.New(hostPort, &minio.Options{
		Creds: credentials.NewStaticV4(
			viper.GetString("rustfs.credential.user"),
			viper.GetString("rustfs.credential.password"),
			"",
		),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return rustfsClient
}
`

const DockerComposeRustfsConfigTemplate string = `
# rustfs
  {{.ProjectName}}_storage:
    image: rustfs/rustfs:1.0.0-alpha.76
    container_name: {{.ProjectName}}_storage
    environment:
      - RUSTFS_CONSOLE_ENABLE=true
      - RUSTFS_ADDRESS=0.0.0.0:9000
      - RUSTFS_CONSOLE_ADDRESS=0.0.0.0:9001
      - RUSTFS_EXTERNAL_ADDRESS=:9000 # Same as internal since no port mapping
      - RUSTFS_CORS_ALLOWED_ORIGINS=*
      - RUSTFS_CONSOLE_CORS_ALLOWED_ORIGINS=*
      - RUSTFS_ACCESS_KEY=${OS_ROOT_USER}
      - RUSTFS_SECRET_KEY=${OS_ROOT_PASSWORD}
    restart: unless-stopped
    ports:
      - "9000:9000"
      - "9001:9001"
    volumes:
      - {{.ProjectName}}_rustfs_data:/data
    healthcheck:
      test:
        [
          "CMD",
          "sh", "-c",
          "curl -f http://localhost:9000/health && curl -f http://localhost:9001/rustfs/console/health"
        ]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
`

const DockerComposeRustfsVolumeTemplate string = `
  {{.ProjectName}}_rustfs_data: {}`
