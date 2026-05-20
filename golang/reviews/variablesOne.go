package reviews

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)

	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	var (
		var1 string = "Hello"
		var2 string = "World"
	)
	fmt.Println(var1 + " " + var2)

	var3, var4 := "var", 50

	fmt.Println(var4)
	fmt.Println(var3)

}
