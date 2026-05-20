package main

import ("fmt")

// alias
// INT32 = RUNE
// BYTE = UINT8

// Sobre números reais, tem-se o float32 e o float64

func main(){
	var uIntNumber uint16 = 20
	fmt.Println(uIntNumber)
	var numberThree float32 = 40.40
	var numberFour float64 = 70.333
	fmt.Printf("Número: %.2f | Número: %.3f\n", numberThree, numberFour)


	var compareOne *int64
	var compareTwo int64 = 20
	
	fmt.Println("Insira o valor: ")
	fmt.Scan(&compareOne)

	if compareOne == compareTwo {
		fmt.Println("Sim")
	} else {
		fmt.Println("Não")
	}

}