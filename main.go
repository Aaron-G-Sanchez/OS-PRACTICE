package main

import (
	"fmt"
	"os"
)

func main() {

	// Reads the current working directory.
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Add a subdir named "sample".
	// err = os.Mkdir("sample", 0750)
	// if err != nil && !os.IsExist(err) {
	// 	log.Fatal(err)
	// }

	// TODO: Pass a flag to name the subdir.

	// Reads the files in the dir.
	// This is just test code.
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// Loop through the directory contents
	// and check to see if entry is a directory or not.
	for _, file := range files {
		// fmt.Println(file)
		if file.IsDir() {
			fmt.Println(file)
		}
	}

	// fmt.Println(files)

}
