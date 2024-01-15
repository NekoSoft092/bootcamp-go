package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Punto 1
type Product struct {
	ID          string
	Name        string
	Price       int
	Description string
	Category    string
}

var product1 Product = Product{
	ID:          "1",
	Name:        "producto 1",
	Price:       3000,
	Description: "Description 1",
	Category:    "A",
}

var product2 Product = Product{
	ID:          "2",
	Name:        "producto 2",
	Price:       4000,
	Description: "Description 2",
	Category:    "B",
}

var product3 Product = Product{
	ID:          "3",
	Name:        "producto 3",
	Price:       4000,
	Description: "Description 3",
	Category:    "B",
}

var product4 Product = Product{
	ID:          "4",
	Name:        "producto 4",
	Price:       44000,
	Description: "Description 4",
	Category:    "C",
}

var (
	products = []Product{product1, product2, product3, product4}
)

// metodos
func (m *Product) Save() (added Product) {
	newProduct := (*m)
	newSlice := append(products, newProduct)
	sliceLen := len(newSlice)
	return newSlice[sliceLen-1]
}

func (m Product) GetAll() (result []Product) {
	result = products
	return
}

func (m Product) GetById(id int) (product Product, err error) {
	for _, value := range products {
		if value.ID == strconv.Itoa(id) {
			product = value
			return
		}
	}
	err = errors.New("not found")
	return
}

// Punto 2
type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Employ struct {
	ID       int
	Position string
	Person   Person
}

func (m Employ) PrintEmploy() {
	fmt.Println("ID", m.ID)
	fmt.Println("Position", m.Position)
	fmt.Println("Person", m.Person)
}

func main() {

	// Punto 1
	product := product1.Save()
	fmt.Println("product added ", product)

	allProducts := product1.GetAll()

	for _, value := range allProducts {
		fmt.Println(value.Name)
	}

	productFind, _ := product1.GetById(2)
	fmt.Println("El producto de id 1 es ", productFind)

	// Punto 2
	person := Person{
		ID:          1,
		Name:        "Alejandro Mora",
		DateOfBirth: time.Now(),
	}

	employ := Employ{
		ID:       1,
		Position: "Developer",
		Person:   person,
	}

	employ.PrintEmploy()

}
