package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	
	file, err := os.Open("test.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	fmt.Println(records[:5])
}
