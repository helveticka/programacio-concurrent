package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Funció per gestionar errors
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Funció principal
func main() {
	// Informació del fumador xivato
	fmt.Println("\nNo sóm fumador. ALERTA! Que ve la policia!\n\n. . .\n")
	// Conexió amb el servidor de RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "No es pot connectar amb RabbitMQ")
	defer conn.Close()
	// Obrir un canal de comunicació amb el servidor
	ch, err := conn.Channel()
	failOnError(err, "No es pot obrir el canal")
	defer ch.Close()
	// Declarar l'exchange d'alerta
	err = ch.ExchangeDeclare(
		"alerta",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "No es pot declarar l'exchange")
	// Publicar un missatge d'alerta
	ch.Publish("alerta", "", false, false, amqp.Publishing{
		Body: []byte("Anem, que ve la policia!"),
	})
}
