package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
)

func InitPkgExample(path *string) error {
	if path != nil {
		folderName := "/internal/pkg"
		fileName := folderName + "/.gitkeep"
		if err := pkg.GenericFileGenerator(path, folderName, fileName, ""); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
