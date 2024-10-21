package main

import (
	"flag"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/undeadops/temporal-examples/pkg/cron"
)

var (
	temporalHostPort string
)

func main() {
	flag.StringVar(&temporalHostPort, "temporalHostPort", "temporal-frontend:7233", "Temporal Host Port")
	flag.Parse()

	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: temporalHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "cron", worker.Options{})

	w.RegisterWorkflow(cron.SampleCronWorkflow)
	w.RegisterActivity(cron.DoSomething)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
