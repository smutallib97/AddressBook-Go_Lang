package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Contact struct {
	FirstName    string
	LastName     string
	Address      string
	City         string
	State        string
	ZipCode      int
	MobileNumber int
	Email        string
}

func main() {

	var err error
	db, err = sql.Open("mysql", "root:A@liya2020@tcp(127.0.0.1:3306)/addressbook")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("==============Welcome to Address Book Program================")

	var option int
	var choice int
	var exit bool = true
	for exit {
		fmt.Println("Select The option to perform operation given below")
		fmt.Println("1. View all contacts from address book")
		fmt.Println("2. Search contact by using City or State")
		fmt.Println("3. Exit")
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Println("Contacts in Address Book: ")
			viewContacts()
		case 2:
			fmt.Println("Contacts by using City or State")
			searchContact()
		case 3:
			fmt.Println("Thank You for using address book")
		default:
			fmt.Println("Invalid Input")
		}
		fmt.Printf("\nDo you wish to continue with the program\nPress 1 to continue\n")
		fmt.Scanln(&choice)
		if choice == 1 {
			continue
		}
		break
	}
}
func viewContacts() {
	rows, err := db.Query("select * from addressbook_data;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var contactDetails []Contact
	for rows.Next() {

		var zipCode, mobileNumber int
		var firstName, lastName, address, city, state, email string
		if err := rows.Scan(&firstName, &lastName, &address, &city, &state, &zipCode, &mobileNumber, &email); err != nil {
			log.Fatal(err)
		}
		newContact := Contact{FirstName: firstName, LastName: lastName, Address: address, City: city, State: state, ZipCode: zipCode, MobileNumber: mobileNumber, Email: email}
		contactDetails = append(contactDetails, newContact)
	}
	fmt.Println(contactDetails)
}
func searchContact() {
	var City string
	var State string
	var choice int
	fmt.Println("Enter the option by which you want to search contact\n1.City\n2.State")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Enter City of Contact")
		fmt.Scanln(&City)
	case 2:
		fmt.Println("Enter State of Contact")
		fmt.Scanln(&State)
	}
	rows, err := db.Query("SELECT * FROM addressbook_data WHERE city = ? or state = ?", City, State)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var contactDetails []Contact
	for rows.Next() {
		var zipCode, mobileNumber int
		var firstName, lastName, address, city, state, email string
		if err := rows.Scan(&firstName, &lastName, &address, &city, &state, &zipCode, &mobileNumber, &email); err != nil {
			log.Fatal(err)
		}
		newContact := Contact{FirstName: firstName, LastName: lastName, Address: address, City: city, State: state, ZipCode: zipCode, MobileNumber: mobileNumber, Email: email}
		contactDetails = append(contactDetails, newContact)
	}
	fmt.Println(contactDetails)

}
