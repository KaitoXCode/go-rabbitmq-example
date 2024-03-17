package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
  // connection
  conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
  if err != nil {
    log.Panicf(" [PRODUCER#1] [PANIC] [Startup] conn error: %s", err)
  }
  defer conn.Close()
  // channel
  ch, err := conn.Channel()
  if err != nil {
    log.Panicf(" [PRODUCER#1] [PANIC] [Startup] chan error: %s", err)
  }
  defer ch.Close()
  // exchange
  if err := ch.ExchangeDeclare(
    "exchange", // name
    "fanout", // kind
    true, // durable
    false, // autoDelete
    false, // internal
    false, // noWait
    nil, // args
  ); err != nil {
    log.Panicf(" [PRODUCER#1] [PANIC] [Startup] exchange error: %s", err)
  }
  // context
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  // message
  body := "from producer #1 with routing key 11 target consumer 11"
  if err := ch.PublishWithContext(
    ctx, // ctx
    "exchange", // exchange
    "11", // routing key
    false, // mandatory
    false, // immediate
    amqp.Publishing{
      ContentType: "text/plain",
      Body: []byte(body),
    }, // msg
  ); err != nil {
    log.Panicf(" [PRODUCER#1] [ERROR] [Send] send error: %s", err)
  }
  log.Printf(" [PRODUCER#1] [INFO] [*] Sent message: %s", body)
}
