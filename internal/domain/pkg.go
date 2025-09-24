package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
)

func InitPkgExample(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/internal/pkg"
		fileName := folderName + "/.gitkeep"
		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, ""); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
