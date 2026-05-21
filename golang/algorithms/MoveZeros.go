package algorithms

import (
	"log"
)

/*
DESAFIO LEETCODE: MOVER ZEROS (Move Zeroes)

O PROBLEMA:
Dado um slice de inteiros chamado 'nums', mova todos os valores '0'
para o FINAL dele, mantendo a ordem relativa dos outros números.

REGRAS:
1. Deve ser feito "in-place" (modificar o slice original diretamente).
2. Não é permitido criar um novo slice auxiliar.
3. Tente minimizar o número de operações.

EXEMPLO 1:
- Entrada: [0, 1, 0, 3, 12]
- Saída:   [1, 3, 12, 0, 0]

EXEMPLO 2:
- Entrada: [0]
- Saída:   [0]

DICA:
Use a técnica de "Dois Ponteiros" (Two Pointers). Um ponteiro busca
números diferentes de zero e o outro guarda a posição onde o próximo
número válido deve ser inserido.
*/

func MoveZerosOne(moveSlices[]int) ([]int) {

	zerosSlice := []int{}
	log.Println("1. Lista com zeros: ", zerosSlice)
	finalSlice := []int{}
	log.Println("2. Lista sem zeros: ", finalSlice)


	for _, v := range moveSlices{
		if v != 0 {
			log.Println("Lista sem zeros: ", finalSlice)
			finalSlice = append(finalSlice, v)
		} else {
			log.Println("Lista com zeros: ", finalSlice)
			zerosSlice = append(zerosSlice, 0)
		}
	}

	finalSlice = append(finalSlice, zerosSlice...)
	return finalSlice
}


/*
	Essa lógica é bem mais interessante, invés de ter dois slices
	um que contém a quantidade de zeros e outro que contém os números
	que não são zeros, eu basicamente faço um incremento na quantidade
	de zeros encontrados e faço um loop aplicando append(finalSlize, 0).
*/
func MoveZerosTwo(moveSlices[]int) ([]int) {
	var amountZeros int = 0
	finalSlice := []int{}

	for _, value := range moveSlices {
		if value != 0 {
			finalSlice = append(finalSlice, value)
		} else {
			amountZeros++
		}
	}

	for i := 0; i < amountZeros; i++ {
		finalSlice = append(finalSlice, 0)
	}

	return finalSlice
}