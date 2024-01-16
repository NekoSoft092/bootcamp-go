package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Punto 1
var (
	Path string = "llamadas123_noviembre_2023.csv"
)

type Llamada struct {
	NumeroIncidente        string
	Fecha                  string
	CodigoLocalidad        string
	Localidad              string
	Edad                   int
	Unidad                 string
	Genero                 string
	TipoIncidente          string
	TipoIncidenteAdicional string
	PrioridadInicial       string
	Recepcion              string
}

func ReadFile(path string) (f *os.File) {
	funcion, err := os.Open(path)
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	fmt.Println("ejecución finalizada")
	f = funcion
	return
}

// Punto 2

func main() {

	// recover app from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error message:", r)
		}
	}()

	// reading file that doesnt exist
	// point := ReadFile("./data.txt")
	// defer point.Close()

	f := ReadFile(Path)
	rd := csv.NewReader(f)
	rd.Read() // skip header

	counter := 0
	// incidentes := []Llamada{}

	for {
		record, err := rd.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			return
		}

		// Take information
		//id, errId := strconv.Atoi(record[0])
		array := strings.Split(record[0], ";")

		if array[0] == "" {
			break
		}

		age, errAge := strconv.Atoi(array[4])
		if errAge != nil {
			break
		}

		incidente := Llamada{
			NumeroIncidente:        array[0],
			Fecha:                  array[1],
			CodigoLocalidad:        array[2],
			Localidad:              array[3],
			Edad:                   age,
			Unidad:                 array[5],
			Genero:                 array[6],
			TipoIncidente:          array[7],
			TipoIncidenteAdicional: array[8],
			PrioridadInicial:       array[9],
			Recepcion:              array[10],
		}

		fmt.Println(incidente)

		/*incidentes = append(incidentes, incidente)*/
		counter++

		/*
			if errId != nil {
				fmt.Println(errId.Error())
				return
			}

			fmt.Println(id)*/
	}
	fmt.Println(counter)

}
