package main

import (
	"csvtodb/pkg"
	"fmt"
)

func main() {
	fmt.Println("Csv to Db started")
	config := &pkg.Config{
		File: "file.csv",
		Match: []pkg.Match{
			{
				From: "name",
				To:   "identifier",
			},
		},
	}
	pkg.InsertData(config)
}
