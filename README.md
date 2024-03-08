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

# Explicação

Este aplicativo é um sistema simples de monitoramento de temperatura que gera valores aleatórios de temperatura para um freezer e um refrigerador. Ele usa MQTT (Message Queuing Telemetry Transport) para publicar dados de temperatura para um corretor MQTT especificado. O sistema envia alarmes caso a temperatura saia da faixa permitida.


## Estrutura da mensagem: 

representa a estrutura de dados de uma mensagem contendo os seguintes campos:

- ID: o ID do dispositivo (por exemplo, lj01f01)
- Tipo: O tipo de dispositivo (freezer ou geladeira)
- Temperatura: O valor da temperatura
- Timestamp: \data e hora de quando a medição foi realizada
- Alarme: Um sinalizador booleano para indicar se um alarme deve ser acionado com base no valor da temperatura
- Função printMessage: Formata e imprime a mensagem de temperatura em um formato legível por humanos

- Função MsgSender: gera uma temperatura aleatória para um freezer ou geladeira e cria um objeto Message

- Função cliente: Conecta-se a um corretor MQTT, publica mensagens de temperatura com QoS 1 e imprime as mensagens em buffer a cada 10 segundos

- Função printBufferedMessages: formata e imprime as mensagens de temperatura armazenadas em buffer em um formato legível


