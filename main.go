package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

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
	// t := reflect.TypeOf(PolStruct{})

	for _, record := range data[0] {
		titles = append(titles, record)
	}
	return titles
}

func buildStruct(titles []string) {
	structFields := []reflect.StructField{}
	for _, title := range titles {
		tag := fmt.Sprintf("%s:\"%s\"", "cs", title)
		fmt.Println(tag)
		structFields = append(structFields, reflect.StructField{
			Name: strings.Title(title),
			Type: reflect.TypeOf(""),
			Tag:  reflect.StructTag(tag),
		})
	}
	t := reflect.StructOf(structFields)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Tag.Lookup("cs"))
		fmt.Println(t.Field(i).Name)
	}
	fmt.Println(titles)
}

func main() {
	fmt.Println("Csv to Db started")
	data, err := readCSV("file.csv")
	if err != nil {
		log.Fatal(err)
	}

	titles := getTitles(data)

	buildStruct(titles)
}
