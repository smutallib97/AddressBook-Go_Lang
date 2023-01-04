package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func addContact() {

	var details Contact
	fmt.Println("Enter First Name")
	fmt.Scanln(&details.FirstName)
	DuplicateFirstName = details.FirstName
	filterDuplicate()
	fmt.Println("Enter Last Name")
	fmt.Scanln(&details.LastName)
	fmt.Println("Enter Address")
	fmt.Scanln(&details.Address)
	fmt.Println("Enter City")
	fmt.Scanln(&details.City)
	fmt.Println("Enter State")
	fmt.Scanln(&details.State)
	fmt.Println("Enter Zip")
	fmt.Scanln(&details.ZipCode)
	fmt.Println("Enter Mobile")
	fmt.Scanln(&details.MobileNumber)
	fmt.Println("Enter Email")
	fmt.Scanln(&details.Email)

	if Duplicatecheck != true {
		f, err := os.OpenFile("Addressbook.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		defer f.Close()
		if err != nil {
			log.Fatalln("failed to open file", err)
		}
		writer := csv.NewWriter(f)
		defer writer.Flush()
		row := []string{details.FirstName, details.LastName, details.Address, details.City, details.State, strconv.Itoa(details.ZipCode), strconv.Itoa(details.MobileNumber), details.Email}
		writer.Write(row)
		fmt.Println("\nContact Added Succesfully")
	} else {
		fmt.Println("\nDuplicate Entry Cannot add Contact")
	}
}

func FileCheck() {
	_, err := os.Stat("Addressbook.csv")
	if os.IsNotExist(err) {
		file, err := os.Create("Addressbook.csv")
		defer file.Close()
		if err != nil {
			log.Fatalln("failed to open file", err)
		}
		writer := csv.NewWriter(file)
		defer writer.Flush()
		_ = writer.Write([]string{"First Name", "Last Name", "Address", "City", "State", "ZipCode", "MobileNumber", "Email"})
	}
}
