#### Comandos no Golang

Iniciar a execução do ```main.go```
```bash
go run main.go
```

Criar o pacote para gerenciamento de dependências (go.mod)
```bash
go mod init modules
```

Buildar o código transformando em binários executáveis
```bash
go build <name>
go build -o ./bin/<name>
```

Instalar uma dependência
```bash
go get <url>
```



#### Sintaxe

Ponteiros servem para acessar valores e modificá-los na memória sem criar cópias.
Uso de & significa "guarda esse valor dentro dessa variável", ou seja, passa o endereço dela.
```golang
var name string

fmt.Scan(&name)
```