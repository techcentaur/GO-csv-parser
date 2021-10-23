package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadCSVLocally(csvUrl, localFileName string) (int, error) {
	response, err := http.Get(csvUrl)
	if err != nil {
		return -1, err
	}
	defer response.Body.Close()

	file, err := os.Create(localFileName)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	bytesWritten, err := io.Copy(file, response.Body)
	if err != nil {
		return -1, err
	}
	return int(bytesWritten), nil
}

func main() {
	url := os.Args[1]
	fileName := os.Args[2]

	// url = "https://people.sc.fsu.edu/~jburkardt/data/csv/cities.csv"
	// filename = "test1.csv"

	bytesWritten, err := downloadCSVLocally(url, fileName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Fetching content for the URL: %q\n", url)
	fmt.Printf("%d Bytes written in file: %q\n", bytesWritten, fileName)
}
