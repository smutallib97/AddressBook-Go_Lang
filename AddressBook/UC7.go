package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func mapCityAndPerson() {
	file, _ := os.Open("Addressbook.csv")
	defer file.Close()

	csvReader := csv.NewReader(file)
	_, _ = csvReader.Read()
	rec, _ := csvReader.ReadAll()

	var City []string
	var Person []string

	for _, record := range rec {

		City = append(City, record[3])
		Person = append(Person, record[0])
	}
	var mapOfCity = make(map[string]string)
	for i, _ := range City {
		mapOfCity[City[i]] = Person[i]
	}
	for key, value := range mapOfCity {
		fmt.Println(key, value)
	}

}
func mapStateAndPerson() {
	file, _ := os.Open("Addressbook.csv")
	defer file.Close()

	csvReader := csv.NewReader(file)
	_, _ = csvReader.Read()
	rec, _ := csvReader.ReadAll()

	var State []string
	var Person []string

	for _, record := range rec {

		State = append(State, record[4])
		Person = append(Person, record[0])
	}
	var mapOfState = make(map[string]string)
	for i, _ := range State {
		mapOfState[State[i]] = Person[i]
	}
	for key, value := range mapOfState {
		fmt.Println(key, value)
	}
}
func mapByStateCity() {
	var choice int
	fmt.Printf("Enter the choice which operation you want to perform\n1.Map of City and Person\n2.Map of State and Person\n")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		mapCityAndPerson()
	case 2:
		mapStateAndPerson()
	default:
		fmt.Println("Invalid Input")
	}
}
