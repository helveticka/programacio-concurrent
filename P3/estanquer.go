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
	// Informació de l'estanquer
	fmt.Println("\nHola, som l'estanquer il·legal\n")
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
	// Creació de la cua per rebre tabac
	tabacQueue, err := ch.QueueDeclare("tabac", false, false, false, false, nil)
	failOnError(err, "No es pot crear la cua de tabac")
	// Creació de la cua per rebre mistos
	mistosQueue, err := ch.QueueDeclare("mistos", false, false, false, false, nil)
	failOnError(err, "No es pot crear la cua de mistos")
	// Creació de la cua per rebre l'alerta de la policia
	alertQueue, err := ch.QueueDeclare("alerta", false, false, false, false, nil)
	failOnError(err, "No es pot crear la cua d'alerta")
	// Binding a la cua d'alertes
	err = ch.QueueBind(alertQueue.Name, "", "alerta", false, nil)
	// Variables
	var tabac int
	var mistos int
	var sortir = false
	// Go routine per gestionar els missatges
	go func() {
		// Missatges dels fumadors de tabac
		missatgeTabac, err := ch.Consume(tabacQueue.Name, "", true, false, false, false, nil)
		failOnError(err, "No es pot consumir la cua de tabac")
		// Missatges dels fumadors de mistos
		missatgeMistos, err := ch.Consume(mistosQueue.Name, "", true, false, false, false, nil)
		failOnError(err, "No es pot consumir la cua de mistos")
		// Missatges d'alerta de la policia
		missatgePolicia, err := ch.Consume(alertQueue.Name, "", true, false, false, false, nil)
		failOnError(err, "No es pot consumir la cua d'alerta")
		// Bucle per a gestionar els missatges
		for {
			select {
			// Cas en el que es rep un missatge de tabac
			case <-missatgeTabac:
				// Incrementar el comptador de tabac
				tabac++
				// Mostrar per pantalla que s'ha posat tabac a la taula
				fmt.Printf("He posat el tabac %d damunt la taula\n", tabac)
				// Publicar el tabac a la cua dels fumadors de tabac
				ch.Publish("", "fumadorTabac", false, false, amqp.Publishing{
					Body: []byte(fmt.Sprintf("%d", tabac)),
				})
				// Cas en el que es rep un missatge de mistos
			case <-missatgeMistos:
				// Incrementar el comptador de mistos
				mistos++
				// Mostrar per pantalla que s'ha posat un misto a la taula
				fmt.Printf("He posat el misto %d damunt la taula\n", mistos)
				// Publicar el misto a la cua dels fumadors de mistos
				ch.Publish("", "fumadorMistos", false, false, amqp.Publishing{
					Body: []byte(fmt.Sprintf("%d", mistos)),
				})
				// Cas en el que es rep un missatge d'alerta de la policia
			case <-missatgePolicia:
				// Sleep per simular que l'estanquer recull la taula
				temps := rand.Intn(1) + 1
				time.Sleep(time.Duration(temps) * time.Second)
				// Mostrar per pantalla que la policia ve
				fmt.Println("\nUyuyuy la policia! Men vaig \n. . . Men duc la taula!!!!")
				// Eliminar les cues
				ch.QueueDelete(tabacQueue.Name, false, false, false)
				ch.QueueDelete(mistosQueue.Name, false, false, false)
				ch.QueueDelete(alertQueue.Name, false, false, false)
				ch.QueueDelete("alertaFumadorTabac", false, false, false)
				ch.QueueDelete("alertaFumadorMistos", false, false, false)
				ch.QueueDelete("fumadorTabac", false, false, false)
				ch.QueueDelete("fumadorMistos", false, false, false)
				sortir = true
				return
			}
		}
	}()

	// Bucle per a esperar fins que arribi la policia
	for !sortir {

	}
}
