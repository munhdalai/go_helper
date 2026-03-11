package helpers_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleFileExists() {
	f, _ := os.CreateTemp("", "test")
	defer os.Remove(f.Name())
	f.Close()
	fmt.Println(helpers.FileExists(f.Name()))
	fmt.Println(helpers.FileExists("/nonexistent/file.txt"))
	// Output:
	// true
	// false
}

func ExampleWriteJSON() {
	f, _ := os.CreateTemp("", "test*.json")
	f.Close()
	defer os.Remove(f.Name())
	data := map[string]string{"name": "test"}
	err := helpers.WriteJSON(f.Name(), data)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleReadJSON() {
	f, _ := os.CreateTemp("", "test*.json")
	defer os.Remove(f.Name())
	f.WriteString(`{"name":"test","age":25}`)
	f.Close()
	result, err := helpers.ReadJSON[map[string]interface{}](f.Name())
	fmt.Println(result["name"], err)
	// Output: test <nil>
}

func ExampleFileExtension() {
	fmt.Println(helpers.FileExtension("document.pdf"))
	fmt.Println(helpers.FileExtension("archive.tar.gz"))
	fmt.Println(helpers.FileExtension("noext"))
	// Output:
	// pdf
	// gz
	//
}

func ExampleFileSize() {
	f, _ := os.CreateTemp("", "test")
	f.WriteString("hello")
	f.Close()
	defer os.Remove(f.Name())
	size, _ := helpers.FileSize(f.Name())
	fmt.Println(size)
	// Output: 5
}

func ExampleEnsureDir() {
	dir := filepath.Join(os.TempDir(), "test_ensure_dir")
	defer os.RemoveAll(dir)
	err := helpers.EnsureDir(dir)
	fmt.Println(err)
	fmt.Println(helpers.FileExists(dir))
	// Output:
	// <nil>
	// true
}
