package algorithms


/*
## Descrição
    Você recebe uma lista de inteiros.
    Seu desafio é contar quantas vezes cada número aparece na lista e retornar essa informação em uma estrutura de dados adequada.

### Objetivos
    Exemplo de entrada: numeros = [1, 2, 2, 3, 1, 4, 2, 3]

    Exemplo de saída: 
        {
            1: 2,
            2: 3,
            3: 2,
            4: 1
        }
*/

import ("fmt")

// Função FrequencyCounter, que terá o parâmetro que será um Array/Slice de int.
// O tipo de retorno será um dicionário (map) do tipo map[int]int
func FrequencyCounter(numbers[]int) map[int]int {

	// Aqui eu crio um dicionário 'make(map[int]int)' para inserir os dados
	// que serão das contagens.
	frequency := make(map[int]int)

	// A instrução 'range' itera sobre cada item de uma lista/dicionário (map, slice)
	for _, v := range numbers{
		fmt.Println("V:", v)
		// Preciso verificar se o número que eu estou iterando já existe salvo no frequency
		// Se não existir, eu salvo, e salvo também a quantidade.
		_, ok := frequency[v]

		// Se não existir
		if !ok {
			frequency[v] = 1
		} else {
		// Se existir
			frequency[v] += 1
		}
	}
	return frequency
}