package pkg

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type Match struct {
	From string
	To   string
}

type Config struct {
	File  string
	Match []Match
}

func readCSV(file string) (data [][]string, err error) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Cannot open")
		fmt.Println(err)
		return
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return [][]string{}, err
	}
	return allRecords, nil
}

func getTitles(data [][]string) []string {
	titles := []string{}
	for _, record := range data[0] {
		titles = append(titles, record)
	}
	return titles
}

func buildStruct(titles []string, val []string) reflect.Type {
	structFields := []reflect.StructField{}
	for index, title := range titles {
		tag := fmt.Sprintf("%s:\"%s\"", "csd", title)
		fmt.Println(tag)
		structFields = append(structFields, reflect.StructField{
			Name: strings.Title(title),
			Type: reflect.TypeOf(val[index]),
			Tag:  reflect.StructTag(tag),
		})
	}
	t := reflect.StructOf(structFields)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Tag.Lookup("csd"))
		fmt.Println(t.Field(i).Name, t.Field(i).Type)
	}
	return t
}

func getColumnTypes(titles interface{}) {

}

func InsertData(c *Config) {
	data, err := readCSV(c.File)
	if err != nil {
		log.Fatal(err)
	}
	titles := getTitles(data)

	s := buildStruct(titles, data[1])
	fmt.Println(s)
}
