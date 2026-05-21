package reviews

import (
	"fmt"
)

func DataStructuresOne(){

	// Isso é um exemplo de Slice/Array
	numbers := []int{10, 20, 30, 40}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Valor:", numbers[i])
	}

	// Isso é um exemplo de HashMap
	mapExampleOne := map[string]int{
		"Abraão Santos": 20,
		"Luciano Silva": 30,
		"Heitor Alberto": 50,
	}

	var i int = 0
	for i < len(mapExampleOne) {
		fmt.Println(mapExampleOne["Abraão Santos"])
		i++
	}	

	for key, value := range mapExampleOne {
		fmt.Println("Chave", key + "Valor", value)
	}

	mapExampleTwo := map[string]string{
		"Abraão": "Santos",
	}

	// Nesse iteração, será pego apenas o valor
	for _, v := range mapExampleTwo {
		fmt.Println(
			"Value:", v,
		)
	}

	// Nesse iteração, será pego apenas a chave
	for k := range mapExampleTwo {
		fmt.Println(
			"Key:", k,
		)
	}

	// Uma list/slice de nomes
	names := []string{"Abraão", "Antônio", "Vitor", "Josué", "Jesus"}

	for i := range names {
		fmt.Println("Name: ", names[i])
	}
}