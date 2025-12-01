package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/garrettladley/gml"
)

const (
	gitignore = ".gitignore"
	golangci  = ".golangci.yml"
	makefile  = "Makefile"
)

func main() {
	var path string
	flag.StringVar(&path, "path", ".", "path to scaffold in (default \".\"")
	flag.Parse()

	for name, data := range gml.ScaffoldFiles {
		strName := string(name)
		if err := createFile(path, strName, data); err != nil {
			log.Fatalf("failed to create `%s`: %v", filepath.Join(path, strName), err)
		}
	}
}

func createFile(basePath string, name string, data []byte) error {
	path := filepath.Join(basePath, name)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}
