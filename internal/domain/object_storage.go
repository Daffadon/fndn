package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	objectstorage_template "github.com/daffadon/fndn/internal/template/object_storage"
)

func InitObjectStorageConfig(i infra.CommandRunner, path *string, os *string) error {
	if path != nil {
		folderName := "/config/storage"
		fileName := folderName + "/storage.go"
		var template string
		switch *os {
		case "rustfs":
			template = objectstorage_template.RustfsConfigTemplate
		case "seaweedfs":
			template = objectstorage_template.SeaweedfsConfigTemplate
		case "minio":
			template = objectstorage_template.MinioConfigTemplate
		}
		if template != "" {
			if err := pkg.GoFileGenerator(i, path, folderName, fileName, template); err != nil {
				log.Fatal(err)
				return err
			}
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitObjectStorageConfigFile(i infra.CommandRunner, p *Project) error {
	if p.Path != nil {
		folderName := "/config/storage"
		fileName := folderName
		var template string

		switch p.ObjectStorage {
		case "seaweedfs":
			fileName += "/s3.json"
			template = objectstorage_template.SeaweedfsConfigFileTemplate
		}

		if template != "" {
			if err := pkg.GenericFileGenerator(i, p.Path, folderName, fileName, template); err != nil {
				log.Fatal(err)
				return err
			}
		}
		return nil
	}
	return errors.New("path is nil")
}

func GenerateSpecificStorage(storage string, infraRunner infra.CommandRunner, path string) error {
	// check folder config/router/ exist or not
	// check filename
	folderName := "/config/storage_"
	fileName := folderName + "/storage.go"

	// if exist, the file name add _framework_name
	exist := pkg.IsFileExists("." + fileName)
	if exist {
		fileName = folderName + "/storage_" + storage + ".go"
	}

	var t string
	switch storage {
	case "rustfs":
		t = objectstorage_template.RustfsConfigTemplate
	case "seaweedfs":
		t = objectstorage_template.SeaweedfsConfigTemplate
	case "minio":
		t = objectstorage_template.MinioConfigTemplate
	}

	if err := pkg.GoFileGenerator(infraRunner, &path, folderName, fileName, t); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
