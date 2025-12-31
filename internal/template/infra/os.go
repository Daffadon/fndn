package infra_template

const MinioInfraTemplate string = `
package storage_infra

import "github.com/minio/minio-go/v7"

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
const RustfsInfraTemplate string = `
package storage_infra

import "github.com/minio/minio-go/v7"

type (
	RustfsInfra interface {
		Set(ctx context.Context, bucketName, objectPath string, reader io.Reader, objectSize int64) error
		Get(ctx context.Context, bucketName, objectPath string) (io.ReadCloser, error)
		Delete(ctx context.Context, bucketName, fileName string) error
		CreateBucketIfNotExist(ctx context.Context, bucketName string) error
		SetPolicy(ctx context.Context, bucketName, policy string) error
		GetPolicy(ctx context.Context, bucketName string) (string, error)
	}
	rustfsInfra struct {
		client *minio.Client
	}
)

func NewRustfsInfra(client *minio.Client) RustfsInfra {
	return &rustfsInfra{
		client: client,
	}
}

func (r *rustfsInfra) Set(ctx context.Context, bucketName, objectPath string, reader io.Reader, objectSize int64) error {
	_, err := r.client.PutObject(ctx, bucketName, objectPath, reader, objectSize, minio.PutObjectOptions{})
	return err
}

func (r *rustfsInfra) Get(ctx context.Context, bucketName, objectPath string) (io.ReadCloser, error) {
	object, err := r.client.GetObject(ctx, bucketName, objectPath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (r *rustfsInfra) Delete(ctx context.Context, bucketName, fileName string) error {
	return r.client.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
}

func (r *rustfsInfra) CreateBucketIfNotExist(ctx context.Context, bucketName string) error {
	exists, err := r.client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		return r.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}
	return nil
}
func (r *rustfsInfra) SetPolicy(ctx context.Context, bucketName, policy string) error {
	return r.client.SetBucketPolicy(ctx, bucketName, policy)
}

func (r *rustfsInfra) GetPolicy(ctx context.Context, bucketName string) (string, error) {
	return r.client.GetBucketPolicy(ctx, bucketName)
}
`

const SeaweedInfraTemplate string = `
package storage_infra

import "github.com/minio/minio-go/v7"

type (
	SeaweedfsInfra interface {
		Set(ctx context.Context, bucketName, objectPath string, reader io.Reader, objectSize int64) error
		Get(ctx context.Context, bucketName, objectPath string) (io.ReadCloser, error)
		Delete(ctx context.Context, bucketName, fileName string) error
		CreateBucketIfNotExist(ctx context.Context, bucketName string) error
		SetPolicy(ctx context.Context, bucketName, policy string) error
		GetPolicy(ctx context.Context, bucketName string) (string, error)
	}
	seaweedfsInfra struct {
		client *minio.Client
	}
)

func NewSeaweedfsInfra(client *minio.Client) SeaweedfsInfra {
	return &seaweedfsInfra{
		client: client,
	}
}

func (s *seaweedfsInfra) Set(ctx context.Context, bucketName, objectPath string, reader io.Reader, objectSize int64) error {
	_, err := s.client.PutObject(ctx, bucketName, objectPath, reader, objectSize, minio.PutObjectOptions{})
	return err
}

func (s *seaweedfsInfra) Get(ctx context.Context, bucketName, objectPath string) (io.ReadCloser, error) {
	object, err := s.client.GetObject(ctx, bucketName, objectPath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (s *seaweedfsInfra) Delete(ctx context.Context, bucketName, fileName string) error {
	return s.client.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
}

func (s *seaweedfsInfra) CreateBucketIfNotExist(ctx context.Context, bucketName string) error {
	exists, err := s.client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		return s.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}
	return nil
}
func (s *seaweedfsInfra) SetPolicy(ctx context.Context, bucketName, policy string) error {
	return s.client.SetBucketPolicy(ctx, bucketName, policy)
}

func (s *seaweedfsInfra) GetPolicy(ctx context.Context, bucketName string) (string, error) {
	return s.client.GetBucketPolicy(ctx, bucketName)
}
`
