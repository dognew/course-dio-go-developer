# Converter Kelvin to Celsius [:brazil:](leiame.md)

This software is a exercise of the [Training DIO of Go Developer](https://dio.me/sign-up?ref=615HKOBZ22)

## Objetive

Make a program that converts the boiling point of water in Kelvin to Celsius.

## Script in Go Language

~~~go
package main

import "fmt"

func main() {
	// Boiling point of water in Kelvin: 373.2
	// Conversion formula for Celsius: Celsius = Kelvin - 273

	var kebulicao float32 = 373.2
	rcelsius := int(kebulicao - 273)
	fmt.Print("\n###    Kelvin to Celsius Converter    ####\n\n")
	fmt.Println("The boiling point of water in Celsius is:")
	fmt.Printf("%d °C\n\n", rcelsius)
}
~~~