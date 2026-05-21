package algorithms

/*
Desafio: Contador de Palavras

Dada uma lista de palavras, conte quantas vezes cada palavra aparece e retorne um map com o resultado.

Casos de teste:
  - ["apple", "banana", "apple", "orange", "banana", "apple"] → {"apple": 3, "banana": 2, "orange": 1}
  - ["cat", "dog", "bird", "fish"]                            → {"cat": 1, "dog": 1, "bird": 1, "fish": 1}
  - ["go", "go", "go", "go", "python"]                       → {"go": 4, "python": 1}
  - ["hello"]                                                 → {"hello": 1}
  - ["a", "b", "a", "c", "b", "a", "c", "c", "c"]           → {"a": 3, "b": 2, "c": 4}
*/

func WordCounter(fruits []string) map[string]int {
	// Aqui constrói o HashMap que utilizarei pra armazenar as informações.
	mapFruits := make(map[string]int)
	
	// Realizo um iteração em todas as frutas.
	for _, fruit := range fruits{

		// Verifico se a fruta existe, sendo que 'valor da fruta (_), exists (bool (true/false)).
		_, exists := mapFruits[fruit]
		
		// Se a fruta não existe, salvo ela e acrescento seu valor de 1.
		if !exists {
			mapFruits[fruit] = 1
		} else {
			// Se existe, acrescento +1 no seu valor atual.
			mapFruits[fruit] += 1
		}
	}

	// Retorno o HashMap (dicionário)
	return mapFruits
}

// Testes que passaram: 1:  map[apple:3 banana:2 orange:1] 2:  map[bird:1 cat:1 dog:1 fish:1] 3:  map[go:4 python:1]