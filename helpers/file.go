package helpers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// FileExists нь файл байгаа эсэхийг шалгана.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ReadJSON нь JSON файлыг уншиж struct-руу decode хийнэ.
func ReadJSON[T any](path string) (T, error) {
	var result T
	data, err := os.ReadFile(path)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// WriteJSON нь утгыг JSON файл руу бичнэ.
func WriteJSON(path string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// FileExtension нь файлын өргөтгөлийг буцаана (цэггүй).
func FileExtension(path string) string {
	ext := filepath.Ext(path)
	return strings.TrimPrefix(ext, ".")
}

// FileSize нь файлын хэмжээг байтаар буцаана.
func FileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// EnsureDir нь хавтас байхгүй бол үүсгэнэ.
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}
