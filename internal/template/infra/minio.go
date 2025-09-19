package infra_template

const MinioInfraTemplate string = `
package storage_infra

type (
	MinioInfra interface {
		Set(ctx context.Context, bucketName, objectPath string, reader io.Reader, objectSize int64) error
		Get(ctx context.Context, bucketName, objectPath string) (io.ReadCloser, error)
		Delete(ctx context.Context, bucketName, fileName string) error
		CreateBucketIfNotExist(ctx context.Context, bucketName string) error
		SetPolicy(ctx context.Context, bucketName, policy string) error
		GetPolicy(ctx context.Context, bucketName string) (string, error)
	}
	minioInfra struct {
		client *minio.Client
	}
)

func NewMinioInfra(client *minio.Client) MinioInfra {
	return &minioInfra{
		client: client,
	}
}

func (m *minioInfra) Set(ctx context.Context, bucketName, objectPath string, reader io.Reader, objectSize int64) error {
	_, err := m.client.PutObject(ctx, bucketName, objectPath, reader, objectSize, minio.PutObjectOptions{})
	return err
}

func (m *minioInfra) Get(ctx context.Context, bucketName, objectPath string) (io.ReadCloser, error) {
	object, err := m.client.GetObject(ctx, bucketName, objectPath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (m *minioInfra) Delete(ctx context.Context, bucketName, fileName string) error {
	return m.client.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
}

func (m *minioInfra) CreateBucketIfNotExist(ctx context.Context, bucketName string) error {
	exists, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		return m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}
	return nil
}
func (m *minioInfra) SetPolicy(ctx context.Context, bucketName, policy string) error {
	return m.client.SetBucketPolicy(ctx, bucketName, policy)
}

func (m *minioInfra) GetPolicy(ctx context.Context, bucketName string) (string, error) {
	return m.client.GetBucketPolicy(ctx, bucketName)
}
`
