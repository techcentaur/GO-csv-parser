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

func read_using_scanner(filename string) (data []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(nil)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		data = append(data, scanner.Text())
	}
	return data
}

func read_using_csv_module(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(nil)
	}

	csvReader := csv.NewReader(file)
	csvReader.LazyQuotes = true

	records, err := csvReader.ReadAll()

	if err != nil {
		panic(err)
	}
	return records
}

type Dataframe struct {
	columns []Column
}

type Column struct {
	values    Values
	name      string
	data_type string
}

type Values []interface{}

func main() {
	filename := "test1.csv"
	data := read_using_csv_module(filename)

	df := new(Dataframe)

	// using first row to determine headers and second row to determine types
	for i, col := range data[0] {
		first_value := strings.TrimSpace(data[1][i])
		_type := reflect.TypeOf(first_value).String()

		if _, err := strconv.Atoi(first_value); err == nil {
			_type = "int"
		} else if _, err := strconv.ParseFloat(first_value, 64); err == nil {
			_type = "float64"
		} else if _, err := strconv.ParseBool(first_value); err == nil {
			_type = "bool"
		}
		df.columns = append(df.columns, Column{name: col, data_type: _type})
	}

	for _, col := range data[1:] {
		for j, val := range col {
			df.columns[j].values = append(df.columns[j].values, val)
		}
	}

	type frequentData struct {
		value     interface{}
		frequency int
	}
	mostFrequent := make(map[string]frequentData)

	for _, col := range df.columns {
		counter := make(map[interface{}]int)
		for _, val := range col.values {
			if _, ok := counter[val]; ok {
				counter[val] += 1
			} else {
				counter[val] = 1
			}
		}
		val := sortByFrequency(counter)[0]
		mostFrequent[col.name] = frequentData{value: val.Key, frequency: val.Value}
	}

	fmt.Print(mostFrequent)

}

func sortByFrequency(counter map[interface{}]int) FreqList {
	_list := make(FreqList, len(counter))

	i := 0
	for k, v := range counter {
		_list[i] = Freq{k, v}
		i++
	}
	sort.Sort(sort.Reverse(_list))
	return _list
}

type Freq struct {
	Key   interface{}
	Value int
}

type FreqList []Freq

func (f FreqList) Len() int           { return len(f) }
func (f FreqList) Less(i, j int) bool { return f[i].Value < f[j].Value }
func (f FreqList) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
