package main

import "fmt"

func main() {
	// 373.2 = ebolição da água em kelvin
	// A fórmula é: C = K - 273

	var kebulicao float32 = 373.2
	rcelsius := int(kebulicao - 273)
	fmt.Print("\n###    Conversor de Kelvin para Celsius    ####\n\n")
	fmt.Println("O ponto de ebulição da água em Celsius é:")
	fmt.Printf("%d °C\n\n", rcelsius)
}
