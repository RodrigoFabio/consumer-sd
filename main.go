package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial(GetStringConnFila)
	nomeFila := os.Getenv("NOME_FILA")

	if err != nil {
		log.Fatalf("Falha ao se conectar ao RabbitMQ: %v verifique se o servidor de fila foi iniciado corretamente", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Fallha ao abrir canal: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		nomeFila,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Falha ao declarar a fila: %v Verifique se ela foi corretamente criada.", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Falha ao registrar um consumer: %v", err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	forever := make(chan bool)

	go func() {
		log.Println("Consumer iniciado. Aguardando mensagens...")
		for d := range msgs {
			log.Printf("message received: %v", string(d.Body))
			EnviarEmail(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-sigchan

	log.Printf("interrupted, shutting down")
	forever <- true
}
