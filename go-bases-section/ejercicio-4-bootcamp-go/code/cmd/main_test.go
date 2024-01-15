package main_test

import (
	main "ejercicio-4/cmd"
	"testing"

	"github.com/stretchr/testify/require"
)

// punto1
func TestTaxValue(t *testing.T) {
	t.Run("case 1: Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.", func(t *testing.T) {

		//act
		empleado := main.Employ{
			Name:   "Alejandro",
			Salary: 40000,
		}
		// assert
		result := main.TaxValue(empleado)
		var resultExpected float64 = 40000

		require.Equal(t, result, resultExpected)
	})
	t.Run("case 2: Calcular el impuesto en caso de que el empleado gane por encima de $50.000.", func(t *testing.T) {
		//act
		empleado := main.Employ{
			Name:   "Alejandro",
			Salary: 60000,
		}
		// assert
		result := main.TaxValue(empleado)
		var resultExpected float64 = 60000 - (60000 * 0.17)

		require.Equal(t, result, resultExpected)
	})
	t.Run("case 3: Calcular el impuesto en caso de que el empleado gane por encima de $150.000.", func(t *testing.T) {
		//act
		empleado := main.Employ{
			Name:   "Alejandro",
			Salary: 160000,
		}
		// assert
		result := main.TaxValue(empleado)
		var resultExpected float64 = 160000 - (160000 * 0.27)
		require.Equal(t, result, resultExpected)
	})
}

// punto2
func TestPromedioPorNotas(t *testing.T) {
	t.Run("case 1: Calcula el promedio", func(t *testing.T) {

		//given
		notas := []int{10, 10, 10, 10, 10}

		//act
		promedio, _ := main.PromedioPorEstudiante(notas)

		//assert
		var valueExpected float64 = 10
		require.Equal(t, promedio, valueExpected)
	})
}

// punto 3
func TestCalcularSalarioEmpleado(t *testing.T) {
	t.Run("case 1: Calcular el salario de la categoría A.", func(t *testing.T) {

		empleado := main.EmployPunto3{
			Name:                  "Alejandro",
			MinutesWorkedPerMouth: 16320,
			Category:              main.CATEGORY_A,
		}

		salary := main.CalcularSalarioEmpleado(empleado)

		var valueExpected float64 = (16320 / 60 * 3000) + (16320 / 60 * 3000 / 2)

		require.Equal(t, valueExpected, salary)
	})

	t.Run("case 2: Calcular el salario de la categoría B.", func(t *testing.T) {

		empleado := main.EmployPunto3{
			Name:                  "Alejandro",
			MinutesWorkedPerMouth: 16320,
			Category:              main.CATEGORY_B,
		}

		salary := main.CalcularSalarioEmpleado(empleado)

		var valueExpected float64 = (16320 / 60 * 1500) + (16320 / 60 * 1500 * 0.2)
		require.Equal(t, valueExpected, salary)
	})

	t.Run("case 3: Calcular el salario de la categoría C.", func(t *testing.T) {

		empleado := main.EmployPunto3{
			Name:                  "Alejandro",
			MinutesWorkedPerMouth: 16320,
			Category:              main.CATEGORY_C,
		}

		salary := main.CalcularSalarioEmpleado(empleado)

		var valueExpected float64 = (16320 / 60 * 1000)
		require.Equal(t, valueExpected, salary)
	})
}

// punto 4
func TestOperation(t *testing.T) {
	t.Run("case 1: Realizar test para calcular el mínimo de calificaciones.", func(t *testing.T) {

		minFunc, _ := main.Operation(main.Minimum)
		var min float64 = minFunc(1, 2, 3, 4, 5, 6, 99)

		var valueExpected float64 = 1

		require.Equal(t, valueExpected, min)

	})

	t.Run("case 2: Realizar test para calcular el máximo de calificaciones.", func(t *testing.T) {

		minFunc, _ := main.Operation(main.Maximum)
		var max float64 = minFunc(1, 2, 3, 4, 5, 6, 99)

		var valueExpected float64 = 99

		require.Equal(t, valueExpected, max)

	})

	t.Run("case 3: Realizar test para calcular el promedio de calificaciones.", func(t *testing.T) {

		minFunc, _ := main.Operation(main.Average)
		var avr float64 = minFunc(99, 99, 99, 99)

		var valueExpected float64 = 99

		require.Equal(t, valueExpected, avr)

	})
}

// Punto 5
func TestDog(t *testing.T) {
	t.Run("case dog: Verificar el cálculo de la cantidad de alimento para los perros.", func(t *testing.T) {
		funcion, _ := main.Animal(main.Dog)
		value := funcion(10) // 10 kg

		var valueExpected float64 = 10 * 10
		require.Equal(t, valueExpected, value)
	})
}
func TestCat(t *testing.T) {
	t.Run("case cat: Verificar el cálculo de la cantidad de alimento para los gatos.", func(t *testing.T) {
		funcion, _ := main.Animal(main.Cat)
		value := funcion(10)

		var valueExpected float64 = 10 * 5
		require.Equal(t, valueExpected, value)
	})
}
func TestHamster(t *testing.T) {
	t.Run("case hamster: Verificar el cálculo de la cantidad de alimento para los hamster.", func(t *testing.T) {
		funcion, _ := main.Animal(main.Hamster)
		value := funcion(10)

		var valueExpected float64 = 10.0 * 250 / 1000
		require.Equal(t, valueExpected, value)
	})
}

func TestTarantula(t *testing.T) {
	t.Run("case dog: Verificar el cálculo de la cantidad de alimento para las tarantulas.", func(t *testing.T) {
		funcion, _ := main.Animal(main.Tarantula)
		value := funcion(10)

		var valueExpected float64 = 10.0 * 150 / 1000
		require.Equal(t, valueExpected, value)
	})
}
