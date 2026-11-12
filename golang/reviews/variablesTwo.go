package reviews

import "fmt"

func main() {

	// variável - nome da variável - tipo = valor da variável
	var myNameIs string = "My Name Is"
	name := "Tony"
	lastName := "Stark"

	// %s -> string
	// %d -> inteiro
	// %t -> booleano
	// %v -> qualquer valor

	// Aqui eu colocado as variáveis dentro da String
	fmt.Printf("%s %s %s\n", myNameIs, name, lastName)

	var myAgeIs int = 24
	fmt.Printf("\nMy age is: %d\n", myAgeIs)

	var bruceWayne string = "Bruce Wayne"
	// fmt.Sprintf SALVA o texto-string
	var message string = fmt.Sprintf("\nSeu nome é %s\n", bruceWayne)
	fmt.Printf(message)

	fmt.Printf("\n%v\n", bruceWayne)
}