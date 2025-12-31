package objectstorage_template

const SeaweedfsConfigTemplate string = `
package storage

import 	(
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func NewSeaweedfsConnection() *minio.Client {
	hostPort := viper.GetString("seaweedfs.host") + ":" + viper.GetString("seaweedfs.port")
	seaweedfsClient, err := minio.New(hostPort, &minio.Options{
		Creds: credentials.NewStaticV4(
			viper.GetString("seaweedfs.credential.user"),
			viper.GetString("seaweedfs.credential.password"),
			"",
		),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return seaweedfsClient
}
`

const DockerComposeSeaweedfsConfigTemplate string = `
# seaweedfs master
  {{.ProjectName}}_storage_master:
    image: chrislusf/seaweedfs:4.03
    container_name: {{.ProjectName}}_storage_master
		command: 'master -ip={{.ProjectName}}_storage_master -port=9333 -disableHttp -volumePreallocate=false'
		healthcheck:
			test: wget http://{{.ProjectName}}_storage_master:9333/cluster/status -O -
			interval: 10s
			timeout: 5s
			retries: 999
		restart: unless-stopped
		volumes:
			- {{.ProjectName}}_seaweed_master_data:/data

# seaweed volume
  {{.ProjectName}}_storage_volume:
    image: chrislusf/seaweedfs:4.03
    container_name: {{.ProjectName}}_storage_volume
		command: 'volume -ip={{.ProjectName}}_storage_volume -port=8080 -mserver={{.ProjectName}}_storage_master:9333'
		healthcheck:
			test: wget http://{{.ProjectName}}_storage_volume:8080/status -O -
			interval: 10s
			timeout: 5s
			retries: 999
		depends_on:
			{{.ProjectName}}_storage_master:
				condition: service_healthy
    restart: unless-stopped
		volumes:
			- {{.ProjectName}}_seaweed_volume_data:/data

# seaweed filer
	{{.ProjectName}}_storage_filer:
		image: chrislusf/seaweedfs:4.03
		container_name: {{.ProjectName}}_storage_filer
		command: 'filer -master={{.ProjectName}}_storage_master:9333 -ip={{.ProjectName}}_storage_filer -ip.bind=0.0.0.0 -port=8888 -port.readonly=8889'
		ports:
			- '8889:8889'
			- '8888:8888'
		healthcheck:
			test: wget http://{{.ProjectName}}_storage_filer:8888/ -O -
			interval: 10s
			timeout: 5s
			retries: 999
		depends_on:
			{{.ProjectName}}_storage_master:
				condition: service_healthy
			{{.ProjectName}}_storage_volume:
				condition: service_healthy
		restart: unless-stopped
		volumes:
				- {{.ProjectName}}_seaweed_filer_data:/data

# seaweed s3
	{{.ProjectName}}_storage_s3:
		image: chrislusf/seaweedfs:4.03
		container_name: {{.ProjectName}}_storage_s3
		command: 's3 -filer={{.ProjectName}}_storage_filer:8888 -port=9000 -ip.bind=0.0.0.0 -config=/etc/seaweedfs/s3.json'
		healthcheck:
			test: wget http://{{.ProjectName}}_storage_s3:9000/ -O -
			interval: 10s
			timeout: 5s
			retries: 999
		ports:
			- '9000:9000'
		depends_on:
			{{.ProjectName}}_storage_filer:
				condition: service_healthy
		restart: unless-stopped
		volumes:
			- ./config/storage/s3.json:/etc/seaweedfs/s3.json:ro
`

const DockerComposeSeaweedfsVolumeTemplate string = `
  {{.ProjectName}}_seaweed_master_data: {}
	{{.ProjectName}}_seaweed_volume_data: {}
	{{.ProjectName}}_seaweed_filer_data: {}
`

const SeaweedfsConfigFileTemplate string = `
{
  "identities": [
    {
      "name": "superclient",
      "credentials": [
        {
          "accessKey": "ROOTUSER",
          "secretKey": "CHANGEME123"
        }
      ],
      "actions": [
        "Read",
        "Write",
        "List",
        "Tagging",
        "Admin"
      ]
    }
  ]
}
`
