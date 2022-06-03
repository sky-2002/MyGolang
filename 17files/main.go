package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	// we use defer to close after use
	// we use ioutils for io operations

	content := "This is some content that needs to be put in some file."

	file, err := os.Create("./mygofile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("The length is", length)

	defer file.Close()

	readFile("./mygofile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename) // data is read in the form of bytes

	checkNilError(err)

	fmt.Println("Text data in bytes:\n", dataByte)
	fmt.Println("Text data in text:\n", string(dataByte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
