package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type PolStruct struct {
	Name   string `cs:"name"`
	Number string `cs:"number"`
}

func main() {
	fmt.Println("Csv to Db started")
	f, err := os.Open("file.csv")
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
		return
	}
	titles := []string{}
	t := reflect.TypeOf(PolStruct{})
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Tag.Lookup("cs"))
		fmt.Println(t.Field(i).Name)
	}
	for _, record := range allRecords[0] {
		titles = append(titles, record)
	}

	x := reflect.StructOf([]reflect.StructField{
		{
			Name: strings.Title(titles[0]),
			Type: reflect.TypeOf(""),
		},
	})
	// g := reflect.New(x)
	fmt.Println(titles, x.Field(0).Name)
}
