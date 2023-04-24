package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	// Implemente sua solução aqui
	var (
		NumAtual int
		Maior    int
		MenorNum int
		soma     float64
	)
	listMed := []int{}
	/*Descobrir o Maior Valor*/
	if len(numbers) > 0 {
		Maior = numbers[0]
	}
	for i := 0; i < len(numbers); i++ {
		NumAtual = numbers[i]
		if NumAtual > Maior {
			Maior = NumAtual
		}
	}
	/*Descobrir o Menor Valor*/
	MenorNum = Maior
	for _, i := range numbers {
		if i < MenorNum {
			MenorNum = i
		}

	}
	/*Fazer a Média da lista*/
	if len(numbers) == 1 {
		media := float64(numbers[0])
		return MenorNum, Maior, media, nil
	} else {
		for _, p := range numbers {
			if p != Maior && p != MenorNum {
				listMed = append(listMed, p)
			}
		}
		for l := 0; l < len(listMed); l++ {
			num := listMed[l]
			soma += float64(num)
		}
	}
	media := soma / float64(len(listMed))
	if len(numbers) != 0 {
		return Maior, MenorNum, media, nil
	} else {
		return 0, 0, 0, fmt.Errorf("A lista está vazia")
	}

}
