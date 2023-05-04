package main

import (
	"context"
	"log"

	"github.com/VrMolodyakov/vgm/email/app/internal/app"
)

func main() {
	ctx := context.Background()
	a := app.New()

	defer func() {
		a.Close()
	}()

	if err := a.ReadConfig(); err != nil {
		log.Fatal(err, "read config")
		return
	}

	a.InitLogger()

	if err := a.InitTracer(); err != nil {
		log.Fatal(err, "init tracer")
		return
	}

	if err := a.Setup(); err != nil {
		log.Fatal(err, "setup dependencies")
		return
	}

	a.StartSubscriber(ctx)
	a.Start(ctx)

}
