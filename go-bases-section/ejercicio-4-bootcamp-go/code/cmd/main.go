package main

import (
	"errors"
	"fmt"
)

// solucion implementada para el ejericio 3 del bootcamp

// Punto 1
type Employ struct {
	Name   string
	Salary int
}

func TaxValue(employ Employ) (result float64) {

	if employ.Salary > 50000 && employ.Salary <= 150000 {
		x := float64(employ.Salary)
		result = x - (float64(employ.Salary) * 0.17)
		return
	} else if employ.Salary > 150000 {
		x := float64(employ.Salary)
		result = x - (x * 0.27)
		return
	} else {
		result = float64(employ.Salary)
		return
	}
}

// Punto 2
func PromedioPorEstudiante(notas []int) (promedio float64, err error) {

	notasLength := len(notas)
	var sumatoria float64 = 0

	if notasLength == 0 {
		err = errors.New("no se ha ingresado ninguna nota")
		return
	}

	for _, value := range notas {
		if value < 0 {
			err = errors.New("todos los valores deben ser mayores a 0")
			return
		} else {
			sumatoria = sumatoria + float64(value)
		}
	}

	promedio = sumatoria / float64(notasLength)
	return
}

// Punto 3
type Category string

const (
	CATEGORY_C Category = "C"
	CATEGORY_B Category = "B"
	CATEGORY_A Category = "A"
)

type EmployPunto3 struct {
	Name                  string
	MinutesWorkedPerMouth float64
	Category
}

func CalcularSalarioEmpleado(employ EmployPunto3) (salary float64) {
	hoursWorked := employ.MinutesWorkedPerMouth / 60
	if employ.Category == CATEGORY_C {
		salary = hoursWorked * 1000
		return
	} else if employ.Category == CATEGORY_B {
		monthBaseSalary := hoursWorked * 1500
		salary = monthBaseSalary + (monthBaseSalary * (0.2))
		return
	} else {
		monthBaseSalary := hoursWorked * 3000
		salary = monthBaseSalary + (monthBaseSalary * (0.5))
		return
	}
}

// Punto 4: Calcular estadisticas
const (
	Minimum = "minimum"

	Average = "average"

	Maximum = "maximum"
)

func Operation(operationParam string) (e func(x ...int) (res float64), err error) {
	if operationParam == Minimum {
		e = MinFunc
		return
	} else if operationParam == Maximum {
		e = MaxFunc
		return
	} else if operationParam == Average {
		e = AverangeFunc
		return
	} else {
		err = errors.ErrUnsupported
		return
	}
}

// Operations availables
func MinFunc(e ...int) (result float64) {
	var min float64 = 100000000000
	for _, value := range e {
		if float64(value) < min {
			min = float64(value)
		}
	}
	result = min
	return
}

func MaxFunc(e ...int) (result float64) {
	var max float64 = 0
	for _, value := range e {
		if float64(value) > max {
			max = float64(value)
		}
	}
	result = max
	return
}

func AverangeFunc(e ...int) (result float64) {
	var sumatoria int = 0
	for _, value := range e {
		sumatoria += value
	}
	tam := len(e)
	result = float64(sumatoria) / float64(tam)
	return
}

// Punto 5
const (
	Dog       string = "dog"
	Cat       string = "cat"
	Hamster   string = "hamster"
	Tarantula string = "tarantula"
)

func Animal(animalType string) (e func(numero int) (cantidad float64), msg error) {
	if animalType == Dog {
		e = AnimalDog
		return
	} else if animalType == Cat {
		e = AnimalCat
		return
	} else if animalType == Hamster {
		e = AnimalHamster
		return
	} else if animalType == Tarantula {
		e = AnimalTarantula
		return
	} else {
		msg = errors.ErrUnsupported
		return
	}
}

func AnimalDog(numero int) (cantidad float64) {
	cantidad = float64(numero) * 10 // en kg
	return
}

func AnimalCat(numero int) (cantidad float64) {
	cantidad = float64(numero) * 5 // en kg
	return
}

func AnimalHamster(numero int) (cantidad float64) {
	cantidad = float64(numero) * 250 / 1000 // en kg
	return
}

func AnimalTarantula(numero int) (cantidad float64) {
	cantidad = float64(numero) * 150 / 1000 //en kg
	return
}

func main() {

	// Punto 1: Impuestos de salario
	myTax := Employ{
		Name:   "Alejandro Mora",
		Salary: 200000,
	}
	tax := TaxValue(myTax)
	fmt.Println("El valor de mi impuesto es ", tax)

	// Punto 2: Calcular promerdio
	notas := []int{1, 2, 3, 4, 5, 1, 2}
	promedio, err := PromedioPorEstudiante(notas)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("El promedio calculado es ", promedio)

	// Punto 3: Calcular salario
	empleado := EmployPunto3{
		Name:     "Alejandro",
		Category: CATEGORY_A,
	}
	salarioEmpleado := CalcularSalarioEmpleado(empleado)
	fmt.Printf("El salario para el empleado %v es %v", empleado.Name, salarioEmpleado)

	// Punto 4: Calcular estadisticas
	minFunc, err1 := Operation(Minimum)
	averageFunc, err2 := Operation(Average)
	maxFunc, err3 := Operation(Maximum)

	if err1 != nil {
		return
	} else if err2 != nil {
		return
	} else if err3 != nil {
		return
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("MIN", minValue)
	fmt.Println("max", maxValue)
	fmt.Println("averange", averageValue)

	// Punto 5: Calcular cantidad de alimento

	animalDog, msg1 := Animal(Dog)
	animalCat, msg2 := Animal(Cat)

	if msg1 != nil {
		return
	} else if msg2 != nil {
		return
	}

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)

	fmt.Printf("Cantidad total de alimento necesario es %v kg \n", amount)

}
