package q4

import "fmt"

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	// Implemente sua solução aqui
	var (
		FreteEstado float64
		aliquota    float64
		pfinal      float64
	)
	if taxCode != 1 && taxCode != 2 && taxCode != 3 {
		return 0, fmt.Errorf("imposto não encontrado")
	} else if basePrice <= 0 {
		return 0, fmt.Errorf("preço base inválido")
	}
	if state == "SP" {
		/*10%*/
		FreteEstado = 0.1
	} else if state == "RJ" {
		/*15%*/
		FreteEstado = 0.15
	} else if state == "MG" {
		/*20%*/
		FreteEstado = 0.2
	} else if state == "ES" {
		/*25%*/
		FreteEstado = 0.25
	} else {
		/*30%*/
		FreteEstado = 0.3
	}
	switch taxCode {
	case 1:
		/*5%*/
		aliquota = 0.05
	case 2:
		/*10%*/
		aliquota = 0.1
	case 3:
		/*15%*/
		aliquota = 0.15
	}

	pfinal = basePrice + ((basePrice * aliquota) + (basePrice * FreteEstado))

	return pfinal, nil
}

/* Uma empresa de comércio eletrônico precisa de uma função para calcular o preço final de um produto, que inclui o preço base, impostos e frete.
A função deve receber o preço base, o estado de destino e o código do imposto.

A alíquota do imposto e o percentual de frete se baseiam nas tabelas abaixo:

Código do imposto	Alíquota
1	5%
2	10%
3	15%
Estado de destino	Percentual de frete
SP	10%
RJ	15%
MG	20%
ES	25%
OUTROS	30%

A função deve retornar o preço final e um erro. O erro deve ser nulo (ou seja, nil) se não houver nenhum erro.
Caso o código do imposto não exista, a função deve retornar um erro com a mensagem "imposto não encontrado".
Caso o preço base seja menor ou igual a zero, a função deve retornar um erro com a mensagem "preço base inválido".
*/
