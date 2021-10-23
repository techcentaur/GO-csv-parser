### Quick Report

- `csvFetch.go`: Fetches csv from url
	- usage: `go run csvFetch.go <url> <fileName.csv>`
	- [code](https://github.com/techcentaur/GO-csv-parser/blob/main/src/csvFetch.go)

- `csvParser.go`: parses csv into a data-structure + outputs most frequent value for each column on stdout and in a file:
	- usage: `go run csvParser.go <fileName.csv> <outputFileName>`
	- [code](https://github.com/techcentaur/GO-csv-parser/blob/main/src/csvParser.go)

- `uploadS3.go`: uploads file to S3
	- usage: `go run uploadS3.go <outputFileName>`
	- [code](https://github.com/techcentaur/GO-csv-parser/blob/main/src/uploadS3.go)
	
data-structure for `csvParser.go`:
```go
// CSV data would be stored in these data structure
type Dataframe struct {
	columns []Column
}
type Column struct {
	elements       Elements
	name, dataType string
}
type Elements []interface{}
```

data-structure for sorting per column:
```go
// FreqList has sorting structure defined on it using `Value` of Freq
type Freq struct {
	Key   interface{}
	Value int
}

func (f *Freq) addValue(i int) {
	f.Value += i
}

type FreqList []Freq
// this is for sorting basd on `Value`
func (f FreqList) Len() int           { return len(f) }
func (f FreqList) Less(i, j int) bool { return f[i].Value < f[j].Value }
func (f FreqList) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
```