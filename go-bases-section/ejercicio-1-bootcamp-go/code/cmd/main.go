package main

import "fmt"

func point1() {
	var lastName string = "Smith" // ningun cambio
	var age int = 35              // cambio de "35" a 35 para asignacion a un tipo de datos entero
	booleanVar := false           // cambios de "false" a false para asignacion a un tipo de datos booleano
	var salary float64 = 45857.90 // cambio de tipo de datos a float64
	var firstName string = "Mary"

	fmt.Println("lastName: ", lastName)
	fmt.Println("age: ", age)
	fmt.Println("booleanVar", booleanVar)
	fmt.Println("salary: ", salary)
	fmt.Println("firstName", firstName)
}

func main() {

	// Punto 1
	myName := "Alejandro Mora"
	myAdress := "Kr 35 # 20-34"

	fmt.Printf("Mi nombre es %v, y mi direccion es %s", myName, myAdress)

	// Punto 2
	type Place struct {
		Nombre             string
		Temperatura        int
		Humedad            int
		PresionAtmosferica float64
	}
	place1 := Place{
		Nombre:             "Bogotà",
		Temperatura:        10,
		Humedad:            48,
		PresionAtmosferica: 564,
	}
	place2 := Place{
		Nombre:             "Buenos Aires",
		Temperatura:        23,
		Humedad:            50,
		PresionAtmosferica: 1018,
	}

	places := []Place{place1, place2}

	for _, value := range places {
		fmt.Printf("Para la ciudad %v, su temperatura es %v, su humedad es %v y su presión atmosferica es %v", value.Nombre, value.Temperatura, value.Humedad, value.PresionAtmosferica)
	}

	// Punto 3
	var firstName string // esta linea fue corregida, los nombres de las variables no pueden empezar por un numero
	var lastName int     // el tipo de esta variable deberia ser int, ya que luego en el codigo se le asigna un valor numerico entero
	var age int          // la estructura de una variables siguiendo esta forma de declaracion debe ser var <nombreVariables> <tipo>

	lastName = 6 // para hacar referencia a lastname debes llamarala por su nombre sin el 1 al inicio

	var driverLicense bool = true // debemos devar una coherencia en el codgo, si todas la variables se estan declarando con camelcase no podemos usar tro formato dentro del programa
	var personHeight int          // Los nombres de la variables no deben contener espacios

	childsNumber := 2

	// - agrego valores a las variables y las imprimo en consola
	firstName = "Primer nombre"
	age = 19
	personHeight = 179

	fmt.Println("El valor de firstName es ", firstName)
	fmt.Println("El valor de lastName es ", lastName)
	fmt.Println("El valorde age es", age)
	fmt.Println("El valor de driverLicense es ", driverLicense)
	fmt.Println("El valor de personHeight es ", personHeight)
	fmt.Println("El valor de childsNumber es ", childsNumber)

	// Punto 4
	point1()
}
