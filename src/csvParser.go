package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// CSV data would be stored in these data structure
type Dataframe struct {
	columns []Column
}
type Column struct {
	elements       Elements
	name, dataType string
}
type Elements []interface{}

// FreqList has sorting structure defined on it using `Value` of Freq
type Freq struct {
	Key   interface{}
	Value int
}

func (f *Freq) addValue(i int) {
	f.Value += i
}

type FreqList []Freq

func (f FreqList) Len() int           { return len(f) }
func (f FreqList) Less(i, j int) bool { return f[i].Value < f[j].Value }
func (f FreqList) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func main() {
	csvFileName := os.Args[1]
	outputFileName := os.Args[2]

	data, err := readUsingCSVModule(csvFileName)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Reading from file: %v\n\n", csvFileName)
	}

	df := new(Dataframe)

	// using first row for headers and second row to determine types
	for i, col := range data[0] {
		first_value := strings.TrimSpace(data[1][i])

		_type := reflect.TypeOf(first_value).String()

		if _, err := strconv.Atoi(first_value); err == nil {
			_type = "int"
		} else if _, err := strconv.ParseFloat(first_value, 64); err == nil {
			_type = "float"
		} else if _, err := strconv.ParseBool(first_value); err == nil {
			_type = "bool"
		}
		df.columns = append(df.columns, Column{name: col, dataType: _type})
	}

	for _, col := range data[1:] {
		for j, val := range col {
			df.columns[j].elements = append(df.columns[j].elements, strings.TrimSpace(val))
		}
	}

	writeData := ""
	for _, col := range df.columns {
		freq := col.getCounter()
		fmt.Printf("Column: %v\n\tMost freq item: %v\n\tfreq count: %v\n", col.name, freq[0].Key, freq[0].Value)
		writeData += fmt.Sprintf("%v,%v,%v\n", col.name, freq[0].Key, freq[0].Value)
	}

	err = writeToFile(outputFileName, writeData)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("\nData written to file: %v\n", outputFileName)
	}
}

func (column *Column) getCounter() FreqList {
	counter := make(map[interface{}]*Freq)

	for _, val := range column.elements {
		if _, ok := counter[val]; ok {
			counter[val].addValue(1)
		} else if val != nil {
			counter[val] = &Freq{Key: val, Value: 1}
		}
	}

	counterList := make(FreqList, 0)
	for _, val := range counter {
		counterList = append(counterList, *val)
	}
	sort.Sort(sort.Reverse(counterList))
	return counterList
}

func writeToFile(fileName string, data string) error {
	err := os.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

func readUsingScanner(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data []string
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		data = append(data, scanner.Text())
	}
	return data, nil
}

func readUsingCSVModule(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.LazyQuotes = true

	return csvReader.ReadAll()
}
