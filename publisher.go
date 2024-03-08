package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"time"
)


type Message struct {
	ID       string    `json:"id"`
	Tipo     string    `json:"tipo"`
	Temperatura int     `json:"temperatura"`
	Timestamp string    `json:"timestamp"`
	Alarm    bool      `json:"alarm,omitempty"`
}

func printMessage(message Message) {
	alarmMessage := ""
	if message.Alarm {
		alarmMessage = fmt.Sprintf("[ALERTA: Temperatura %s]", func() string {
			if message.Temperatura < 0 {
				return "BAIXA"
			}
			return "ALTA"
		}())
	}

	temperature := message.Temperatura
	if message.Temperatura < 0 {
		temperature = -temperature
	}
	formattedTemperature := fmt.Sprintf("%d°C", temperature)

	fmt.Printf("Lj %s: %s %s| %s\n", message.ID[0:2], message.Tipo, message.ID[2:], formattedTemperature)
	if message.Alarm {
		fmt.Println(alarmMessage)
	}
}

func MsgSender() Message {
	rand.Seed(time.Now().Unix())
	temperature := rand.Intn(56) - 35
	tipo := ""
	if rand.Intn(2) == 0 {
		tipo = "freezer"
	} else {
		tipo = "geladeira"
	}
	id := "lj01f01"

	var message Message
	if tipo == "freezer" {
		message = Message{
			ID:       id,
			Tipo:     tipo,
			Temperatura:  temperature,
		}

		if temperature > -15 || temperature < -25 {
			freezerMsg := message
			freezerMsg.Alarm = true
			message = freezerMsg
		}
	} else {
		message = Message{
			ID:       id,
			Tipo:     tipo,
			Temperatura:  temperature,
		}

		if temperature > 10 || temperature < 2 {
			fridgeMsg := message
			fridgeMsg.Alarm = true
			message = fridgeMsg
		}
	}

	message.Timestamp = time.Now().Format(time.RFC3339)
	return message
}

func Client() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("publisher")

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	messageBuffer := make([]Message, 0)

	for {
		data := MsgSender()
		messageBuffer = append(messageBuffer, data)
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}

		msg := string(jsonData) + time.Now().Format(time.RFC3339) 

		token := client.Publish("/p1", 1, false, msg) //QoS 1
		token.Wait()

		// Print the buffered messages every 10 seconds
		if len(messageBuffer) >= 4 {
			printBufferedMessages(messageBuffer)
			messageBuffer = messageBuffer[:0] // Clear the buffer
		}

		fmt.Println("Published:", msg)
		time.Sleep(2 * time.Second)
	}
}
func printBufferedMessages(messages []Message) {
	for i, message := range messages {
		alarmMessage := ""
		if message.Alarm {
			alarmMessage = fmt.Sprintf(" [ALERTA: Temperatura %s]", func() string {
				if message.Temperatura < 0 {
					return "BAIXA"
				}
				return "ALTA"
			}())
		}

		temperature := message.Temperatura
		if message.Temperatura < 0 {
			temperature = -temperature
		}
		formattedTemperature := fmt.Sprintf("%d°C", temperature)

		fmt.Printf("Lj %d: %s %s| %s%s\n", i+1, message.ID[0:2], message.Tipo, message.ID[2:], formattedTemperature)
		if message.Alarm {
			fmt.Println(alarmMessage)
		}
	}
}

func main() {
	Client()

}

