package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"testing"

)

// Testa a conexão com o Broker
func TestConection(t *testing.T) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("test-client")

	client := MQTT.NewClient(opts)
	//mensagem := MsgSender(resultado, "/pond2")

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Errorf("Error in connection with broker MQTT: %v", token.Error())
	} else {
		t.Log("Conection with broker MQTT successful")
	}

}

