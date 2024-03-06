package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {
	// Create a zip file to write our archive to.
	file, err := os.Create("xxx.zip")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new zip archive.
	w := zip.NewWriter(file)

	// Add some files to the archive.
	files := []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
