package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Punto 1: Hallar el numero de letras que tiene la palabra
func LenghtOfWord(word string) (length int) {
	// no contar espacios en el string
	var counter int = 0
	for _, value := range word {
		if value == rune(' ') {
			counter++
		} else {
			fmt.Println(strconv.QuoteRune(value))
		}
	}
	if len(word)-counter >= 0 {
		length = len(word) - counter
	} else {
		length = 0
	}
	return
}

// Punto 2: Prestamo
type Customer struct {
	Name             string
	Age              int
	IsEmployed       bool
	YearsOnActualJob float64
	Salary           int
}

func Prestamos(customers []Customer) {
	for _, value := range customers {
		if value.Age > 22 && value.IsEmployed && value.YearsOnActualJob > 1 {
			if value.Salary > 100000 {
				fmt.Printf("El cliente %v puede acceder al prestamo sin intereses \n", value.Name)
			} else {
				fmt.Printf("El cliente %v puede acceder al prestam con intereses \n", value.Name)
			}
		} else {
			fmt.Printf("No es posible otorgar a el cliente %v el prestamo \n", value.Name)
		}
	}
}

// Punto 3: A que mes corresponse
func MesCorrespondiente(month int) (result string, err error) {
	// La solucion que he escogido es hacer un arreglo pero no es necesario recorrerlo

	if month > 12 || month < 1 {
		err = errors.New("el numero de mes debe ser menor a 12 y mayor de 0")
		return
	}

	index := month - 1
	months := [12]string{
		"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
		"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre",
	}
	result = months[index]
	return
}

func main() {

	// Punto 1
	word1 := "Alejan dro"
	tamaño := LenghtOfWord(word1)
	fmt.Println("Cantidad de letras", tamaño)

	// Punto 2

	var customer1 Customer = Customer{
		Name:             "Roberto",
		Age:              40,
		IsEmployed:       true,
		YearsOnActualJob: 2,
		Salary:           200000,
	}
	var customer2 Customer = Customer{
		Name:             "Sol",
		Age:              40,
		IsEmployed:       true,
		YearsOnActualJob: 3,
		Salary:           70000,
	}
	var customer3 Customer = Customer{
		Name:             "Alejandro",
		Age:              17,
		IsEmployed:       false,
		YearsOnActualJob: 0,
		Salary:           0,
	}

	aspirantes := []Customer{customer1, customer2, customer3}
	Prestamos(aspirantes)

	// Punto 3
	mes1, err1 := MesCorrespondiente(1)
	mes12, err2 := MesCorrespondiente(12)
	mes9, err3 := MesCorrespondiente(9)

	if err1 != nil {
		fmt.Println(err1)
		return
	} else if err2 != nil {
		fmt.Println(err2)
		return
	} else if err3 != nil {
		fmt.Println(err3)
		return
	}

	fmt.Printf("El mes numero 1 es %v \n", mes1)
	fmt.Printf("El mes numero 12 es %v \n", mes12)
	fmt.Printf("El mes numero 9 es %v \n", mes9)

	// Punto 4
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es ", employees["Benjamin"])

	counter := 0
	for _, value := range employees {
		if value > 21 {
			counter++
		}
	}
	fmt.Printf("Hay %v empleados mayores de 21 \n", counter)

	employees["Federico"] = 25
	delete(employees, "Pedro")
}
