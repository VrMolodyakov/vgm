package main

import (
	"context"
	"log"

	"github.com/VrMolodyakov/vgm/youtube/internal/app"
)

func main() {
	ctx := context.Background()
	a := app.New()

	defer func() {
		a.Close(ctx)
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

	if err := a.Setup(ctx); err != nil {
		log.Fatal(err, "setup dependencies")
		return
	}

	a.Start(ctx)

}
