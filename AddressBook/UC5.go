package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func filterDuplicate() {
	file, err := os.Open("Addressbook.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	_, err = csvReader.Read()
	rec, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	var data []Contact
	for _, record := range rec {
		Contacts := Contact{
			FirstName: record[0],
		}
		data = append(data, Contacts)
	}
	for _, record := range data {
		if record.FirstName == DuplicateFirstName {
			Duplicatecheck = true
		}

	}
}
