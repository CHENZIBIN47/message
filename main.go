package main
import (
	"rabbitmq/rabbitmq/client"
	"rabbitmq/utils"
)

func main() {
	rabbitClient, err := client.NewRabbitClient("","")

	utils.FailOnError(err,"rabbit client fail")

	defer rabbitClient.Close()

}