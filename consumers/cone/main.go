package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
  // connection
  conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
  if err != nil {
    log.Panicf(" [CONSUMER#1] [PANIC] [Startup] conn error: %s", err)
  }
  defer conn.Close()
  // channel
  ch, err := conn.Channel()
  if err != nil {
    log.Panicf(" [CONSUMER#1] [PANIC] [Startup] chan error: %s", err)
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
    log.Panicf(" [CONSUMER#1] [PANIC] [Startup] exchange error: %s", err)
  }
  // queue declare
  q, err := ch.QueueDeclare(
    "queue#11", // queue name
    false, // durable
    false, // autoDelete
    true, // exclusive
    false, // noWait
    nil, // args
  )
  if err != nil {
    log.Panicf(" [CONSUMER#1] [PANIC] [Startup] queue declare error: %s", err)
  }
  // queue bind
  if err = ch.QueueBind(
    q.Name, // name
    "11", // routing key
    "exchange", // exchange
    false, // noWait
    nil, // args
  ); err != nil {
    log.Panicf(" [CONSUMER#1] [PANIC] [Startup] queue bind error: %s", err)
  }
  // consumer
  msgs, err := ch.Consume(
    q.Name, // queue
    "", // consumer
    true, // autoAck
    false, // exclusive
    false, // noLocal
    false, // noWait
    nil,
  )
  if err != nil {
    log.Panicf(" [CONSUMER#1] [PANIC] [Startup] consume error: %s", err)
  }
  var forever chan struct{}
  go func() {
    for d := range msgs {
      log.Printf(" [CONSUMER#1] [INFO] [Received] msg: %s", d.Body)
    }
  }()
  log.Print(" [CONSUMER#1] [INFO] [*] Waiting for messages")
  <-forever
}
