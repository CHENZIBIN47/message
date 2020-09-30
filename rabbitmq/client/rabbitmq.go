package client

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/utils"
)

type Client interface {
	Publish(body []byte)
	Close()
}


type MQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	Queue amqp.Queue
}

func NewRabbitClient(url string,queue string)(client *MQClient,err error){
	client = &MQClient{}
	client.conn, err = amqp.Dial(url)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	client.channel, err = client.conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	client.Queue, err = client.channel.QueueDeclare(queue, false, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare a queue")
	return
}

// Publish 发布
func (c *MQClient) Publish(body []byte) {
	if c.channel == nil {
		log.Fatalf("rabbitmq publish %s fail: invalie channel", c.Queue.Name)
		return
	}
	err := c.channel.Publish(
		"",
		c.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	utils.FailOnError(err,"rabbitmq publish %s success")
	return
}

// Close 关闭
func (c *MQClient) Close() {
	if c.conn != nil {
		err := c.conn.Close()
		utils.FailOnError(err,"c.conn.Close")
	}
	if c.channel != nil {
		err := c.channel.Close()
		utils.FailOnError(err,"c.channel.Close")
	}
}



