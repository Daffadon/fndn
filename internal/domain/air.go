package domain

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/pelletier/go-toml"
)

func InitAirConfig(i infra.CommandRunner, path *string, initAir bool) error {
	if initAir {
		_, err := exec.LookPath("air")
		if err != nil {
			log.Default().Println("air is not installed. Please run: go install github.com/air-verse/air@latest")
			return nil
		}
		if err := i.Run("air", []string{"init"}, *path); err != nil {
			return err
		}

		configPath := *path + "/.air.toml"

		tree, err := toml.LoadFile(configPath)
		if err != nil {
			return err
		}

		goos := os.Getenv("GOOS")
		if goos == "" {
			goos = runtime.GOOS
		}

		var buildCmd string
		if goos == "windows" {
			buildCmd = "go build -o ./tmp/main.exe ./cmd"
		} else {
			buildCmd = "go build -o ./tmp/main ./cmd"
		}

		tree.Set("build.cmd", buildCmd)

		tomlStr, err := tree.ToTomlString()
		if err != nil {
			return err
		}
		if err := os.WriteFile(configPath, []byte(tomlStr), 0644); err != nil {
			return err
		}
	}
	return nil
}
