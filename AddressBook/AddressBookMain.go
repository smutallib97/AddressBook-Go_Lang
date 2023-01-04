package main

import (
	"fmt"
)

var DuplicateFirstName string
var Duplicatecheck bool

func main() {
	fmt.Println("==============Welcome to Address Book Program================")
	FileCheck()
	var option int
	var choice int
	var exit bool = true
	for exit {
		fmt.Println("Select The option to perform operation given below")
		fmt.Printf("1.Add Contact\n2.Edit Contact\n3.Delete Contact\n4.Search Contact\n5.PersonNamebyState\n6.Count Of Contacts\n")
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			{
				addContact()

			}
		case 2:
			{
				editContact()
			}
		case 3:
			{
				deleteContact()
			}
		case 4:
			{
				searchContact()
			}
		case 5:
			{
				mapByStateCity()
			}
		case 6:
			{
				countOfContact()
			}
		default:
			{
				fmt.Println("Invalid Option Entered")
			}
		}
		fmt.Printf("\nDo you wish to continue with the program\nPress 1 to continue\n")
		fmt.Scanln(&choice)
		if choice == 1 {
			continue
		}
		break
	}
}
