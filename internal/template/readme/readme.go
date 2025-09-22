package readme

import (
	"io"
	"net/http"
)

func CopyReadmeTemplate() (string, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/Daffadon/fndn/refs/heads/main/internal/template/readme/readme.md")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
