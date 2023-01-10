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
		fmt.Println("3. Count of Contacts")
		fmt.Println("4. Add Contact in Address Book")
		fmt.Println("5. Update Contact details of Address Book")
		fmt.Println("6. Delete Contact from Address Book")
		fmt.Println("7. Exit")
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Println("Contacts in Address Book: ")
			viewContacts()
		case 2:
			fmt.Println("Contacts by using City or State")
			searchContact()
		case 3:
			fmt.Println("Count of Contacts by City --")
			countOfContact()
		case 4:
			addContact()
			fmt.Println("Contact Added Successfully!!!")
		case 5:
			updateContact()
		case 6:
			fmt.Print("Enter first name to delete contact: ")
			var firstName string
			fmt.Scanln(&firstName)
			deleteContact(firstName)
		case 7:
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
func countOfContact() {
	var City string
	fmt.Println("Enter City Name")
	fmt.Scanln(&City)
	var count int
	err := db.QueryRow("SELECT COUNT(city) FROM addressbook_data WHERE city = ?", City).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Contacts count of %s is %d", City, count)
}
func addContact() []Contact {

	newContact := &Contact{}

	fmt.Println("Enter First name: ")
	fmt.Scanln(&newContact.FirstName)
	fmt.Println("Enter Last name: ")
	fmt.Scanln(&newContact.LastName)
	fmt.Println("Enter address: ")
	fmt.Scanln(&newContact.Address)
	fmt.Println("Enter city: ")
	fmt.Scanln(&newContact.City)
	fmt.Println("Enter state: ")
	fmt.Scanln(&newContact.State)
	fmt.Println("Enter zip code: ")
	fmt.Scanln(&newContact.ZipCode)
	fmt.Println("Enter mobile number: ")
	fmt.Scanln(&newContact.MobileNumber)
	fmt.Println("Enter email: ")
	fmt.Scanln(&newContact.Email)

	_, err := db.Exec("INSERT INTO addressbook_data (firstName,lastName, address, city, state, zipcode, mobileNumber, email) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", &newContact.FirstName, &newContact.LastName, &newContact.Address, &newContact.City, &newContact.State, &newContact.ZipCode, &newContact.MobileNumber, &newContact.Email)
	if err != nil {
		log.Fatal(err)
	}
	var contactDetails []Contact
	contactDetails = append(contactDetails, *newContact)
	return contactDetails
}
func updateContact() {
	var FirstName string
	var result bool
	fmt.Println("Enter first name of contact you want to edit")
	fmt.Scanln(&FirstName)
	var zipCode, mobileNumber int
	var firstName, lastName, address, city, state, email string
	var choice int
	if FirstName == FirstName {
		result = true
		fmt.Printf("Enter the details you want to edit\n1.FirstName\n2.LastName\n3.Address\n4.City\n5.State\n6.ZipCode\n7.MobileNumber\n8.Email\n")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter new first name: ")
			fmt.Scanln(firstName)

		case 2:
			fmt.Println("Enter new last name: ")
			fmt.Scanln(lastName)
		case 3:
			fmt.Println("Enter new address: ")
			fmt.Scanln(address)

		case 4:
			fmt.Println("Enter new city: ")
			fmt.Scanln(city)

		case 5:
			fmt.Println("Enter new state: ")
			fmt.Scanln(state)

		case 6:
			fmt.Println("Enter new zip code: ")
			fmt.Scanln(zipCode)

		case 7:
			fmt.Println("Enter new phone number: ")
			fmt.Scanln(mobileNumber)

		case 8:
			fmt.Println("Enter new email: ")
			fmt.Scanln(email)
		default:
			fmt.Println("Invalid Input")

		}
		var contactDetails []Contact
		newContact := Contact{FirstName: firstName, LastName: lastName, Address: address, City: city, State: state, ZipCode: zipCode, MobileNumber: mobileNumber, Email: email}
		contactDetails = append(contactDetails, newContact)
		rows, err := db.Query("UPDATE addressbook_data SET firstName = ?, lastName = ?, address = ?, city = ?, state = ?, zipcode = ?, mobileNumber = ?, email = ? ", firstName, lastName, address, city, state, zipCode, mobileNumber, email)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		//var contactDetails []Contact
		///for rows.Next() {
		// var zipCode, mobileNumber int
		// var firstName, lastName, address, city, state, email string
		//if err := rows.Scan(&firstName, &lastName, &address, &city, &state, &zipCode, &mobileNumber, &email); err != nil {
		//	log.Fatal(err)
		//}
		//newContact := Contact{FirstName: firstName, LastName: lastName, Address: address, City: city, State: state, ZipCode: zipCode, MobileNumber: mobileNumber, Email: email}
		//contactDetails = append(contactDetails, newContact)
		//}
		fmt.Println(contactDetails)

	}
	if result == true {
		fmt.Println("Contact edited Succesfully")
	} else {
		fmt.Println("Contact not found")
	}
}
func deleteContact(firstName string) {
	_, err := db.Exec("DELETE FROM addressbook_data WHERE firstName = ?", firstName)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Contact deleted successfully!")
}
