package q5

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	// Implemente sua solução aqui
	var (
		TempConver float64
	)
	if fromScale != "F" && fromScale != "C" && fromScale != "K" || toScale != "F" && toScale != "C" && toScale != "K" {
		return 0, fmt.Errorf("escala inválida")
	}

	switch fromScale {
	case "C":
		if toScale == "F" {
			TempConver = temp*9/5 + 32
		} else {
			TempConver = temp + 273.15
		}
	case "F":
		if toScale == "C" {
			TempConver = (temp - 32) * 5 / 9
		} else {
			TempConver = (temp-32)*5/9 + 273.15
		}
	case "K":
		if toScale == "C" {
			TempConver = temp - 273.15
		} else {
			TempConver = (temp-273.15)*9/5 + 32
		}
	}

	return TempConver, nil
}

/*
Uma aplicação de clima precisa de uma função para converter a temperatura entre diferentes escalas: Celsius, Fahrenheit e Kelvin.
A função deve receber a temperatura, a escala de origem e a escala de destino.

As escalas de temperatura são representadas pelas seguintes letras:

C: Celsius
F: Fahrenheit
K: Kelvin
A função deve retornar a temperatura convertida na escala de destino. As fórmulas de conversão são:

Celsius para Fahrenheit: F = C * 9/5 + 32
Celsius para Kelvin: K = C + 273.15
Fahrenheit para Celsius: C = (F - 32) * 5/9
Fahrenheit para Kelvin: K = (F - 32) * 5/9 + 273.15
Kelvin para Celsius: C = K - 273.15
Kelvin para Fahrenheit: F = (K - 273.15) * 9/5 + 32
A função deve retornar um erro caso a escala de origem ou a escala de destino não sejam válidas, com a mensagem "escala inválida".
*/
