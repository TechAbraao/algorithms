package algorithms

/*
    ## Descrição
    Você recebe uma lista de inteiros.
    Seu desafio é encontrar o primeiro número que não se repete na lista.

    ### Objetivos
    Exemplo um de entrada: numeros = [4, 5, 1, 2, 1, 4, 5, 3]
    Exemplo dois de entrada: numeros = [1, 2, 3, 2, 1]
    Exemplo três de entrada: numeros = [7, 7, 8, 8, 9, 9]

    Exemplo um de saída: 2
    Exemplo dois de saída: 3
    Exemplo três de saída: None
*/


func DoNotRepeat(numbers[]int) (int, bool) {

	/*
	Antes de tudo, eu faço um loop em todos os números existentes na lista enviada
	via parâmetro. Preciso fazer com que o valor iterado atual tenha algo repetido!!
	*/
	for _, n := range numbers {

		/* essa variável armazena o estado da repetição. */
		var repeat int = 0

		/* 
		Outro loop O(N²) é problemático, porém nesse loop eu verifico na iteração atual se ele repete em todo
		conjunto da lista e incremento no 'repeat'.
		*/
		for _, y := range numbers {

			// Ou seja, se a iteração primária for igual a iteração secundária, incrementa
			// Isso significa na realidade que o item repetiu.
			if n == y {
				repeat++
			} 
		}

		// Agora eu verifico se a iteração primária se tiver APENAS 1 repetido, implica
		// Que só existe ele, não mais outro, logo encontramos o correto.
		if repeat == 1 {
			return n, true
		}
	}

	// Essa lógica é pra ocasiões/cenários que a repetição foi igual, tipo isso: "[]int{7, 7, 8, 8, 9, 9}"
	return 0, false
}


/*
package main

import (
	"fmt"
	"modules/algorithms"
)

func main() {
	result1, ok1 := algorithms.DoNotRepeat([]int{4, 5, 1, 2, 1, 4, 5, 3})
	fmt.Println("Result:", result1)
	fmt.Println("Ok:", ok1)


	result2, ok2 := algorithms.DoNotRepeat([]int{1, 2, 3, 2, 1})
	fmt.Println("Result:", result2)
	fmt.Println("Ok:", ok2)


	result3, ok3 := algorithms.DoNotRepeat([]int{7, 7, 8, 8, 9, 9})
	fmt.Println("Result:", result3)
	fmt.Println("Ok:", ok3)
}
*/