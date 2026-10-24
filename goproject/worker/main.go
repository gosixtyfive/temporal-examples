package main

import (
	"log"

	"my-org/quickstart"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "my-task-queue", worker.Options{})

	w.RegisterWorkflow(quickstart.SayHelloWorkflow)
	w.RegisterActivity(quickstart.Greet)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

