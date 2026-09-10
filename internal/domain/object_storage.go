package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
	objectstorage_template "github.com/daffadon/fndn/internal/template/object_storage"
)

var objectStorageTemplates = map[string]string{
	"rustfs":    objectstorage_template.RustfsConfigTemplate,
	"seaweedfs": objectstorage_template.SeaweedfsConfigTemplate,
	"minio":     objectstorage_template.MinioConfigTemplate,
}

func InitObjectStorageConfig(path *string, os *string) error {
	if *os == None {
		return nil
	}
	if path != nil {
		folderName := "/config/storage"
		fileName := folderName + "/storage.go"
		template, err := lookup(objectStorageTemplates, "object storage", *os)
		if err != nil {
			return err
		}
		if err := pkg.GoFileGenerator(path, folderName, fileName, template); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitObjectStorageConfigFile(p *Project) error {
	if p.ObjectStorage == None {
		return nil
	}
	if p.Path != nil {
		if p.ObjectStorage != "seaweedfs" {
			return nil
		}
		folderName := "/config/storage"
		fileName := folderName + "/s3.json"
		if err := pkg.GenericFileGenerator(p.Path, folderName, fileName, objectstorage_template.SeaweedfsConfigFileTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func GenerateSpecificStorage(storage string, path string) error {
	folderName := "/config/storage_"
	fileName := resolveTarget(folderName+"/storage.go", folderName+"/storage", storage)

	t, err := lookup(objectStorageTemplates, "object storage", storage)
	if err != nil {
		return err
	}
	if err := pkg.GoFileGenerator(&path, folderName, fileName, t); err != nil {
		return err
	}
	return nil
}
