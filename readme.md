# Conversor de Escalas Termométricas

Este projeto é um conversor simples de temperaturas entre Celsius, Fahrenheit e Kelvin, desenvolvido em Go.

## Como usar

1. Certifique-se de ter o [Go](https://golang.org/dl/) instalado em sua máquina.
2. Clone este repositório ou baixe o arquivo `main.go`.

## Executando o programa

No terminal, navegue até a pasta onde está o arquivo `main.go` e execute:

```sh
go run main.go
```

## Funcionamento

O programa solicitará que você informe a unidade de origem da temperatura (Celsius, Fahrenheit ou Kelvin) e, em seguida, o valor da temperatura. Ele exibirá as conversões para as outras duas escalas.

Exemplo de uso:

```
Digite a unidade de origem (Celsius = c, Fahrenheit = f ou Kelvin = k): c
Digite a temperatura em Celsius: 25
25.00°C = 77.00°F
25.00°C = 298.15K
```
## Builds
- Para Windows (64 bits):
```
GOOS=windows GOARCH=amd64 go build -o conversor.exe
```

- Para Mac (64 bits):
```
GOOS=darwin GOARCH=amd64 go build -o conversor-mac
```

- Para Mac (Apple Silicon/ARM):
```
GOOS=darwin GOARCH=arm64 go build -o conversor-mac-arm
```

- Para Linux (64 bits):
```
GOOS=linux GOARCH=amd64 go build -o conversor-linux
```

- Para Linux (ARM 64 bits):
```
GOOS=linux GOARCH=arm64 go build -o conversor-linux-arm
```


## Licença

Este projeto é de uso livre para fins educacionais.