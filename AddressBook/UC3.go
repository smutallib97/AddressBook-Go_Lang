package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func editContact() {
	var FirstName string
	var result bool
	fmt.Println("Enter first name of contact you want to edit")
	fmt.Scanln(&FirstName)
	file, err := os.Open("Addressbook.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	header, err := csvReader.Read()
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
	var choice int
	var updateddetails []Contact
	for _, record := range details {
		if record.FirstName != FirstName {
			updateddetails = append(updateddetails, record)
		}
		if record.FirstName == FirstName {
			result = true
			fmt.Printf("Enter the details you want to edit\n1.FirstName\n2.LastName\n3.Address\n4.City\n5.State\n6.ZipCode\n7.MobileNumber\n8.Email\n")
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				fmt.Println("Enter new first name: ")
				fmt.Scanln(&record.FirstName)

			case 2:
				fmt.Println("Enter new last name: ")
				fmt.Scanln(&record.LastName)
			case 3:
				fmt.Println("Enter new address: ")
				fmt.Scanln(&record.Address)

			case 4:
				fmt.Println("Enter new city: ")
				fmt.Scanln(&record.City)

			case 5:
				fmt.Println("Enter new state: ")
				fmt.Scanln(&record.State)

			case 6:
				fmt.Println("Enter new zip code: ")
				fmt.Scanln(&record.ZipCode)

			case 7:
				fmt.Println("Enter new phone number: ")
				fmt.Scanln(&record.MobileNumber)

			case 8:
				fmt.Println("Enter new email: ")
				fmt.Scanln(&record.Email)
			default:
				fmt.Println("Invalid Input")

			}
			updateddetails = append(updateddetails, record)

		}

	}
	f, err := os.Create("Addressbook.csv")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(f)
	defer w.Flush()
	_ = w.Write([]string(header))

	var writedata [][]string
	for _, record := range updateddetails {
		row := []string{record.FirstName, record.LastName, record.Address, record.City, record.State, strconv.Itoa(record.ZipCode), strconv.Itoa(record.MobileNumber), record.Email}
		writedata = append(writedata, row)
	}
	w.WriteAll(writedata)
	if result == true {
		fmt.Println("Contact edited Succesfully")
	} else {
		fmt.Println("Contact not found")
	}
}
