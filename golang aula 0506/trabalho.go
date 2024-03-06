package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     string `json:"age"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

var People []Person

func main() {
	for {
		showMenu()
		reader := bufio.NewReader(os.Stdin)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
		cmdString = strings.TrimSuffix(cmdString, "\n")
		cmdString = strings.TrimSpace(cmdString)

		switch cmdString {
		case "1":
			fmt.Println("Enter Person Details")
			addUser()
		case "2":
			fmt.Println("Get People")
			getPeople()
		case "3":
			fmt.Println("Delete Person")
			deletePerson()
		case "4":
			fmt.Println("Update Person")
			updatePerson()
		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid Option")
		}
	}
}

func showMenu() {
	fmt.Println("1. Add Person")
	fmt.Println("2. Get People")
	fmt.Println("3. Delete Person")
	fmt.Println("4. Update Person")
	fmt.Println("5. Exit")
}

func addUser() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Println("Enter Address: ")
	address, _ := reader.ReadString('\n')

	fmt.Println("Enter Age: ")
	age, _ := reader.ReadString('\n')

	fmt.Println("Enter Email: ")
	email, _ := reader.ReadString('\n')

	fmt.Println("Enter Phone: ")
	phone, _ := reader.ReadString('\n')

	person := Person{
		Name:    strings.TrimSpace(name),
		Address: strings.TrimSpace(address),
		Age:     strings.TrimSpace(age),
		Email:   strings.TrimSpace(email),
		Phone:   strings.TrimSpace(phone),
	}

	People = loadPeople()
	People = append(People, person)
	SaveUsers()
}

func getPeople() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name to search (or leave empty to list all): ")
	searchName, _ := reader.ReadString('\n')
	searchName = strings.TrimSpace(searchName)

	if searchName == "" {
		// Print all people
		for _, person := range People {
			fmt.Printf("Name: %s\nAddress: %s\nAge: %s\nEmail: %s\nPhone: %s\n\n", person.Name, person.Address, person.Age, person.Email, person.Phone)
		}
		return
	}

	// Search for people by name (case-insensitive)
	found := false
	for _, person := range People {
		if strings.ToLower(person.Name) == strings.ToLower(searchName) {
			fmt.Printf("Name: %s\nAddress: %s\nAge: %s\nEmail: %s\nPhone: %s\n\n", person.Name, person.Address, person.Age, person.Email, person.Phone)
			found = true
		}
	}

	if !found {
		fmt.Printf("Person with name '%s' not found\n", searchName)
	}
}

func deletePerson() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Name of the person to delete:")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}

	name = strings.TrimSpace(name)
	foundIndex := -1
	for i, person := range People {
		if person.Name == name {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("Person not found.")
		return
	}

	// Remove the person from the slice
	People = append(People[:foundIndex], People[foundIndex+1:]...)
	fmt.Println("Person deleted successfully.")
	SaveUsers() // Update the file after deletion
}

func updatePerson() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Name of the person to update:")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}

	name = strings.TrimSpace(name)
	foundIndex := -1
	for i, person := range People {
		if person.Name == name {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("Person not found.")
		return
	}

	fmt.Println("Enter new details (leave blank to keep existing):")

	fmt.Println("New Name:")
	newName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	newName = strings.TrimSpace(newName)

	fmt.Println("New Address:")
	newAddress, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	newAddress = strings.TrimSpace(newAddress)

	fmt.Println("New Age:")
	newAge, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	newAge = strings.TrimSpace(newAge)

	fmt.Println("New Email:")
	newEmail, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	newEmail = strings.TrimSpace(newEmail)

	fmt.Println("New Phone:")
	newPhone, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	newPhone = strings.TrimSpace(newPhone)

	// Update the person details
	if newName != "" {
		People[foundIndex].Name = newName
	}
	if newAddress != "" {
		People[foundIndex].Address = newAddress
	}
	if newAge != "" {
		People[foundIndex].Age = newAge
	}
	if newEmail != "" {
		People[foundIndex].Email = newEmail
	}
	if newPhone != "" {
		People[foundIndex].Phone = newPhone
	}

	fmt.Println("Person updated successfully.")
	SaveUsers() // Update the file after update
}

func loadPeople() []Person {
	file, err := os.ReadFile("people.json")
	if err != nil {
		fmt.Println("Error reading file")
	}
	_ = json.Unmarshal(file, &People)
	return People
}

func SaveUsers() {
	file, _ := json.MarshalIndent(People, "", " ")
	_ = os.WriteFile("people.json", file, 0644)
}
