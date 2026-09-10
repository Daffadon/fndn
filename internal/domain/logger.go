package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
	logger_template "github.com/daffadon/fndn/internal/template/logger"
)

func InitZerologConfig(path *string) error {
	if path != nil {
		folderName := "/config/logger"
		fileName := folderName + "/zerolog.go"
		if err := pkg.GoFileGenerator(path, folderName, fileName, logger_template.ZerologConfigTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
