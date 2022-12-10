package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.Name() != "main.go" && f.Name() != "go.mod" {
			start(f.Name())
		}
	}
}

func start(filename string) {
	removeLine(filename, 0)
}

func removeLine(path string, line int) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	mode := info.Mode()
	array := strings.Split(string(file), "\n")
	array = append(array[:line], array[line+1:]...)
	ioutil.WriteFile(path, []byte(strings.Join(array, "\n")), mode)
}
