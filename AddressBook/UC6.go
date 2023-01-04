package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func searchContact() {
	var City string
	var State string
	var choice int
	var result bool
	fmt.Printf("Enter the option by which you want to search contact\n1.City\n2.State")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Enter City of Contact")
		fmt.Scanln(&City)
	case 2:
		fmt.Println("Enter State of Contact")
		fmt.Scanln(&State)
	}
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
	var details []Contact
	for _, record := range rec {
		zipCode, _ := strconv.Atoi(record[5])
		mobileNumber, _ := strconv.Atoi(record[6])

		Contacts := Contact{
			FirstName:    record[0],
			LastName:     record[1],
			Address:      record[2],
			City:         record[3],
			State:        record[4],
			ZipCode:      zipCode,
			MobileNumber: mobileNumber,
			Email:        record[7],
		}
		details = append(details, Contacts)
	}
	for _, record := range details {
		if record.City == City || record.State == State {
			result = true
			fmt.Println(record)
		}

	}

	if result != true {
		fmt.Println("None Contact Found")
	}
}
