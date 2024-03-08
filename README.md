# prova1-m9

## Requisitos

Golang (Go)
Mosquitto MQTT Broker

## Instalação

1. Instale o Mosquitto MQTT Broker. Você pode seguir as instruções de instalação no site oficial: [mosquitto.org](https://mosquitto.org/download/).


2. Navegue até o diretório do projeto:

    ```bash
    cd prova1-m9
    ```

## Uso

1. Inicie o broker MQTT Mosquitto (se ainda não estiver em execução). Se estiver usando Linux ou macOS, você pode iniciar o broker com:

    ```bash
    mosquitto -c mosquitto.conf
    ```

    Certifique-se de que o arquivo de configuração `mosquitto.conf` esteja apontando para o listener na porta 1891 (ou ajuste o código e o arquivo de configuração conforme necessário).

### Instalando Dependências - Go Mod

Acesse o diretorio que contem as dependências necessárias para cada função: 

Acione as dependências para cada uma das pastas, com: 
```
go mod tidy
```
Após isso, execute o publisher com:

```
go run publisher.go
```

E o subscriber no diretorio com:

```
mosquitto_sub -h localhost -p 1891 -v -t "/p1" 

```

### Testes

No diretório que você executou o `publisher.go` execute o seguinte comando:

```
go test -v -cover
```

Resultado esperado:

```
=== RUN   TestConection
    prova_test.go:20: Conection with broker MQTT successful
--- PASS: TestConection (0.00s)
PASS
coverage: 0.0% of statements
ok  	p1	0.003s

```
## Teste de Conexão com o Broker

### Propósito

Recebimento - garante que os dados enviados pelo simulador são recebidos pelo broker.

### Entrada:
- Nenhuma.

### Saída Esperada:
- Conection with broker MQTT successful.

## Vídeo





