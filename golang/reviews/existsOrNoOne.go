package reviews

import "fmt"

func ExistsOrNo() {

	// Aqui eu crio um dicionário 'map[string]string{}'.
	dictionary := map[string]string{
		"Abraão": "Santos",
		"Antônio": "Nunes",
		"Luiz": "Antônio",
	}

	// Crio uma variável chamada 'scan' do tipo string.
	var scan string
	fmt.Println("Deseja buscar: ")

	// Faço uma entrada de dados e, com ponteiro, adiciono o valor em scan.
	// & -> "Me da o endereço de memória dessa variável"
	fmt.Scan(&scan) 
	
	_, ok := dictionary[scan]

	if ok {
		fmt.Println("Encontrou.")
	} else {
		fmt.Println("Não encontrou.")
	}
}