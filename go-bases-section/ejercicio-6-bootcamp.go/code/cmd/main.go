package main

import (
	"errors"
	"fmt"
	"time"
)

// Punto 1
type Alumno struct {
	Name     string
	Apellido string
	DNI      int
	Fecha    time.Time
}

func (m Alumno) detalle() {
	fmt.Printf("Alumno %v %v con número de DNI %v el cual ingreso el dia %v \n", m.Name, m.Apellido, m.DNI, m.Fecha.String())
}

// Punto 2
type Category string

const (
	Small  Category = "small"
	Medium Category = "medium"
	Large  Category = "large"
)

type ProductInterface interface {
	Price(category Category, price int) (total float64, err error)
}

type Product struct {
	Name string
}

func (m Product) Price(category Category, price int) (total float64, err error) {
	if category == Small {
		total = float64(price)
		return
	} else if category == Medium {
		total = float64(price) + (float64(price) * 0.97)
		return
	} else if category == Large {
		total = float64(price) + (float64(price) * 0.94) + 2500
		return
	} else {
		err = errors.New("invalid category")
		return
	}
}

func main() {
	// Punto 1
	alumno := Alumno{
		Name:     "Alejandro",
		Apellido: "Mora",
		DNI:      1063552128,
		Fecha:    time.Now(),
	}
	alumno.detalle()

	// Punto 2
	product := Product{
		Name: "Producto 1",
	}

	price, _ := product.Price(Small, 10000)
	fmt.Println("El precio del producto es ", price)
}
