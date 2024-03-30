package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cname := "CONSUMER#1"
	qname := "queue#11"
	rkey := "11"
	// connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Panicf(" [%s] [PANIC] [Startup] conn error: %s", cname, err)
	}
	defer conn.Close()
	// channel
	ch, err := conn.Channel()
	if err != nil {
		log.Panicf(" [%s] [PANIC] [Startup] chan error: %s", cname, err)
	}
	defer ch.Close()
	// exchange
	if err := ch.ExchangeDeclare(
		"exchange", // name
		"direct",   // kind
		true,       // durable
		false,      // autoDelete
		false,      // internal
		false,      // noWait
		nil,        // args
	); err != nil {
		log.Panicf(" [%s] [PANIC] [Startup] exchange error: %s", cname, err)
	}
	// queue declare
	q, err := ch.QueueDeclare(
		qname, // queue name
		false, // durable
		false, // autoDelete
		true,  // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		log.Panicf(" [%s] [PANIC] [Startup] queue declare error: %s", cname, err)
	}
	// queue bind
	if err = ch.QueueBind(
		q.Name,     // name
		rkey,       // routing key
		"exchange", // exchange
		false,      // noWait
		nil,        // args
	); err != nil {
		log.Panicf(" [%s] [PANIC] [Startup] queue bind error: %s", cname, err)
	}
	// consumer
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // autoAck
		false,  // exclusive
		false,  // noLocal
		false,  // noWait
		nil,
	)
	if err != nil {
		log.Panicf(" [%s] [PANIC] [Startup] consume error: %s", cname, err)
	}
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf(" [%s] [INFO] [Received] msg: %s", cname, d.Body)
		}
	}()
	log.Printf(" [%s] [INFO] [*] Waiting for messages", cname)
	<-forever
}
