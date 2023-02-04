# Conversor Kelvin para Celsius

Este software é um exercício de [Treinamento DIO de Go Developer](https://dio.me/sign-up?ref=615HKOBZ22)

## Objetivo

Faça um programa que converta o ponto de ebulição da água em Kelvin para Celsius.

## Código em Linguagem Go

~~~go
package main

import "fmt"

func main() {
	// Boiling point of water in Kelvin: 373.2
	// Conversion formula for Celsius: Celsius = Kelvin - 273

	var kebulicao float32 = 373.2
	rcelsius := int(kebulicao - 273)
	fmt.Print("\n###    Conversor de Kelvin para Celsius    ####\n\n")
	fmt.Println("O ponto de ebulição da água em Celsius é:")
	fmt.Printf("%d °C\n\n", rcelsius)
}
~~~