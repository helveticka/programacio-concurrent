package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

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
	// Informació del fumador de tabac
	fmt.Println("\nSóm fumador. Tinc mistos però me falta tabac\n")
	// Connexió amb el servidor de RabbitMQ
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
	// Creació de la cua per rebre tabac
	tabacQueue, err := ch.QueueDeclare("fumadorTabac", false, false, false, false, nil)
	failOnError(err, "No es pot crear la cua de tabac")
	// Creació de la cua per rebre alertes
	alertQueue, err := ch.QueueDeclare("alertaFumadorTabac", false, false, false, false, nil)
	failOnError(err, "No es pot crear la cua d'alerta")
	// Binding a la cua d'alertes
	err = ch.QueueBind(alertQueue.Name, "", "alerta", false, nil)
	failOnError(err, "No es pot fer el binding de la cua d'alerta")
	// Anar a cercar tabac
	ch.Publish("", "tabac", false, false, amqp.Publishing{
		Body: []byte("Petició de tabac"),
	})
	// Consumir missatges de la cua de tabac
	msgs, err := ch.Consume(tabacQueue.Name, "", true, false, false, false, nil)
	failOnError(err, "No es pot consumir de la cua de tabac")
	// Consumir missatges de la cua d'alertes
	alertMsgs, err := ch.Consume(alertQueue.Name, "", true, false, false, false, nil)
	failOnError(err, "No es pot consumir de la cua d'alerta")
	// Bucle per gestionar els tabacs
	for {
		select {
		// Cas en que arriba un missatge de la cua de tabac
		case msg := <-msgs:
			// Sleep per simular el temps que triga a agafar el tabac
			temps := rand.Intn(3) + 1
			time.Sleep(time.Duration(temps) * time.Second)
			fmt.Printf("He agafat el tabac %s. Gràcies! \n. . . \nMe dones més tabac?\n", string(msg.Body))
			// Sol·licitar més tabac
			ch.Publish("", "tabac", false, false, amqp.Publishing{
				Body: []byte("Petició de tabac"),
			})
			// Cas de rebre un missatge d'alerta
		case <-alertMsgs:
			fmt.Println("Anem, que ve la policia!")
			return
		}
	}
}