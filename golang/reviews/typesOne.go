package reviews

// Para inteiros (int), têm os seguintes tipos de dados:
// int, int8, int16, int32, int64
// O 'int' isolado utiliza da arquitetura do seu computador.
// A numerologia se refere a quantidade de bites que a variável suporta.

// Há também o uint, que é uint (unsigned integer)
// Ele suport apenas números positivos e zero, nada negativo.

import "fmt"

func TypesOne() {
	var numberOne int16 = 100
	var numberTwo int8 = 23

	var value int32 = 20
	
	valuePointer := &value

	fmt.Println("Endereço 'value':", &valuePointer)
	fmt.Println("Endereço 'valuePointer':", &valuePointer)
	fmt.Println("Valor 'value':", *valuePointer)
	fmt.Println("Valor 'valuePointer':", *valuePointer)
	*valuePointer = 30

	fmt.Println("Endereço 'value':", &valuePointer)
	fmt.Println("Endereço 'valuePointer':", &valuePointer)
	fmt.Println("Valor 'value':", *valuePointer)
	fmt.Println("Valor 'valuePointer':", *valuePointer)
	
	fmt.Println(numberOne)
	fmt.Println(numberTwo)

	if numberOne > 100 {
		fmt.Println("Hello")
	} else {
		fmt.Print(numberTwo)
	}

	var name string

	fmt.Println("Guarde o nome da pessoa: ")
	fmt.Scan(&name)
	
	if name == "Abraão" {
		fmt.Println("Digitou o nome correto!", name)
	} else {
		fmt.Println("Digitou o nome errado.")
	}

}