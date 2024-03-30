package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	pname := "PRODUCER#1"
	// grab routing key from launch cmd
	rkey := os.Args[1]
	// validate rkey
	if !utils.IsValidArg(rkey) {
		panic(
			errors.New(
				fmt.Sprintf("InvalidRoutingKey: provided %s; must be in ['11', '22', '33']!", rkey),
			),
		)
	}
	destname := utils.GetExpectedDestination(rkey)
	// connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Panicf(" [%s] [PANIC] [Startup] conn error: %s", pname, err)
	}
	defer conn.Close()
	// channel
	ch, err := conn.Channel()
	if err != nil {
		log.Panicf(" [%s] [PANIC] [Startup] chan error: %s", pname, err)
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
		log.Panicf(" [%s] [PANIC] [Startup] exchange error: %s", pname, err)
	}
	// context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// message
	body := fmt.Sprintf("from:[%s] with routing key %s target:[%s]", pname, rkey, destname)
	if err := ch.PublishWithContext(
		ctx,        // ctx
		"exchange", // exchange
		rkey,       // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		}, // msg
	); err != nil {
		log.Panicf(" [%s] [ERROR] [Send] send error: %s", pname, err)
	}
	log.Printf(" [%s] [INFO] [*] Sent message: %s", pname, body)
}
