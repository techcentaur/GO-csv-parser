package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func download_csv_locally(csv_url, local_file_name string) int {
	response, err := http.Get(csv_url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	file, err := os.Create(local_file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data_len, err := io.Copy(file, response.Body)
	if err != nil {
		panic(err)
	}
	return int(data_len)
}

func main() {
	var url, filename string

	url = "https://people.sc.fsu.edu/~jburkardt/data/csv/cities.csv"
	filename = "test1.csv"

	data_len := download_csv_locally(url, filename)

	fmt.Printf("Fetching content for the URL: %q\n", url)
	fmt.Printf("%d Bytes written in file: %q\n", data_len, filename)
}
